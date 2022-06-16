package wire

import "github.com/blazejsewera/notipie/core/pkg/model"

var exampleAppNotification = model.ExampleAppNotification

func GetAppNotification() (model.AppNotification, error) {
	return exampleAppNotification, nil
}
