package domain

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApp_Send(t *testing.T) {
	// given
	app := getTestApp()
	notification := getTestNotification()

	// when
	err := app.Send(notification)

	// then
	assert.NotNil(t, err)
	// TODO: refactor test to use error format constant
	assert.Equal(t, fmt.Sprintf("no tags for TestApp#1 when sending %s", notification), err.Error())
}
