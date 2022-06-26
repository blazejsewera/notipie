package wire

import (
	"fmt"
	"github.com/blazejsewera/notipie/core/pkg/lib/util"
	"github.com/blazejsewera/notipie/producer/pkg/lib/config"
	"github.com/blazejsewera/notipie/producer/pkg/lib/converter"
	"github.com/blazejsewera/notipie/producer/pkg/lib/nnp"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const cfgYAML = `
endpointConfig:
  address: localhost
  port: 8080
`

var (
	ConfigPath string
	BaseConfig config.Config
)

var (
	unsupportedFileError = fmt.Errorf("parse config: provided file is not supported, supported files are .yml, .yaml, and .json; file path: %s", ConfigPath)
	emptyConfigPathError = fmt.Errorf("empty config path")
)

var AppIDSaver = nnp.AppIDSaverFunc(func(appID string) error {
	if ConfigPath == "" {
		fmt.Println("appID:", appID)
		return nil
	}

	cfg := BaseConfig
	cfg.AppID = appID
	return SaveConfig(cfg)
})

func SetConfigPath(customPath string) {
	if customPath != "" {
		ConfigPath = customPath
	} else if fileExists(DefaultProducerConfigFilePath) {
		ConfigPath = DefaultProducerConfigFilePath
	} else {
		fmt.Fprintln(os.Stderr, "warning: custom config not specified and default config does not exist, falling back to sample config")
	}
}

func GetProducerConfig(patch config.Config) (nnp.ProducerConfig, error) {
	var err error
	BaseConfig, err = getBaseConfig()
	if err != nil {
		return nnp.ProducerConfig{}, err
	}
	merged := mergeConfig(BaseConfig, patch)
	return converter.ProducerConfigFrom(merged)
}

func SaveConfig(cfg config.Config) error {
	if ConfigPath == "" {
		return emptyConfigPathError
	}

	ext := filepath.Ext(ConfigPath)
	var bytes []byte
	var err error

	switch ext {
	case ".yaml", ".yml":
		bytes, err = cfg.ToYAML()
	case ".json":
		bytes, err = cfg.ToJSON()
	default:
		return unsupportedFileError
	}
	if err != nil {
		return err
	}

	return ioutil.WriteFile(ConfigPath, bytes, 0644)
}

func ConfigOf(addr string, port int, appID string) config.Config {
	return config.Config{
		EndpointConfig: config.EndpointConfig{
			Address: addr,
			Port:    port,
		},
		AppID: appID,
	}
}

func mergeConfig(base, patch config.Config) config.Config {
	return util.Merge(base, patch)
}

func getBaseConfig() (config.Config, error) {
	if ConfigPath == "" {
		return sampleConfig(), nil
	}

	file, err := os.Open(ConfigPath)
	defer file.Close()
	if err != nil {
		return config.Config{}, err
	}

	return parseConfig(file)
}

func parseConfig(r io.Reader) (config.Config, error) {
	ext := filepath.Ext(ConfigPath)

	switch ext {
	case ".yaml", ".yml":
		return config.FromYAML(r)
	case ".json":
		return config.FromJSON(r)
	default:
		return config.Config{}, unsupportedFileError
	}
}

func sampleConfig() config.Config {
	r := strings.NewReader(cfgYAML)
	c, _ := config.FromYAML(r)
	return c
}
