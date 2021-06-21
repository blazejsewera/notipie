package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestString(t *testing.T) {
	notification := getTestNotification()
	str := `@2021-01-21T12:49:30Z | M | Test Notification
|> First line of body
|> Second line of body
`
	assert.Equal(t, str, notification.String())
}
