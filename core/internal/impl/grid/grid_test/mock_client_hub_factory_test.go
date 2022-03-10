package grid_test

import (
	"github.com/blazejsewera/notipie/core/internal/impl/model"
	"github.com/blazejsewera/notipie/core/internal/impl/net/ws"
	"github.com/blazejsewera/notipie/core/pkg/lib/util"
	"github.com/gorilla/websocket"
)

var MockClientHubFactory = ws.ClientHubFactoryFunc(func() ws.ClientHub {
	return NewMockClientHub()
})

type MockClientHub struct {
	Notifications []model.ClientNotification
	Done          chan struct{}
}

func (m *MockClientHub) Start() {}

func (m *MockClientHub) Broadcast(notification model.ClientNotification) {
	m.Notifications = append(m.Notifications, notification)
	m.Done <- struct{}{}
}

func (m *MockClientHub) Register(conn *websocket.Conn) {}

func (m *MockClientHub) Unregister(clientUUID string) {}

func NewMockClientHub() *MockClientHub {
	return &MockClientHub{Done: make(chan struct{})}
}

var _ ws.ClientHub = (*MockClientHub)(nil)
var _ util.Starter = (*MockClientHub)(nil)
