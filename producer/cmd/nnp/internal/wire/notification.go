package wire

import "github.com/blazejsewera/notipie/core/pkg/model"

var exampleAppNotification = model.AppNotification{
	HashableNetNotification: model.HashableNetNotification{
		AppName:    "Example App Name",
		AppImgURI:  "https://www.sewera.dev/magpie_dark.svg",
		Title:      "Example Notification",
		Subtitle:   "Subtitle",
		Body:       "Body",
		ExtURI:     "https://www.sewera.dev/",
		ReadURI:    "https://www.sewera.dev/",
		ArchiveURI: "https://www.sewera.dev/",
	},
}

func GetAppNotification() (model.AppNotification, error) {
	return exampleAppNotification, nil
}
