package grid_test

import (
	model "github.com/blazejsewera/notipie/core/pkg/model"
)

func newTestAppNotification() model.AppNotification {
	return model.ExampleAppNotification
}

func newTestClientNotification() model.ClientNotification {
	return model.ExampleClientNotification
}
