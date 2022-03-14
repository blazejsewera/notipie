package broadcast

import (
	"github.com/blazejsewera/notipie/core/internal/domain"
	"github.com/blazejsewera/notipie/core/internal/impl/net/ws"
	"github.com/blazejsewera/notipie/core/internal/model"
	"github.com/blazejsewera/notipie/core/pkg/lib/log"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type WebSocketBroadcaster struct {
	hub ws.Hub
	l   *zap.Logger
}

//@impl
var _ domain.NotificationBroadcaster = (*WebSocketBroadcaster)(nil)

func NewWebSocketBroadcaster(hub ws.Hub) *WebSocketBroadcaster {
	return &WebSocketBroadcaster{
		hub: hub,
		l:   log.For("impl").Named("broadcast").Named("notification_broadcaster"),
	}
}

func (b *WebSocketBroadcaster) Broadcast(notification domain.Notification) {
	cn := model.ClientNotificationFromDomain(notification)
	b.hub.Broadcast(cn)
	b.l.Debug("broadcast notification to hub", zap.String("notificationID", notification.ID))
}

func (b *WebSocketBroadcaster) RegisterClient(conn interface{}) {
	b.hub.Register(conn.(*websocket.Conn))
}
