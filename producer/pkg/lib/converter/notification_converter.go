package converter

import (
	"github.com/blazejsewera/notipie/core/pkg/model"
	"io"
)

func FromJSON(r io.Reader) (model.AppNotification, error) {
	an, err := model.AppNotificationFromJSON(r)
	if err != nil {
		return model.AppNotification{}, err
	}
	return an, nil
}

func FromYAML(r io.Reader) (model.AppNotification, error) {
	an, err := model.AppNotificationFromYAML(r)
	if err != nil {
		return model.AppNotification{}, err
	}
	return an, nil
}
