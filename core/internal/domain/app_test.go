package domain

import "testing"

func TestSendNotification(t *testing.T) {
	// given
	notification := getTestNotification()
	app := App{}

	// when
	app.Send(notification)
}
