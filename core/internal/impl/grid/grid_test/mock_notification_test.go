package grid_test

import "github.com/jazzsewera/notipie/core/internal/impl/model"

func NewTestNetNotification() model.NetNotification {
	return model.NetNotification{
		AppName:    "TestApp",
		Timestamp:  "2022-02-19T11:00:00.000Z",
		Title:      "Test NetNotification",
		Subtitle:   "Test Subtitle",
		Body:       "Test Body",
		ID:         "1",
		Read:       false,
		ExtUri:     "testExtUri",
		ReadUri:    "testReadUri",
		ArchiveUri: "testArchiveUri",
	}
}
