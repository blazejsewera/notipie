package domain_test

import (
	"github.com/blazejsewera/notipie/core/internal/domain"
	"time"
)

func NewTestUser() (*domain.User, *MockNotificationRepository) {
	repo := &MockNotificationRepository{}
	return domain.NewUser("1", "TestUser", repo), repo
}

func NewTestUserWithAsyncRepo() (*domain.User, *MockAsyncNotificationRepository) {
	repo := NewAsyncNotificationRepository()
	return domain.NewUser("1", "TestUser", repo), repo
}

type MockNotificationRepository struct {
	Notifications []domain.Notification
}

func (r *MockNotificationRepository) GetNotificationCount() int {
	return len(r.Notifications)
}

func (r *MockNotificationRepository) GetNotifications(from, to int) []domain.Notification {
	return r.Notifications[from:to]
}

func (r *MockNotificationRepository) SaveNotification(notification domain.Notification) {
	r.Notifications = append(r.Notifications, notification)
}

func (r *MockNotificationRepository) GetLastNotifications(n int) []domain.Notification {
	return r.Notifications[len(r.Notifications)-n:]
}

func GetAllNotificationsFor(u *domain.User) []domain.Notification {
	return u.GetNotifications(0, u.GetNotificationCount())
}

type MockAsyncNotificationRepository struct {
	MockNotificationRepository
	NotificationSaved chan struct{}
}

func (r *MockAsyncNotificationRepository) SaveNotification(notification domain.Notification) {
	r.Notifications = append(r.Notifications, notification)
	select {
	case r.NotificationSaved <- struct{}{}:
		return
	case <-time.After(200 * time.Millisecond):
		panic("no receiver for r.NotificationSaved for 200ms")
	}
}

func NewAsyncNotificationRepository() *MockAsyncNotificationRepository {
	return &MockAsyncNotificationRepository{NotificationSaved: make(chan struct{})}
}
