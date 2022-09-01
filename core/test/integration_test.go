package test

import (
	"testing"

	"github.com/blazejsewera/notipie/core/pkg/lib/fp"
	"github.com/blazejsewera/notipie/core/pkg/lib/netutil"
	"github.com/blazejsewera/notipie/core/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestNotificationFlow(t *testing.T) {
	init := func() int {
		port, err := netutil.FindFreePort()
		if err != nil {
			panic("could not find free port")
		}
		initCore(t, port)
		return port
	}

	// given
	port := init()

	t.Run("push notification - notification is pushed to ws client", func(t *testing.T) {
		// given
		notification := appNotification
		pushNotificationTitle := "push notification"
		notification.Title = pushNotificationTitle

		producer := newAppProducer(t, port)

		userWSClient := newUserWSClient(t, port)
		userWSClient.connect()
		defer userWSClient.close()

		// when
		producer.pushNotification(notification)

		// then
		<-userWSClient.saved
		actual := userWSClient.notifications[0]
		assertClientNotification(t, pushNotificationTitle, actual)
	})

	t.Run("get notifications - notification list is returned", func(t *testing.T) {
		// given
		notification := appNotification
		getNotificationsTitle := "get notifications"
		notification.Title = getNotificationsTitle

		producer := newAppProducer(t, port)

		userRestClient := newUserRestClient(t, port)

		// when
		producer.pushNotification(notification)
		userRestClient.getNotifications()

		// then
		assertContainsClientNotification(t, userRestClient.notifications, getNotificationsTitle)
	})
}

func assertClientNotification(
	t testing.TB,
	expectedTitle string,
	actual model.ClientNotification,
) {
	t.Helper()
	assert.Equal(t, expectedTitle, actual.Title)
}

func assertContainsClientNotification(
	t testing.TB,
	notifications []model.ClientNotification,
	expectedTitle string,
) {
	t.Helper()

	notificationTitles := fp.Map(func(n model.ClientNotification) string {
		return n.Title
	}, notifications)

	assert.Contains(t, notificationTitles, expectedTitle)
}
