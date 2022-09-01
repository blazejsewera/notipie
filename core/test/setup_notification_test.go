package test

import "github.com/blazejsewera/notipie/core/pkg/model"

const timestamp = "2022-03-20T15:55:00.000Z"

var hnn = model.HashableNetNotification{
	Timestamp: timestamp,
	AppName:   "TestApp",
	Title:     "Test Title",
	Subtitle:  "Test Subtitle",
	Body:      "Test Body",
}

var appNotification = model.AppNotification{
	HashableNetNotification: hnn,
}

var clientNotification = model.ClientNotification{
	HashableNetNotification: hnn,
}
