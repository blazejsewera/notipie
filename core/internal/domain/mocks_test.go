package domain

import (
	"fmt"
	"time"
)

type MockNotificationRepository struct {
	Notifications []Notification
}

func (r *MockNotificationRepository) SaveNotification(notification Notification) {
	r.Notifications = append(r.Notifications, notification)
}

func (r *MockNotificationRepository) GetAllNotifications() []Notification {
	return r.Notifications
}

func (r *MockNotificationRepository) GetLastNotifications(n int) []Notification {
	return r.Notifications[3:4]
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
		repo:     &MockNotificationRepository{},
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
