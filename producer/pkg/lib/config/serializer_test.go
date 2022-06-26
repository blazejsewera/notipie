package config_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSerializer(t *testing.T) {
	t.Run("to JSON", func(t *testing.T) {
		// when
		actual, err := exampleConfig.ToJSON()

		// then
		if assert.NoError(t, err) {
			assert.Equal(t, exampleConfigJSON, string(actual))
		}
	})

	t.Run("to YAML", func(t *testing.T) {
		// when
		actual, err := exampleConfig.ToYAML()

		// then
		if assert.NoError(t, err) {
			assert.Equal(t, exampleConfigYAML, string(actual))
		}
	})
}
