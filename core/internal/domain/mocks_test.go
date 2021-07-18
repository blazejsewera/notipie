package domain

import (
	"fmt"
	"time"
)

type mockNotificationRepository struct {
	Notifications     []Notification
	NotificationSaved chan struct{}
}

func newMockNotificationRepository() mockNotificationRepository {
	return mockNotificationRepository{NotificationSaved: make(chan struct{})}
}

func (r *mockNotificationRepository) GetNotificationCount() int {
	return len(r.Notifications)
}

func (r *mockNotificationRepository) GetNotifications(from, to int) []Notification {
	return r.Notifications[from:to]
}

func (r *mockNotificationRepository) SaveNotification(notification Notification) {
	r.Notifications = append(r.Notifications, notification)
	select {
	case r.NotificationSaved <- struct{}{}:
		return
	case <-time.After(200 * time.Millisecond):
		panic("no receiver for r.NotificationSaved for 200ms")
	}
}

func (r *mockNotificationRepository) GetLastNotifications(n int) []Notification {
	return r.Notifications[len(r.Notifications)-n:]
}

type mockCommandHandler struct {
	Command Command
}

func (h *mockCommandHandler) HandleCommand(command Command) {
	h.Command = command
}

func getTestNotification() Notification {
	app := getTestApp()

	timestamp, _ := time.Parse(time.RFC3339, "2021-01-01T00:00:00Z")
	return Notification{
		ID:        "1",
		App:       &app,
		Timestamp: timestamp,
		Title:     "Test Notification",
		Body:      "First line of body\nSecond line of body",
		Urgency:   Medium,
	}
}

func getTestApp() App {
	return App{
		ID:   "1",
		Name: "TestApp",
	}
}

func getTestUser() User {
	return User{
		ID:       "1",
		Username: "TestUser",
	}
}

func getTestTag() Tag {
	return Tag{
		Name: "TestTag",
	}
}

func get5TestNotifications() (notifications []Notification) {
	for i := 1; i <= 5; i++ {
		app := &App{
			ID:   fmt.Sprint(i),
			Name: fmt.Sprintf("TestApp%d", i),
		}
		timestamp, _ := time.Parse(time.RFC3339, fmt.Sprintf("2021-01-0%dT0%d:0%d:0%dZ", i, i, i, i))
		notifications = append(notifications, Notification{
			ID:        fmt.Sprint(i),
			App:       app,
			Timestamp: timestamp,
			Title:     fmt.Sprintf("Test Notification #%d", i),
			Body:      fmt.Sprint(i),
			Urgency:   Urgency(i - 1),
		})
	}
	return
}
