package persistence

import (
	"github.com/blazejsewera/notipie/core/internal/domain"
	"github.com/blazejsewera/notipie/core/pkg/lib/log"
	"go.uber.org/zap"
)

type RealtimeNotificationRepo interface {
	GetNotificationChan() chan domain.Notification
	domain.NotificationRepository
}

type MemRealtimeNotificationRepository struct {
	notifications    []domain.Notification
	notificationChan chan domain.Notification
	l                *zap.Logger
}

func NewMemRealtimeNotificationRepository() *MemRealtimeNotificationRepository {
	return &MemRealtimeNotificationRepository{
		notificationChan: make(chan domain.Notification),
		l:                log.For("impl").Named("persistence").Named("notification_repo"),
	}
}

func (r *MemRealtimeNotificationRepository) SaveNotification(notification domain.Notification) {
	r.l.Debug("received notification", zap.String("notificationID", notification.ID), zap.String("notificationTitle", notification.Title), zap.String("notificationAppID", notification.App.ID))
	r.notifications = append(r.notifications, notification)
	r.notificationChan <- notification
	r.l.Debug("sent notification to notificationChan")
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

func (r *MemRealtimeNotificationRepository) GetNotificationChan() chan domain.Notification {
	return r.notificationChan
}
