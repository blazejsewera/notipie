package wire

import (
	"fmt"
	"github.com/blazejsewera/notipie/core/pkg/lib/util"
	"github.com/blazejsewera/notipie/core/pkg/model"
	"github.com/blazejsewera/notipie/producer/pkg/lib/converter"
	"os"
	"path/filepath"
)

type GetAppNotificationConfig struct {
	UseDefaultNotificationFile bool
	NotificationFilePath       string
	PartialNotification        model.AppNotification
}

func GetAppNotification(c GetAppNotificationConfig) (model.AppNotification, error) {
	base, err := getBaseNotification(c)
	if err != nil {
		return model.AppNotification{}, err
	}

	return mergeNotifications(base, c.PartialNotification), nil
}

func getBaseNotification(c GetAppNotificationConfig) (base model.AppNotification, err error) {
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
		appNotification, err = converter.FromYAML(file)
		break
	case ".json":
		appNotification, err = converter.FromJSON(file)
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

var UserHomeDir, _ = os.UserConfigDir()
var DefaultNotificationFilePath = filepath.Join(UserHomeDir, "notipie", "producer", "notification.yaml")
