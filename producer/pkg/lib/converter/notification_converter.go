package converter

import (
	"github.com/blazejsewera/notipie/core/pkg/model"
	"io"
)

func FromJSON(r io.Reader) (model.AppNotification, error) {
	appNotification, err := model.AppNotificationFromReader(r)
	if err != nil {
		return model.AppNotification{}, err
	}
	return appNotification, nil
}
