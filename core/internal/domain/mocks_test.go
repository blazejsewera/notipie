package domain

import "time"

type MockAppHandler struct {
	HandledNotification Notification
}

func (h *MockAppHandler) HandleNotification(notification Notification) {
	h.HandledNotification = notification
}

type MockUserHandler struct {
	HandledNotification Notification
	HandledApp          App
}

func (h *MockUserHandler) Handle(app App, noti Notification) {
	h.HandledNotification = noti
	h.HandledApp = app
}

type MockUserNotificationRepository struct {
	// TODO: Add notifications slice
}

func (r *MockUserNotificationRepository) SaveNotification(notification Notification) {
}

func (r *MockUserNotificationRepository) GetAllNotifications() []Notification {
	return nil
}

func getTestNotification() Notification {
	app := App{
		ID:           "1",
		Name:         "TestApp",
		SmallIconURL: ".",
		BigIconURL:   ".",
	}

	timestamp, _ := time.Parse(time.RFC3339, "2021-01-21T12:49:30Z")
	return Notification{
		App:       &app,
		Timestamp: timestamp,
		Title:     "Test Notification",
		Body:      "First line of body\nSecond line of body",
		Urgency:   Medium,
	}
}
