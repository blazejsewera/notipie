package wire

import (
	"github.com/blazejsewera/notipie/core/pkg/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeNotifications(t *testing.T) {
	// given
	base := model.AppNotification{
		HashableNetNotification: model.HashableNetNotification{
			Timestamp: "Timestamp",
			AppName:   "AppName",
			Title:     "Old Title",
		},
	}

	patch := model.AppNotification{HashableNetNotification: model.HashableNetNotification{Title: "New Title"}}

	expected := model.AppNotification{
		HashableNetNotification: model.HashableNetNotification{
			Timestamp: "Timestamp",
			AppName:   "AppName",
			Title:     "New Title",
		},
	}

	// when
	actual := mergeNotifications(base, patch)

	// then
	assert.Equal(t, expected, actual)
	assert.NotEqual(t, expected, base)
}
