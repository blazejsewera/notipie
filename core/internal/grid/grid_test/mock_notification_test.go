package grid_test

import (
	model2 "github.com/blazejsewera/notipie/core/pkg/model"
)

func newTestAppNotification() model2.AppNotification {
	return model2.AppNotification{
		HashableNetNotification: model2.HashableNetNotification{
			AppName:    "TestApp",
			Title:      "Test AppNotification",
			Subtitle:   "Test Subtitle",
			Body:       "Test Body",
			ExtURI:     "testExtUri",
			ReadURI:    "testReadUri",
			ArchiveURI: "testArchiveUri",
		},
		Timestamp: "2022-02-19T11:00:00.000Z",
		ID:        "1",
		Read:      false,
	}
}

func newTestClientNotification() model2.ClientNotification {
	an := newTestAppNotification()
	return model2.ClientNotification{
		HashableNetNotification: an.HashableNetNotification,
		ID:                      an.ID,
		Timestamp:               an.Timestamp,
		Read:                    an.Read,
	}
}
