package message_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/ONSdigital/dp-dimension-importer/message"
	mock "github.com/ONSdigital/dp-dimension-importer/message/mock"
	kafka "github.com/ONSdigital/dp-kafka"
	"github.com/ONSdigital/dp-kafka/kafkatest"
	"github.com/ONSdigital/log.go/log"
	. "github.com/smartystreets/goconvey/convey"
)

var (
	errExpected   = errors.New("bork")
	errExitedTest = errors.New("forced test exit")
)

const (
	fileURL           = "test-url"
	instanceID        = "1234567890"
	timeout           = 5
	timeoutFailureMsg = "Concurrent Test did not complete within the configured timeout window. Failing test."
	consumerExitedMsg = "Consumer exited"
	consumerCloserMsg = "consumer closer called"
	producerCloserMsg = "Producer closer called"
)

func TestConsumer_Listen(t *testing.T) {
	Convey("Given the Consumer has been configured correctly", t, func() {
		cgChannels := kafka.CreateConsumerGroupChannels(true)
		handlerInvoked := make(chan kafka.Message, 1)

		kafkaConsumer := &kafkatest.IConsumerGroupMock{
			ChannelsFunc: func() *kafka.ConsumerGroupChannels { return cgChannels },
			ReleaseFunc:  func() {},
		}

		handleCalls := []kafka.Message{}
		recieverMock := mock.MessageReciever{
			OnMessageFunc: func(message kafka.Message) {
				handleCalls = append(handleCalls, message)
				handlerInvoked <- message
			},
		}

		msg := &kafkatest.MessageMock{}

		consumer := message.NewConsumer(kafkaConsumer, recieverMock, time.Second*10)
		consumer.Listen()

		Convey("When the consumer receives a valid message", func() {
			cgChannels.Upstream <- msg

			select {
			case <-handlerInvoked:
				log.Event(context.Background(), "Handler invoked", log.INFO)
			case <-time.After(time.Second * 3):
				log.Event(context.Background(), "Test timed out.", log.INFO)
				t.FailNow()
			}
			consumer.Close(nil)

			Convey("Then messageReciever.OnMessage is called 1 time with the expected parameters", func() {
				So(len(handleCalls), ShouldEqual, 1)
				So(len(kafkaConsumer.ReleaseCalls()), ShouldEqual, 1)
			})
		})

	})
}
