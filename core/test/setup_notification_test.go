package test

import (
	model2 "github.com/blazejsewera/notipie/core/pkg/model"
)

const timestamp = "2022-03-20T15:55:00.000Z"

var hnn = model2.HashableNetNotification{
	AppName:  "TestApp",
	Title:    "Test Title",
	Subtitle: "Test Subtitle",
	Body:     "Test Body",
}

var appNotification = model2.AppNotification{
	HashableNetNotification: hnn,
	Timestamp:               timestamp,
	ID:                      "1",
}

var clientNotification = model2.ClientNotification{
	HashableNetNotification: hnn,
	Timestamp:               timestamp,
	ID:                      "1",
}
