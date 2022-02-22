package persistence

import "github.com/jazzsewera/notipie/core/internal/domain"

type MemNotificationRepository struct {
	notifications []domain.Notification
}

func (r *MemNotificationRepository) SaveNotification(notification domain.Notification) {
	r.notifications = append(r.notifications, notification)
}

func (r *MemNotificationRepository) GetLastNotifications(n int) []domain.Notification {
	return r.notifications[len(r.notifications)-n:]
}

func (r *MemNotificationRepository) GetNotifications(from, to int) []domain.Notification {
	return r.notifications[from:to]
}

func (r *MemNotificationRepository) GetNotificationCount() int {
	return len(r.notifications)
}
