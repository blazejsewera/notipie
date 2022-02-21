package model_test

import (
	"github.com/jazzsewera/notipie/core/internal/impl/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNetNotification(t *testing.T) {
	// given
	netNotification := model.NetNotification{
		HashableNetNotification: model.HashableNetNotification{
			AppName: "1",
			AppID:   "2",
			Title:   "3",
		},
		Timestamp: "4",
		Read:      true,
		ApiKey:    "5",
	}
	netNotificationJSON := `{"appName":"1","appId":"2","title":"3","timestamp":"4","read":true,"apiKey":"5"}`

	t.Run("marshal json", func(t *testing.T) {
		// when
		marshaled := netNotification.ToJSON()

		// then
		assert.Equal(t, netNotificationJSON, marshaled)
	})

	t.Run("unmarshal json", func(t *testing.T) {
		t.Run("valid", func(t *testing.T) {
			// when
			unmarshaled, err := model.NetNotificationFromJSON(netNotificationJSON)

			// then
			if assert.NoError(t, err) {
				assert.Equal(t, netNotification, unmarshaled)
			}
		})

		t.Run("invalid", func(t *testing.T) {
			// given
			invalidJSON := `{"title":"1"}`

			// when
			_, err := model.NetNotificationFromJSON(invalidJSON)

			// then
			if assert.Error(t, err) {
				assert.Equal(t, model.NotEnoughInfoInNotificationErrorMessage, err.Error())
			}
		})
	})

	t.Run("add id to net notification", func(t *testing.T) {
		// given
		expectedHash := "8Mkt7MhqpOfj27kg8m6Ss+KWcwsA2IIL+Et9UBMCJUs="

		// when
		nnWithID := model.AddIDTo(netNotification)

		// then
		assert.Equal(t, expectedHash, nnWithID.ID)
	})
}
