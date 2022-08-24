package model_test

import (
	"github.com/blazejsewera/notipie/core/pkg/model"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestAppNotification(t *testing.T) {
	// given
	appNotification := model.ExampleAppNotification
	appNotificationJSON := appNotificationJSONWithoutWhitespace()
	appNotificationJSONReader := strings.NewReader(appNotificationJSON)

	appNotificationYAML := model.ExampleAppNotificationYAML
	appNotificationYAMLReader := strings.NewReader(appNotificationYAML)

	t.Run("marshal json", func(t *testing.T) {
		// when
		marshaled, err := appNotification.ToJSON()
		actual := string(marshaled)

		// then
		if assert.NoError(t, err) {
			assert.Equal(t, appNotificationJSON, actual)
		}
	})

	t.Run("unmarshal json", func(t *testing.T) {
		// when
		unmarshaled, err := model.AppNotificationFromJSON(appNotificationJSONReader)

		// then
		if assert.NoError(t, err) {
			assert.Equal(t, appNotification, unmarshaled)
		}
	})

	t.Run("marshal yaml", func(t *testing.T) {
		// when
		marshaled, err := appNotification.ToYAML()
		actual := string(marshaled)

		// then
		if assert.NoError(t, err) {
			assert.Equal(t, appNotificationYAML, actual)
		}
	})

	t.Run("unmarshal yaml", func(t *testing.T) {
		// when
		unmarshaled, err := model.AppNotificationFromYAML(appNotificationYAMLReader)

		// then
		if assert.NoError(t, err) {
			assert.Equal(t, appNotification, unmarshaled)
		}
	})

	t.Run("add id to net notification", func(t *testing.T) {
		// given
		expectedHash := "j/M0l0NIdJHGvAFTKY2uc7EW4+FiN3jDytCoPJL84ZM="
		anWithoutID := appNotification
		anWithoutID.ID = ""

		// when
		anWithID := model.AddIDTo(anWithoutID)

		// then
		assert.Equal(t, expectedHash, anWithID.ID)
	})
}

func appNotificationJSONWithoutWhitespace() string {
	appNotificationJSON := model.ExampleAppNotificationJSON
	appNotificationJSON = strings.ReplaceAll(appNotificationJSON, " ", "")
	appNotificationJSON = strings.ReplaceAll(appNotificationJSON, "\t", "")
	appNotificationJSON = strings.ReplaceAll(appNotificationJSON, "\n", "")
	return appNotificationJSON
}
