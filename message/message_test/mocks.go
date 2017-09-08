package message_test

import (
	"github.com/ONSdigital/dp-dimension-importer/event"
)

type EventHandlerMock struct {
	Param           []event.NewInstanceEvent
	HandleEventFunc func(event.NewInstanceEvent) error
}

func (m *EventHandlerMock) HandleEvent(event event.NewInstanceEvent) error {
	m.Param = append(m.Param, event)
	return m.HandleEventFunc(event)
}

type MessageMock struct {
	Committed chan bool
	Data []byte
}

func (m *MessageMock) Commit() {
	m.Committed <- true
}

func (m *MessageMock) GetData() []byte {
	return m.Data
}