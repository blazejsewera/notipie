package grid_test

import (
	"github.com/jazzsewera/notipie/core/internal/impl/grid"
	"github.com/jazzsewera/notipie/core/internal/impl/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGrid(t *testing.T) {
	// given
	cnExpected := NewTestClientNotification()
	t.Run("send notification - receive on user proxy", func(t *testing.T) {
		// given
		g := grid.NewGrid(MockClientHubFactory{})
		an := NewTestAppNotification()
		g.Start()

		// when
		g.GetAppNotificationChan() <- an

		// then
		userProxy, _ := g.GetUserProxy(grid.RootUsername)
		cn := <-userProxy.GetClientHub().GetBroadcastChan()
		assertClientNotificationEqual(t, cnExpected, cn)
	})
}

func assertClientNotificationEqual(t testing.TB, expected model.ClientNotification, actual model.ClientNotification) {
	t.Helper()
	expected.AppID = actual.AppID
	assert.Equal(t, expected, actual)
}
