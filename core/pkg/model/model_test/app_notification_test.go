package model_test

import (
	model2 "github.com/blazejsewera/notipie/core/pkg/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAppNotification(t *testing.T) {
	// given
	appNotification := model2.AppNotification{
		HashableNetNotification: model2.HashableNetNotification{
			AppName: "1",
			AppID:   "2",
			Title:   "3",
		},
		Timestamp: "4",
		Read:      true,
		ApiKey:    "5",
	}
	appNotificationJSON := `{"appName":"1","appId":"2","title":"3","timestamp":"4","read":true,"apiKey":"5"}`

	t.Run("marshal json", func(t *testing.T) {
		// when
		marshaled := appNotification.ToJSON()

		// then
		assert.Equal(t, appNotificationJSON, marshaled)
	})

	t.Run("unmarshal json", func(t *testing.T) {
		t.Run("valid", func(t *testing.T) {
			// when
			unmarshaled, err := model2.AppNotificationFromJSON(appNotificationJSON)

			// then
			if assert.NoError(t, err) {
				assert.Equal(t, appNotification, unmarshaled)
			}
		})

		t.Run("invalid", func(t *testing.T) {
			// given
			invalidJSON := `{"title":"1"}`

			// when
			_, err := model2.AppNotificationFromJSON(invalidJSON)

			// then
			if assert.Error(t, err) {
				assert.Equal(t, model2.NotEnoughInfoInNotificationErrorMessage, err.Error())
			}
		})
	})

	t.Run("add id to net notification", func(t *testing.T) {
		// given
		expectedHash := "8Mkt7MhqpOfj27kg8m6Ss+KWcwsA2IIL+Et9UBMCJUs="

		// when
		anWithID := model2.AddIDTo(appNotification)

		// then
		assert.Equal(t, expectedHash, anWithID.ID)
	})
}
