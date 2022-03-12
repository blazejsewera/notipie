package persistence

import (
	"github.com/blazejsewera/notipie/core/internal/domain"
	"github.com/blazejsewera/notipie/core/pkg/lib/log"
	"go.uber.org/zap"
)

type MemRealtimeNotificationRepository struct {
	notifications []domain.Notification
	l             *zap.Logger
}

// MemRealtimeNotificationRepository implements interfaces below
var _ domain.NotificationRepository = (*MemRealtimeNotificationRepository)(nil)

func NewMemRealtimeNotificationRepository() *MemRealtimeNotificationRepository {
	return &MemRealtimeNotificationRepository{l: log.For("impl").Named("persistence").Named("notification_repo")}
}

func (r *MemRealtimeNotificationRepository) SaveNotification(notification domain.Notification) {
	r.notifications = append(r.notifications, notification)
	r.l.Debug("saved notification", zap.String("notificationID", notification.ID))
}

func (r *MemRealtimeNotificationRepository) GetLastNotifications(n int) []domain.Notification {
	return r.notifications[len(r.notifications)-n:]
}

func (r *MemRealtimeNotificationRepository) GetNotifications(from, to int) []domain.Notification {
	return r.notifications[from:to]
}

func (r *MemRealtimeNotificationRepository) GetNotificationCount() int {
	return len(r.notifications)
}
