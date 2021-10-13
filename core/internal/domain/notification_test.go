package domain

import (
	"github.com/jazzsewera/notipie/core/internal/domain/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNotification_String(t *testing.T) {
	notification := mock.NewTestNotification()
	str := `[TestApp#1@2021-01-01T00:00:00Z|M] Test Notification#1
|> First line of body
|> Second line of body
`
	assert.Equal(t, str, notification.String())
}
