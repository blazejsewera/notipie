package wire

import (
	"errors"
	"os"
	"path/filepath"
)

const (
	NotipieConfigDir          = "notipie"
	ProducerConfigDir         = "producer"
	DefaultProducerConfigFile = "config.yaml"
	DefaultNotificationFile   = "notification.yaml"
)

var UserConfigDir, _ = os.UserConfigDir()
var NotipieProducerConfigDir = filepath.Join(UserConfigDir, NotipieConfigDir, ProducerConfigDir)
var DefaultProducerConfigFilePath = filepath.Join(NotipieProducerConfigDir, DefaultProducerConfigFile)
var DefaultNotificationFilePath = filepath.Join(NotipieProducerConfigDir, DefaultNotificationFile)

func MkConfigDirIfDoesNotExist() error {
	if !fileExists(NotipieProducerConfigDir) {
		return os.MkdirAll(NotipieProducerConfigDir, 0755)
	}
	return nil
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}
