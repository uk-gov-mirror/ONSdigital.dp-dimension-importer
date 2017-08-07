package main

import (
	"os"

	"github.com/ONSdigital/dp-dimension-importer/client"
	"github.com/ONSdigital/dp-dimension-importer/config"
	"github.com/ONSdigital/dp-dimension-importer/handler"
	"github.com/ONSdigital/dp-dimension-importer/message"
	"github.com/ONSdigital/dp-dimension-importer/repository"
	"github.com/ONSdigital/go-ns/kafka"
	"github.com/ONSdigital/go-ns/log"
	logKeys "github.com/ONSdigital/dp-dimension-importer/common"
	"fmt"
)

var incomingKafka chan kafka.Message

func main() {
	log.Namespace = "dimension-importer"
	cfg, err := config.Load()
	if err != nil {
		os.Exit(1)
	}

	log.Debug("Application configuration", log.Data{"config": cfg})

	consumer, err := kafka.NewConsumerGroup(cfg.KafkaAddr, cfg.DimensionsExtractedTopic, log.Namespace, kafka.OffsetNewest)
	if err != nil {
		log.ErrorC("Could not create consumer", err, nil)
		panic("Could not create consumer")
	}

	client.Host = cfg.ImportAddr
	fmt.Println(cfg.ImportAuthToken)
	client.AuthToken = cfg.ImportAuthToken

	var databaseClient *client.Neo4j
	if databaseClient, err = client.NewDatabase(cfg.DatabaseURL, cfg.PoolSize); err != nil {
		log.ErrorC("Unexpected error while to create database connection pool", err, log.Data{
			logKeys.URL:      cfg.DatabaseURL,
			logKeys.PoolSize: cfg.PoolSize,
		})
		os.Exit(1)
	}

	newDimensionInserterFunc := func() handler.DimensionRepository {
		return repository.DimensionRepository{
			Neo:              *databaseClient,
			ConstraintsCache: map[string]string{},
		}
	}

	eventHandler := &handler.DimensionsExtractedEventHandler{
		NewDimensionInserter: newDimensionInserterFunc,
		InstanceRepository:   &repository.InstanceRepository{Neo: databaseClient},
		ImportAPI:            client.ImportAPI{},
	}

	err = message.Consume(consumer, eventHandler)
	if err != nil {
		log.ErrorC("consumer", err, nil)
		panic("Consumer returned error")
	}
}
