package persistence

import "github.com/blazejsewera/notipie/core/internal/domain"

type RealtimeNotificationRepo interface {
	GetNotificationChan() chan domain.Notification
	domain.NotificationRepository
}

type MemRealtimeNotificationRepository struct {
	notifications    []domain.Notification
	notificationChan chan domain.Notification
}

func NewMemRealtimeNotificationRepository() *MemRealtimeNotificationRepository {
	return &MemRealtimeNotificationRepository{notificationChan: make(chan domain.Notification)}
}

func (r *MemRealtimeNotificationRepository) SaveNotification(notification domain.Notification) {
	r.notifications = append(r.notifications, notification)
	r.notificationChan <- notification
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
