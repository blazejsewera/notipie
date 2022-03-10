package broadcast

import (
	"github.com/blazejsewera/notipie/core/internal/domain"
	"github.com/blazejsewera/notipie/core/internal/impl/model"
	"github.com/blazejsewera/notipie/core/internal/impl/net/ws"
	"github.com/blazejsewera/notipie/core/pkg/lib/log"
	"github.com/blazejsewera/notipie/core/pkg/lib/timeformat"
	"github.com/blazejsewera/notipie/core/pkg/lib/util"
	"go.uber.org/zap"
)

type WebSocketNotificationBroadcaster struct {
	hub ws.ClientHub
	l   *zap.Logger
}

// WebSocketNotificationBroadcaster implements interfaces below
var _ domain.NotificationBroadcaster = (*WebSocketNotificationBroadcaster)(nil)
var _ util.Starter = (*WebSocketNotificationBroadcaster)(nil)

func NewWebSocketNotificationBroadcaster(hub ws.ClientHub) *WebSocketNotificationBroadcaster {
	return &WebSocketNotificationBroadcaster{
		hub: hub,
		l:   log.For("impl").Named("broadcast").Named("notification_broadcaster"),
	}
}

func (b *WebSocketNotificationBroadcaster) Start() {
	b.hub.Start()
}

func (b *WebSocketNotificationBroadcaster) Broadcast(notification domain.Notification) {
	cn := clientNotificationOf(notification)
	b.hub.GetBroadcastChan() <- cn
	b.l.Debug("broadcast notification to hub", zap.String("notificationID", notification.ID))
}

func clientNotificationOf(n domain.Notification) model.ClientNotification {
	timestamp := n.Timestamp.Format(timeformat.RFC3339Milli)
	return model.ClientNotification{
		HashableNetNotification: model.HashableNetNotification{
			AppName:    n.App.Name,
			AppID:      n.App.ID,
			AppImgURI:  n.App.IconURI,
			Title:      n.Title,
			Subtitle:   n.Subtitle,
			Body:       n.Body,
			ExtURI:     n.ExtURI,
			ReadURI:    n.ReadURI,
			ArchiveURI: n.ArchiveURI,
		},
		ID:        n.ID,
		Timestamp: timestamp,
		Read:      false,
	} // TODO: implement urgency
}
