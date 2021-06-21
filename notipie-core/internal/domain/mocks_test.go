package domain

import "time"

type MockHandler struct {
	HandledNotification Notification
	Err                 error
}

func (h *MockHandler) HandleNotification(notification Notification) error {
	h.HandledNotification = notification
	return nil
}

func (h *MockHandler) HandleError(err error) {
	h.Err = err
}

type ErroredHandler struct {
	HandledNotification Notification
	Err                 error
}

func (h *ErroredHandler) HandleNotification(_ Notification) error {
	return HandlerError{msg: "an error occurred"}
}

func (h *ErroredHandler) HandleError(err error) {
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

func getTestHandlerError() HandlerError {
	return HandlerError{msg: "an error occurred"}
}
