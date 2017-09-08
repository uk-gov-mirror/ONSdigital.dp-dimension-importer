package message

import (
	logKeys "github.com/ONSdigital/dp-dimension-importer/common"
	"github.com/ONSdigital/dp-dimension-importer/event"
	"github.com/ONSdigital/dp-dimension-importer/schema"
	"github.com/ONSdigital/go-ns/kafka"
	"github.com/ONSdigital/go-ns/log"
)

//go:generate moq -out ./message_test/consumer_generated_mocks.go -pkg message_test . KafkaMessage KafkaConsumer CompletedProducer ErrorEventHandler

const (
	eventRecieved       = "Recieved NewInstanceEvent"
	eventKey            = "event"
	eventHandlerErr     = "Unexpected error encountered while handling NewInstanceEvent"
	eventHandlerSuccess = "Instance has been successfully imported"
	errorRecieved       = "Consumer exit channel recieved error. Exiting dimensionExtractedConsumer"
	conumserErrMsg      = "Kafka Consumer Error recieved"
	producerErrMsg      = "Completed Error recieved"
	consumerStoppedMsg  = "Exiting Consumer loop"
	unmarshallErrMsg    = "Unexpected error when unmarshalling avro message to newInstanceEvent"
)

var closer chan bool = make(chan bool)

// KafkaMessage type representing a kafka message.
type KafkaMessage kafka.Message

// eventHandler defines an eventHandler.
type EventHandler interface {
	HandleEvent(event event.NewInstanceEvent) error
}

type KafkaConsumer interface {
	Incoming() chan kafka.Message
}

// Completed defines an KafkaProducer for dimensions inserted events
type CompletedProducer interface {
	Completed(e event.InstanceCompletedEvent) error
}

type ErrorEventHandler interface {
	Handle(instanceID string, err error, data log.Data)
}

func CloseConsumer() {
	closer <- true
}

func Consume(consumer KafkaConsumer, producer CompletedProducer, eventHandler EventHandler, errorEventHandler ErrorEventHandler) {
	go func() {
		consuming := true
		for consuming {
			select {
			case consumedMessage := <-consumer.Incoming():
				processMessage(consumedMessage.GetData(), producer, eventHandler, errorEventHandler)
				consumedMessage.Commit()
			case <-closer:
				log.Info(consumerStoppedMsg, nil)
				consuming = false
			}
		}
	}()
}

func processMessage(consumedData []byte, producer CompletedProducer, eventHandler EventHandler, errorEventHandler ErrorEventHandler) {
	var newInstanceEvent event.NewInstanceEvent
	if err := schema.NewInstanceSchema.Unmarshal(consumedData, &newInstanceEvent); err != nil {
		log.ErrorC(unmarshallErrMsg, err, nil)
		errorEventHandler.Handle("", err, nil)
		return
	}

	logData := map[string]interface{}{eventKey: newInstanceEvent}
	log.Debug(eventRecieved, logData)

	if err := eventHandler.HandleEvent(newInstanceEvent); err != nil {
		log.ErrorC(eventHandlerErr, err, logData)
		errorEventHandler.Handle(newInstanceEvent.InstanceID, err, nil)
		return
	}

	logData[logKeys.InstanceID] = newInstanceEvent.InstanceID
	log.Debug(eventHandlerSuccess, logData)

	insertedEvent := event.InstanceCompletedEvent{
		FileURL:    newInstanceEvent.FileURL,
		InstanceID: newInstanceEvent.InstanceID,
	}

	if err := producer.Completed(insertedEvent); err != nil {
		errorEventHandler.Handle(newInstanceEvent.InstanceID, err, nil)
	}
}
