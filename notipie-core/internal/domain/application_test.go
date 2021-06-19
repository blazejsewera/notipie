package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApplication_RegisterListener(t *testing.T) {
	application := NewApplication("TestApp", "", "")
	application.RegisterListener("TestID", NewListener(""))
	t.Run("register unique listener", func(t *testing.T) {
		assert.Contains(t, application.Listeners, NewListener(""))
	})
}
