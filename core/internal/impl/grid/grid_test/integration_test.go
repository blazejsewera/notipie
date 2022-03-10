package grid_test

import (
	"github.com/blazejsewera/notipie/core/internal/impl/grid"
	"github.com/blazejsewera/notipie/core/internal/impl/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGrid(t *testing.T) {
	// given
	cnExpected := NewTestClientNotification()
	t.Run("send notification - receive on user proxy", func(t *testing.T) {
		// given
		g := grid.NewGrid(MockClientHubFactory)
		an := NewTestAppNotification()
		g.Start()

		// when
		g.ReceiveAppNotification(an)

		// then
		appID := g.GetAppID()
		userProxy, _ := g.GetUserProxy(grid.RootUsername)
		<-userProxy.GetClientHub().(*MockClientHub).Done
		cn := userProxy.GetClientHub().(*MockClientHub).Notifications[0]
		assertClientNotificationEqual(t, cnExpected, cn, appID)
	})
}

func assertClientNotificationEqual(t testing.TB, expected model.ClientNotification, actual model.ClientNotification, actualAppID string) {
	t.Helper()
	expected.AppID = actualAppID
	assert.Equal(t, expected, actual)
}
