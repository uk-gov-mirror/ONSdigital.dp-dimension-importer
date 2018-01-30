package config

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/ONSdigital/go-ns/log"
	"github.com/kelseyhightower/envconfig"
)

// Config struct to hold application configuration.
type Config struct {
	BindAddr                       string        `envconfig:"BIND_ADDR"`
	KafkaAddr                      []string      `envconfig:"KAFKA_ADDR"`
	IncomingInstancesTopic         string        `envconfig:"DIMENSIONS_EXTRACTED_TOPIC"`
	IncomingInstancesConsumerGroup string        `envconfig:"DIMENSIONS_EXTRACTED_CONSUMER_GROUP"`
	OutgoingInstancesTopic         string        `envconfig:"DIMENSIONS_INSERTED_TOPIC"`
	EventReporterTopic             string        `envconfig:"EVENT_REPORTER_TOPIC"`
	DatasetAPIAddr                 string        `envconfig:"DATASET_API_ADDR"`
	DatasetAPIAuthToken            string        `envconfig:"DATASET_API_AUTH_TOKEN"`
	DatabaseURL                    string        `envconfig:"DB_URL"`
	PoolSize                       int           `envconfig:"DB_POOL_SIZE"`
	GracefulShutdownTimeout        time.Duration `envconfig:"GRACEFUL_SHUTDOWN_TIMEOUT"`
	HealthCheckInterval            time.Duration `envconfig:"HEALTHCHECK_INTERVAL"`
}

func (c *Config) String() string {
	authTokenFound := "NOT FOUND"
	if len(c.DatasetAPIAuthToken) > 0 {
		authTokenFound = "FOUND"
	}

	masked := Config(*c)
	masked.DatasetAPIAuthToken = authTokenFound

	b, _ := json.Marshal(masked)
	return string(b)
}

// Load load the configuration & apply defaults where necessary
func Load() (*Config, error) {
	cfg := Config{
		BindAddr:                       ":23000",
		DatasetAPIAddr:                 "http://localhost:22000",
		DatasetAPIAuthToken:            "FD0108EA-825D-411C-9B1D-41EF7727F465",
		DatabaseURL:                    "bolt://localhost:7687",
		PoolSize:                       20,
		KafkaAddr:                      []string{"localhost:9092"},
		IncomingInstancesTopic:         "dimensions-extracted",
		IncomingInstancesConsumerGroup: "dp-dimension-importer",
		OutgoingInstancesTopic:         "dimensions-inserted",
		EventReporterTopic:             "report-events",
		GracefulShutdownTimeout:        time.Second * 5,
		HealthCheckInterval:            time.Minute,
	}

	err := envconfig.Process("", &cfg)

	if len(cfg.DatasetAPIAuthToken) == 0 {
		err := errors.New("error while attempting to load config. dataset api auth token is required but has not been configured")
		log.Error(err, nil)
		return nil, err
	}
	return &cfg, err
}
