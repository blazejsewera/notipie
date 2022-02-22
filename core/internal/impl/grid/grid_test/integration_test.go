package grid_test

import (
	"github.com/jazzsewera/notipie/core/internal/impl/grid"
	"github.com/jazzsewera/notipie/core/pkg/lib/util"
	"testing"
)

func TestAppToGrid(t *testing.T) {
	// given
	mockGrid := NewMockGrid()
	appProxy := grid.NewAppProxy(mockGrid)
	appNotification := NewTestAppNotification()
	mockGrid.Start()
	appProxy.Listen()

	// when
	util.AsyncRun(func() {
		appProxy.AppNotificationChan <- appNotification
	})

	// then
	done := mockGrid.MockUserRepository.Done
	util.AsyncAssert(t, done).Equal(1, mockGrid.MockUserRepository.GetNotificationCount())
}
