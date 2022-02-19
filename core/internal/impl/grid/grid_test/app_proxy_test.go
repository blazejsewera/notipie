package grid_test

import (
	"github.com/jazzsewera/notipie/core/internal/impl/grid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAppProxy(t *testing.T) {
	// given
	netNotification := NewTestNetNotification()
	appProxy := grid.NewAppProxy(grid.NewGrid())
	appProxy.Listen()

	// when
	appProxy.Receive(netNotification)

	// then
	assert.Equal(t, 1, appProxy.GetAppCount())
}
