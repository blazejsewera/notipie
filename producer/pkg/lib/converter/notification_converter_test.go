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
		appNotificationJson := model.ExampleAppNotificationJSON
		appNotificationJsonReader := strings.NewReader(appNotificationJson)

		// when
		actual, err := converter.FromJSON(appNotificationJsonReader)

		// then
		if assert.NoError(t, err) {
			assert.Equal(t, expected, actual)
		}
	})

	t.Run("from args", func(t *testing.T) {
		// TODO: write test and impl
	})
}
