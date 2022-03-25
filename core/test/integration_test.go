package test

import (
	"github.com/blazejsewera/notipie/core/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserClient(t *testing.T) {
	// given
	expectedCN := clientNotification

	t.Run("push notification - notification is pushed to ws client", func(t *testing.T) {
		// given
		initCore(t)
		ac := newAppRestClient(t)

		uWSC := newUserWSClient(t)
		uWSC.connect()
		defer uWSC.close()

		// when
		ac.pushNotification(appNotification)

		// then
		<-uWSC.saved
		actualCN := uWSC.notifications[0]
		assertClientNotification(t, expectedCN, actualCN, ac.appID)
	})

	t.Run("get notifications - notification list is returned", func(t *testing.T) {
		// given
		initCore(t)
		ac := newAppRestClient(t)

		uRC := newUserRestClient(t)

		// when
		ac.pushNotification(appNotification)

		// then
		uRC.getNotifications()
		actualCN := uRC.notifications[0]
		assertClientNotification(t, expectedCN, actualCN, ac.appID)
		// TODO: write test
	})
}

func assertClientNotification(
	t testing.TB,
	expected model.ClientNotification,
	actual model.ClientNotification,
	actualAppID string,
) {
	t.Helper()
	expected.AppID = actualAppID
	assert.Equal(t, expected, actual)
}
