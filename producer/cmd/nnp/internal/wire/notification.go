package wire

import (
	"fmt"
	"github.com/blazejsewera/notipie/core/pkg/lib/util"
	"github.com/blazejsewera/notipie/core/pkg/model"
	"os"
	"path/filepath"
)

type AppNotificationConfig struct {
	UseDefaultNotificationFile bool
	NotificationFilePath       string
	PartialNotification        model.AppNotification
}

func AppNotificationFrom(c AppNotificationConfig) (model.AppNotification, error) {
	base, err := getBaseNotification(c)
	if err != nil {
		return model.AppNotification{}, err
	}

	return mergeNotifications(base, c.PartialNotification), nil
}

func getBaseNotification(c AppNotificationConfig) (base model.AppNotification, err error) {
	base = model.AppNotification{
		HashableNetNotification: model.HashableNetNotification{
			AppName: "Example App",
			Title:   "Example Title",
		},
	}

	if c.UseDefaultNotificationFile {
		base, err = notificationFromFile(DefaultNotificationFilePath)
		if err != nil {
			return model.AppNotification{}, err
		}
	} else if c.NotificationFilePath != "" {
		base, err = notificationFromFile(c.NotificationFilePath)
		if err != nil {
			return model.AppNotification{}, err
		}
	}
	return base, nil
}

func notificationFromFile(path string) (model.AppNotification, error) {
	file, err := os.Open(path)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	if err != nil {
		return model.AppNotification{}, err
	}

	extension := filepath.Ext(path)

	var appNotification model.AppNotification

	switch extension {
	case ".yaml":
	case ".yml":
		appNotification, err = model.AppNotificationFromYAML(file)
		break
	case ".json":
		appNotification, err = model.AppNotificationFromJSON(file)
		break
	default:
		return model.AppNotification{}, fmt.Errorf("parse notification: provided file is not supported, supported files are .yml, .yaml, and .json; file path: %s", path)
	}

	if err != nil {
		return model.AppNotification{}, err
	}

	return appNotification, nil
}

func mergeNotifications(base, patch model.AppNotification) model.AppNotification {
	return util.Merge(base, patch)
}
