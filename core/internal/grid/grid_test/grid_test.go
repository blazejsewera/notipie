package grid_test

import (
	"github.com/blazejsewera/notipie/core/internal/grid"
	"github.com/blazejsewera/notipie/core/pkg/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGrid(t *testing.T) {
	// given
	cnExpected := newTestClientNotification()
	t.Run("send notification - receive on user proxy", func(t *testing.T) {
		// given
		g := grid.NewGrid(mockRepositoryFactory, mockBroadcasterFactory)
		an := newTestAppNotification()
		g.Start()

		// when
		appID := g.ReceiveAppNotification(an)

		// then
		<-mockBroadcasterInstance.Done
		cn := mockBroadcasterInstance.ClientNotification
		assertClientNotificationEqual(t, cnExpected, cn, appID)
	})
}

func assertClientNotificationEqual(t testing.TB, expected model.ClientNotification, actual model.ClientNotification, actualAppID string) {
	t.Helper()
	expected.AppID = actualAppID
	assert.Equal(t, expected, actual)
}
