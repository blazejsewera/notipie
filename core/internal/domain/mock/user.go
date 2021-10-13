package mock

import (
	"github.com/jazzsewera/notipie/core/internal/domain"
	"time"
)

func NewTestUser() (*domain.User, *NotificationRepository) {
	repo := NotificationRepository{}
	return domain.NewUser("1", "TestUser", &repo), &repo
}

func NewTestUserWithAsyncRepo() (*domain.User, *AsyncNotificationRepository) {
	repo := NewAsyncNotificationRepository()
	return domain.NewUser("1", "TestUser", &repo), &repo
}

type NotificationRepository struct {
	Notifications []domain.Notification
}

func (r *NotificationRepository) GetNotificationCount() int {
	return len(r.Notifications)
}

func (r *NotificationRepository) GetNotifications(from, to int) []domain.Notification {
	return r.Notifications[from:to]
}

func (r *NotificationRepository) SaveNotification(notification domain.Notification) {
	r.Notifications = append(r.Notifications, notification)
}

func (r *NotificationRepository) GetLastNotifications(n int) []domain.Notification {
	return r.Notifications[len(r.Notifications)-n:]
}

type AsyncNotificationRepository struct {
	NotificationRepository
	NotificationSaved chan struct{}
}

func (r *AsyncNotificationRepository) SaveNotification(notification domain.Notification) {
	r.Notifications = append(r.Notifications, notification)
	select {
	case r.NotificationSaved <- struct{}{}:
		return
	case <-time.After(200 * time.Millisecond):
		panic("no receiver for r.NotificationSaved for 200ms")
	}
}

func NewAsyncNotificationRepository() AsyncNotificationRepository {
	return AsyncNotificationRepository{NotificationSaved: make(chan struct{})}
}
