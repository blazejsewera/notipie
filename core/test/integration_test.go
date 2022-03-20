package test

import (
	"github.com/blazejsewera/notipie/core/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNotipieCore(t *testing.T) {
	// given
	initCore(t)
	ac := newAppRestClient(t)
	wsc := newWSClient(t)
	wsc.connect()
	defer wsc.close()

	t.Run("push notification", func(t *testing.T) {
		// when
		ac.pushNotification(appNotification)

		// then
		<-wsc.saved
		actualClientNotification := wsc.notifications[0]
		assertClientNotification(t, clientNotification, actualClientNotification, ac.appID)
	})

	t.Run("get notifications", func(t *testing.T) {
		// when
		// TODO: write test
	})
}

func assertClientNotification(t testing.TB, expected model.ClientNotification, actual model.ClientNotification, actualAppID string) {
	t.Helper()
	expected.AppID = actualAppID
	assert.Equal(t, expected, actual)
}
