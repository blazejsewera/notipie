package config_test

import (
	"github.com/blazejsewera/notipie/producer/pkg/lib/config"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestParser(t *testing.T) {
	t.Run("from JSON", func(t *testing.T) {
		// given
		r := strings.NewReader(exampleConfigJSON)

		// when
		actual, err := config.FromJSON(r)

		// then
		if assert.NoError(t, err) {
			assert.Equal(t, exampleConfig, actual)
		}
	})

	t.Run("from YAML", func(t *testing.T) {
		// given
		r := strings.NewReader(exampleConfigYAML)

		// when
		actual, err := config.FromYAML(r)

		// then
		if assert.NoError(t, err) {
			assert.Equal(t, exampleConfig, actual)
		}
	})
}
