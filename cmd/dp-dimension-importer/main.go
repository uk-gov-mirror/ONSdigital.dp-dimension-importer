package main

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"

	"github.com/ONSdigital/dp-dimension-importer/client"
	"github.com/ONSdigital/dp-dimension-importer/config"
	"github.com/ONSdigital/dp-dimension-importer/handler"
	"github.com/ONSdigital/dp-dimension-importer/initialise"
	"github.com/ONSdigital/dp-dimension-importer/message"
	"github.com/ONSdigital/dp-dimension-importer/schema"
	"github.com/ONSdigital/dp-dimension-importer/store"
	"github.com/ONSdigital/dp-graph/v2/graph"
	"github.com/ONSdigital/dp-healthcheck/healthcheck"
	kafka "github.com/ONSdigital/dp-kafka/v2"
	dphttp "github.com/ONSdigital/dp-net/http"
	"github.com/ONSdigital/dp-reporter-client/reporter"
	"github.com/ONSdigital/log.go/log"
	"github.com/gorilla/mux"
)

var (
	// BuildTime represents the time in which the service was built
	BuildTime string
	// GitCommit represents the commit (SHA-1) hash of the service that is running
	GitCommit string
	// Version represents the version of the service that is running
	Version string
)

func main() {
	log.Namespace = "dimension-importer"
	ctx := context.Background()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	cfg, err := config.Get(ctx)
	if err != nil {
		log.Event(ctx, "config load returned an error", log.FATAL, log.Error(err))
		os.Exit(1)
	}

	log.Event(ctx, "application configuration", log.INFO, log.Data{"config": cfg})

	// serviceList keeps track of what dependency services have been initialised
	serviceList := initialise.ExternalServiceList{}

	// Incoming kafka topic for instances to process
	instanceConsumer, err := serviceList.GetConsumer(ctx, cfg)
	if err != nil {
		os.Exit(1)
	}

	// Outgoing topic for instances that have completed processing
	instanceCompleteProducer, err := serviceList.GetProducer(ctx, cfg.OutgoingInstancesTopic, initialise.InstanceComplete, cfg)
	if err != nil {
		os.Exit(1)
	}

	// Outgoing topic for any errors while processing an instance
	errorReporterProducer, err := serviceList.GetProducer(ctx, cfg.EventReporterTopic, initialise.ErrorReporter, cfg)
	if err != nil {
		os.Exit(1)
	}

	// Connection to graph DB
	graphDB, err := serviceList.GetGraphDB(ctx)
	if err != nil {
		os.Exit(1)
	}

	var graphErrorConsumer *graph.ErrorConsumer
	if serviceList.GraphDB {
		graphErrorConsumer = graph.NewLoggingErrorConsumer(ctx, graphDB.ErrorChan())
	}

	// MessageProducer for instanceComplete events.
	instanceCompletedProducer := message.InstanceCompletedProducer{
		Producer:   instanceCompleteProducer,
		Marshaller: schema.InstanceCompletedSchema,
	}

	// Dataset Client wrapper.
	datasetAPICli, err := client.NewDatasetAPIClient(cfg)

	// Receiver for NewInstance events.
	instanceEventHandler := &handler.InstanceEventHandler{
		Store:         graphDB,
		DatasetAPICli: datasetAPICli,
		Producer:      instanceCompletedProducer,
		BatchSize:     cfg.BatchSize,
	}

	// Errors handler
	errorReporter, err := reporter.NewImportErrorReporter(errorReporterProducer, log.Namespace)
	if err != nil {
		log.Event(ctx, "new import error reporter error", log.FATAL, log.Error(err))
		os.Exit(1)
	}

	// Create healthcheck object with versionInfo
	hc, err := serviceList.GetHealthChecker(ctx, BuildTime, GitCommit, Version, cfg)
	if err != nil {
		os.Exit(1)
	}

	if err := registerCheckers(hc, instanceConsumer, instanceCompleteProducer, errorReporterProducer, datasetAPICli.Client, graphDB); err != nil {
		os.Exit(1)
	}

	httpServer := startHealthCheck(ctx, hc, cfg.BindAddr)

	messageReceiver := message.KafkaMessageReceiver{
		InstanceHandler: instanceEventHandler,
		ErrorReporter:   errorReporter,
	}

	// Create Consumer with kafkaConsmer
	consumer := serviceList.NewConsumer(ctx, instanceConsumer, messageReceiver, cfg.GracefulShutdownTimeout)
	consumer.Listen()

	instanceConsumer.Channels().LogErrors(ctx, "incoming instance kafka consumer received an error")
	instanceCompleteProducer.Channels().LogErrors(ctx, "completed instance kafka producer received an error")
	errorReporterProducer.Channels().LogErrors(ctx, "error reporter kafka producer received an error")

	// If we receive a signal (SIGINT or SIGTERM), start graceful shutdown
	signal := <-signals
	log.Event(ctx, "os signal received, attempting graceful shutdown", log.INFO, log.Data{"signal": signal.String()})

	shutdownCtx, cancel := context.WithTimeout(ctx, cfg.GracefulShutdownTimeout)
	hasShutdownError := false

	go func() {

		defer cancel() // cancel shutdown context timer

		if serviceList.HealthCheck {
			log.Event(ctx, "stopping health check", log.INFO)
			hc.Stop()
		}

		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			log.Event(ctx, "error shutting down http server", log.ERROR, log.Error(err))
			hasShutdownError = true
		}

		if serviceList.InstanceConsumer {
			log.Event(shutdownCtx, "stop listening to instance kafka consumer", log.INFO)
			if err := instanceConsumer.StopListeningToConsumer(shutdownCtx); err != nil {
				log.Event(ctx, "error on stop listening to instance kafka consumer", log.ERROR, log.Error(err))
				hasShutdownError = true
			}
		}

		if serviceList.Consumer {
			log.Event(ctx, "closing event consumer", log.INFO)
			consumer.Close(shutdownCtx)
		}

		if serviceList.InstanceConsumer {
			log.Event(shutdownCtx, "closing instance kafka consumer", log.INFO)
			if err := instanceConsumer.Close(shutdownCtx); err != nil {
				log.Event(ctx, "error closing instance kafka consumer", log.ERROR, log.Error(err))
				hasShutdownError = true
			}
		}

		if serviceList.InstanceCompleteProducer {
			log.Event(shutdownCtx, "closing instance complete kafka producer", log.INFO)
			if err := instanceCompleteProducer.Close(shutdownCtx); err != nil {
				log.Event(ctx, "error closing instance complete kafka consumer", log.ERROR, log.Error(err))
				hasShutdownError = true
			}
		}

		if serviceList.GraphDB {
			log.Event(shutdownCtx, "closing graph db", log.INFO)
			if err := graphDB.Close(shutdownCtx); err != nil {
				log.Event(ctx, "error closing graph db", log.ERROR, log.Error(err))
				hasShutdownError = true
			}

			log.Event(shutdownCtx, "closing graph db error consumer", log.INFO)
			if err := graphErrorConsumer.Close(shutdownCtx); err != nil {
				log.Event(ctx, "error closing graph db error consumer", log.ERROR, log.Error(err))
				hasShutdownError = true
			}
		}

		if serviceList.ErrorReporterProducer {
			log.Event(shutdownCtx, "closing error reporter kafka producer", log.INFO)
			if err := errorReporterProducer.Close(shutdownCtx); err != nil {
				log.Event(ctx, "error closing error reporter kafka producer", log.ERROR, log.Error(err))
				hasShutdownError = true
			}
		}
	}()

	// wait for timeout or success (cancel)
	<-shutdownCtx.Done()

	if hasShutdownError {
		err = errors.New("failed to shutdown gracefully")
		log.Event(ctx, "failed to shutdown gracefully ", log.ERROR, log.Error(err))
		os.Exit(1)
	}

	log.Event(ctx, "graceful shutdown complete", log.INFO)
	os.Exit(0)
}

// StartHealthCheck sets up the Handler, starts the healthcheck and the http server that serves health endpoint
func startHealthCheck(ctx context.Context, hc *healthcheck.HealthCheck, bindAddr string) *dphttp.Server {
	router := mux.NewRouter()
	router.Path("/health").HandlerFunc(hc.Handler)
	hc.Start(ctx)

	httpServer := dphttp.NewServer(bindAddr, router)
	httpServer.HandleOSSignals = false

	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			log.Event(ctx, "http server error", log.ERROR, log.Error(err))
			hc.Stop()
		}
	}()
	return httpServer
}

// RegisterCheckers adds the checkers for the provided clients to the healthcheck object.
func registerCheckers(hc *healthcheck.HealthCheck,
	instanceConsumer *kafka.ConsumerGroup,
	instanceCompleteProducer *kafka.Producer,
	errorReporterProducer *kafka.Producer,
	datasetClient client.IClient,
	db store.Storer) (err error) {

	hasErrors := false

	if err = hc.AddCheck("Kafka Instance Consumer", instanceConsumer.Checker); err != nil {
		hasErrors = true
		log.Event(nil, "error adding check for kafka instance consumer checker", log.ERROR, log.Error(err))
	}

	if err = hc.AddCheck("Kafka InstanceComplete Producer", instanceCompleteProducer.Checker); err != nil {
		hasErrors = true
		log.Event(nil, "error adding check for kafka instance complete producer checker", log.ERROR, log.Error(err))
	}

	if err = hc.AddCheck("Kafka ErrorReporter Producer", errorReporterProducer.Checker); err != nil {
		hasErrors = true
		log.Event(nil, "error adding check for kafka error reporter checker", log.ERROR, log.Error(err))
	}

	if err = hc.AddCheck("Dataset", datasetClient.Checker); err != nil {
		hasErrors = true
		log.Event(nil, "error adding check for dataset checker", log.ERROR, log.Error(err))
	}

	if err = hc.AddCheck("Graph DB", db.Checker); err != nil {
		hasErrors = true
		log.Event(nil, "error adding check for graph db", log.ERROR, log.Error(err))
	}

	if hasErrors {
		return errors.New("Error(s) registering checkers for healthcheck")
	}
	return nil
}
