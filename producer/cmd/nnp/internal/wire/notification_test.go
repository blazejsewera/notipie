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
			AppName: "AppName",
			Title:   "Old Title",
		},
		Timestamp: "Timestamp",
	}

	patch := model.AppNotification{HashableNetNotification: model.HashableNetNotification{Title: "New Title"}}

	expected := model.AppNotification{
		HashableNetNotification: model.HashableNetNotification{
			AppName: "AppName",
			Title:   "New Title",
		},
		Timestamp: "Timestamp",
	}

	// when
	actual := mergeNotifications(base, patch)

	// then
	assert.Equal(t, expected, actual)
	assert.NotEqual(t, expected, base)
}
