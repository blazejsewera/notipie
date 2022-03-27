package test

import (
	"github.com/blazejsewera/notipie/core/internal/model"
)

const timestamp = "2022-03-20T15:55:00.000Z"

var hnn = model.HashableNetNotification{
	AppName:  "TestApp",
	Title:    "Test Title",
	Subtitle: "Test Subtitle",
	Body:     "Test Body",
}

var appNotification = model.AppNotification{
	HashableNetNotification: hnn,
	Timestamp:               timestamp,
	ID:                      "1",
}

var clientNotification = model.ClientNotification{
	HashableNetNotification: hnn,
	Timestamp:               timestamp,
	ID:                      "1",
}