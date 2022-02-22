package grid_test

import (
	"github.com/jazzsewera/notipie/core/internal/impl/grid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAppProxy(t *testing.T) {
	// given
	appNotification := NewTestAppNotification()
	appProxy := grid.NewAppProxy(grid.NewGrid())
	appProxy.Listen()

	// when
	appProxy.Receive(appNotification)

	// then
	assert.Equal(t, 1, appProxy.GetAppCount())
}
