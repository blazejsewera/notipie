package converter_test

import (
	"github.com/blazejsewera/notipie/core/pkg/model"
	"github.com/blazejsewera/notipie/producer/pkg/lib/converter"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestConvert(t *testing.T) {
	// given
	expected := model.ExampleAppNotification

	t.Run("from JSON", func(t *testing.T) {
		// given
		appNotificationJSON := model.ExampleAppNotificationJSON
		appNotificationJSONReader := strings.NewReader(appNotificationJSON)

		// when
		actual, err := converter.FromJSON(appNotificationJSONReader)

		// then
		if assert.NoError(t, err) {
			assert.Equal(t, expected, actual)
		}
	})

	t.Run("from YAML", func(t *testing.T) {
		// given
		appNotificationYAML := model.ExampleAppNotificationYAML
		appNotificationYAMLReader := strings.NewReader(appNotificationYAML)

		// when
		actual, err := converter.FromYAML(appNotificationYAMLReader)

		// then
		if assert.NoError(t, err) {
			assert.Equal(t, expected, actual)
		}
	})
}
