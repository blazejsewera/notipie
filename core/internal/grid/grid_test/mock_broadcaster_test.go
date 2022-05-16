package grid_test

import (
	"github.com/blazejsewera/notipie/core/internal/domain"
	"github.com/blazejsewera/notipie/core/pkg/model"
)

type mockBroadcaster struct {
	Done               chan struct{}
	ClientNotification model.ClientNotification
}

//@impl
var _ domain.NotificationBroadcaster = (*mockBroadcaster)(nil)

var mockBroadcasterInstance = &mockBroadcaster{Done: make(chan struct{})}

func (m *mockBroadcaster) Broadcast(notification domain.Notification) {
	m.ClientNotification = model.ClientNotificationFromDomain(notification)
	m.Done <- struct{}{}
}

func (m *mockBroadcaster) RegisterClient(interface{}) {}
