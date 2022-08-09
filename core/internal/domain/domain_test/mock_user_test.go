package domain_test

import (
	"github.com/blazejsewera/notipie/core/internal/domain"
	"github.com/blazejsewera/notipie/core/pkg/lib/util"
	"time"
)

func NewTestUser() (*domain.User, *MockNotificationRepository) {
	repo := &MockNotificationRepository{}
	broadcaster := &MockNotificationBroadcaster{}
	return domain.NewUser("1", "TestUser", repo, broadcaster), repo
}

func NewTestUserWithAsyncRepo() (*domain.User, *MockAsyncNotificationRepository) {
	repo := NewAsyncNotificationRepository()
	broadcaster := &MockNotificationBroadcaster{}
	return domain.NewUser("1", "TestUser", repo, broadcaster), repo
}

type MockNotificationBroadcaster struct{}

// @impl
var _ domain.NotificationBroadcaster = (*MockNotificationBroadcaster)(nil)

func (b *MockNotificationBroadcaster) Broadcast(domain.Notification) {}

func (b *MockNotificationBroadcaster) RegisterClient(interface{}) {}

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
	NotificationSaved chan util.Signal
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
	return &MockAsyncNotificationRepository{NotificationSaved: make(chan util.Signal)}
}
