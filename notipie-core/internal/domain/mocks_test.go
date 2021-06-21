package domain

import "time"

type MockAppHandler struct {
	HandledNotification Notification
	Err                 error
}

func (h *MockAppHandler) HandleNotification(notification Notification) error {
	if notification.Title == "Error" {
		return getTestAppHandlerError()
	}
	h.HandledNotification = notification
	return nil
}

func (h *MockAppHandler) HandleError(err error) {
	h.Err = err
}

type MockUserHandler struct {
	HandledNotification Notification
	Err                 error
	HandledApp          Application
}

func (h *MockUserHandler) Handle(app Application, noti Notification) error {
	if noti.Title == "Error" {
		return getTestAppHandlerError()
	}
	h.HandledNotification = noti
	h.HandledApp = app
	return nil
}

func (h *MockUserHandler) HandleError(err error) {
	h.Err = err
}

func getTestNotification() Notification {
	timestamp, _ := time.Parse(time.RFC3339, "2021-01-21T12:49:30Z")
	return Notification{
		Timestamp: timestamp,
		Title:     "Test Notification",
		Body:      "First line of body\nSecond line of body",
		Urgency:   Medium,
	}
}

// getTestErrNotification returns a valid Notification,
// but it triggers an error in tests to reduce code repetition.
func getTestErrNotification() Notification {
	return Notification{Title: "Error"}
}

func getTestAppHandlerError() AppHandlerError {
	return AppHandlerError{msg: "an error occurred"}
}

func getTestUserHandlerError() UserHandlerError {
	return UserHandlerError{msg: "an error occurred"}
}
