package model_test

import (
	"bytes"
	"github.com/blazejsewera/notipie/core/pkg/model"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestAppNotification(t *testing.T) {
	// given
	appNotification := model.ExampleAppNotification
	appNotificationJSON := appNotificationJSONWithoutWhitespace()
	appNotificationJSONReader := bytes.NewReader(appNotificationJSON)
	invalidJSON := []byte(`{"title":"1"}`)
	invalidJSONReader := bytes.NewReader(invalidJSON)

	t.Run("marshal json", func(t *testing.T) {
		// when
		marshaled, err := appNotification.ToJSON()

		// then
		if assert.NoError(t, err) {
			assert.Equal(t, appNotificationJSON, marshaled)
		}
	})

	t.Run("unmarshal json", func(t *testing.T) {
		t.Run("valid from string", func(t *testing.T) {
			// when
			unmarshaled, err := model.AppNotificationFromJSON(appNotificationJSON)

			// then
			if assert.NoError(t, err) {
				assert.Equal(t, appNotification, unmarshaled)
			}
		})

		t.Run("valid from reader", func(t *testing.T) {
			// when
			unmarshaled, err := model.AppNotificationFromReader(appNotificationJSONReader)

			// then
			if assert.NoError(t, err) {
				assert.Equal(t, appNotification, unmarshaled)
			}
		})

		t.Run("invalid from string", func(t *testing.T) {
			// when
			_, err := model.AppNotificationFromJSON(invalidJSON)

			// then
			if assert.Error(t, err) {
				assert.Equal(t, model.NotEnoughInfoInNotificationErrorMessage, err.Error())
			}
		})

		t.Run("invalid from reader", func(t *testing.T) {
			// when
			_, err := model.AppNotificationFromReader(invalidJSONReader)

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

func appNotificationJSONWithoutWhitespace() []byte {
	appNotificationJSON := model.ExampleAppNotificationJSON
	appNotificationJSON = strings.ReplaceAll(appNotificationJSON, " ", "")
	appNotificationJSON = strings.ReplaceAll(appNotificationJSON, "\t", "")
	appNotificationJSON = strings.ReplaceAll(appNotificationJSON, "\n", "")
	return []byte(appNotificationJSON)
}
