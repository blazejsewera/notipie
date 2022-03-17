package test

import "testing"

func TestNotipieCore(t *testing.T) {
	cfg := initCore(t)
	c := initRestClient(t, cfg)
	appNotification := appNotification

	t.Run("push notification", func(t *testing.T) {
		c.pushNotification(appNotification)
	})
}
