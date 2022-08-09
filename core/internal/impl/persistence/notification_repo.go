package persistence

import (
	"github.com/blazejsewera/notipie/core/internal/domain"
	"github.com/blazejsewera/notipie/core/pkg/lib/log"
	"go.uber.org/zap"
)

type MemRepository struct {
	notifications []domain.Notification
	l             *zap.Logger
}

// @impl
var _ domain.NotificationRepository = (*MemRepository)(nil)

func NewMemRepository() *MemRepository {
	return &MemRepository{l: log.For("impl").Named("persistence").Named("notification_repo")}
}

func (r *MemRepository) SaveNotification(notification domain.Notification) {
	r.notifications = append(r.notifications, notification)
	r.l.Debug("saved notification", zap.String("notificationID", notification.ID))
}

func (r *MemRepository) GetLastNotifications(n int) []domain.Notification {
	return r.notifications[len(r.notifications)-n:]
}

func (r *MemRepository) GetNotifications(from, to int) []domain.Notification {
	return r.notifications[from:to]
}

func (r *MemRepository) GetNotificationCount() int {
	return len(r.notifications)
}
