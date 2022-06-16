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
	invalidJSON := `{"title":"1"}`
	invalidJSONReader := strings.NewReader(invalidJSON)

	appNotificationYAML := model.ExampleAppNotificationYAML
	appNotificationYAMLReader := strings.NewReader(appNotificationYAML)
	invalidYAML := `
title: '1'
`
	invalidYAMLReader := strings.NewReader(invalidYAML)

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
		t.Run("valid", func(t *testing.T) {
			// when
			unmarshaled, err := model.AppNotificationFromJSON(appNotificationJSONReader)

			// then
			if assert.NoError(t, err) {
				assert.Equal(t, appNotification, unmarshaled)
			}
		})

		t.Run("invalid", func(t *testing.T) {
			// when
			_, err := model.AppNotificationFromJSON(invalidJSONReader)

			// then
			if assert.Error(t, err) {
				assert.Equal(t, model.NotEnoughInfoInNotificationErrorMessage, err.Error())
			}
		})
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
		t.Run("valid", func(t *testing.T) {
			// when
			unmarshaled, err := model.AppNotificationFromYAML(appNotificationYAMLReader)

			// then
			if assert.NoError(t, err) {
				assert.Equal(t, appNotification, unmarshaled)
			}
		})

		t.Run("invalid", func(t *testing.T) {
			// when
			_, err := model.AppNotificationFromYAML(invalidYAMLReader)

			// then
			if assert.Error(t, err) {
				assert.Equal(t, model.NotEnoughInfoInNotificationErrorMessage, err.Error())
			}
		})
	})

	t.Run("add id to net notification", func(t *testing.T) {
		// given
		expectedHash := "frGOwBO6bNL/kbixYn3eJ6xS8WAewHK7qzt8q1cLVLs="
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
