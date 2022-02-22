package grid_test

import "github.com/jazzsewera/notipie/core/internal/impl/model"

func NewTestNetNotification() model.AppNotification {
	return model.AppNotification{
		HashableNetNotification: model.HashableNetNotification{
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
