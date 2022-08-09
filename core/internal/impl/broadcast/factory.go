package broadcast

import (
	"github.com/blazejsewera/notipie/core/internal/domain"
	"github.com/blazejsewera/notipie/core/internal/grid"
	"github.com/blazejsewera/notipie/core/internal/impl/net/ws"
)

type BroadcasterFactoryFunc func() domain.NotificationBroadcaster

// @impl
var _ grid.BroadcasterFactory = BroadcasterFactoryFunc(nil)

func (f BroadcasterFactoryFunc) GetBroadcaster() domain.NotificationBroadcaster {
	return f()
}

var WebSocketBroadcasterFactory = BroadcasterFactoryFunc(func() domain.NotificationBroadcaster {
	hub := ws.WebSocketHubFactory.GetHub()
	return NewWebSocketBroadcaster(hub)
})
