package test

import (
	"github.com/blazejsewera/notipie/core/internal/model"
	"github.com/blazejsewera/notipie/core/pkg/lib/netutil"
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
	expectedCN := clientNotification

	t.Run("push notification - notification is pushed to ws client", func(t *testing.T) {
		// given
		port := init()
		ac := newAppRestClient(t, port)

		uwsc := newUserWSClient(t, port)
		uwsc.connect()
		defer uwsc.close()

		// when
		ac.pushNotification(appNotification)

		// then
		<-uwsc.saved
		actualCN := uwsc.notifications[0]
		assertClientNotification(t, expectedCN, actualCN, ac.appID)
	})

	t.Run("get notifications - notification list is returned", func(t *testing.T) {
		// given
		port := init()
		ac := newAppRestClient(t, port)

		urc := newUserRestClient(t, port)

		// when
		ac.pushNotification(appNotification)

		// then
		urc.getNotifications()
		actualCN := urc.notifications[0]
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
