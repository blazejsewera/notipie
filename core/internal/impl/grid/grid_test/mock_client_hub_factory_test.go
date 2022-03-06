package grid_test

import (
	"github.com/blazejsewera/notipie/core/internal/impl/model"
	"github.com/blazejsewera/notipie/core/internal/impl/net/ws"
	"github.com/gorilla/websocket"
)

type MockClientHubFactory struct{}

func (f MockClientHubFactory) GetClientHub() ws.ClientHub {
	return NewMockClientHub()
}

type MockClientHub struct {
	broadcastChan  chan model.ClientNotification
	registerChan   chan *websocket.Conn
	unregisterChan chan string
}

func NewMockClientHub() *MockClientHub {
	return &MockClientHub{
		broadcastChan:  make(chan model.ClientNotification),
		registerChan:   make(chan *websocket.Conn),
		unregisterChan: make(chan string),
	}
}

func (m *MockClientHub) GetBroadcastChan() chan model.ClientNotification {
	return m.broadcastChan
}

func (m *MockClientHub) GetRegisterChan() chan *websocket.Conn {
	return m.registerChan
}

func (m *MockClientHub) GetUnregisterChan() chan string {
	return m.unregisterChan
}

func (m *MockClientHub) Start() {}
