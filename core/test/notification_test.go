package test

import (
	"github.com/blazejsewera/notipie/core/internal/model"
	"github.com/blazejsewera/notipie/core/pkg/lib/timeformat"
	"time"
)

var appNotification = model.AppNotification{
	HashableNetNotification: model.HashableNetNotification{
		AppName:  "TestApp",
		Title:    "Test Title",
		Subtitle: "Test Subtitle",
		Body:     "Test Body",
	},
	Timestamp: time.Now().Format(timeformat.RFC3339Milli),
}
