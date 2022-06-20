package test

import (
	"github.com/blazejsewera/notipie/core/pkg/lib/fp"
	"github.com/blazejsewera/notipie/core/pkg/lib/netutil"
	"github.com/blazejsewera/notipie/core/pkg/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserClient(t *testing.T) {
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
		an := appNotification
		pushNotificationTitle := "push notification"
		an.Title = pushNotificationTitle

		ac := newAppRestClient(t, port)

		uwsc := newUserWSClient(t, port)
		uwsc.connect()
		defer uwsc.close()

		// when
		ac.pushNotification(an)

		// then
		<-uwsc.saved
		actualCN := uwsc.notifications[0]
		assertClientNotification(t, pushNotificationTitle, actualCN)
	})

	t.Run("get notifications - notification list is returned", func(t *testing.T) {
		// given
		an := appNotification
		getNotificationsTitle := "get notifications"
		an.Title = getNotificationsTitle

		ac := newAppRestClient(t, port)

		urc := newUserRestClient(t, port)

		// when
		ac.pushNotification(an)
		urc.getNotifications()

		// then
		assertContainsCN(t, urc.notifications, getNotificationsTitle)
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

func assertContainsCN(
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
