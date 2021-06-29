package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestString(t *testing.T) {
	notification := getTestNotification()
	str := `[TestApp#1@2021-01-01T00:00:00Z|M] Test Notification
|> First line of body
|> Second line of body
`
	assert.Equal(t, str, notification.String())
}
