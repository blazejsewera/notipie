package wire

import (
	"fmt"
	"github.com/blazejsewera/notipie/producer/pkg/lib/config"
	"github.com/blazejsewera/notipie/producer/pkg/lib/converter"
	"github.com/blazejsewera/notipie/producer/pkg/lib/nnp"
	"os"
	"path/filepath"
	"strings"
)

var producerInstance nnp.Producer

func init() {
}

func GetProducer(configPath string) nnp.Producer {
	if producerInstance != nil {
		return producerInstance
	}

	cfg, err := getProducerConfig(configPath)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "read config:", err)
		os.Exit(2)
	}

	producerInstance = nnp.NewProducer(cfg, nnp.AppIDSaverFunc(func(appID string) error {
		fmt.Println("appID:", appID)
		return nil
	}))
	return producerInstance
}

const cfgYAML = `
endpointConfig:
  address: localhost
  port: 8080
  prefix: "/"
  root: ""
  push: push
`

func getProducerConfig(path string) (nnp.ProducerConfig, error) {
	if path != "" {
		return parseConfig(path)
	} else if fileExists(DefaultProducerConfigFilePath) {
		return parseConfig(DefaultProducerConfigFilePath)
	}

	_, _ = fmt.Fprintln(os.Stderr, "warning: custom config not specified and default config does not exist, falling back to sample config")
	return sampleConfig(), nil
}

func parseConfig(path string) (nnp.ProducerConfig, error) {
	file, err := os.Open(path)
	if err != nil {
		return nnp.ProducerConfig{}, err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	extension := filepath.Ext(path)
	var cfg config.Config

	switch extension {
	case ".yaml":
	case ".yml":
		cfg, err = config.FromYAML(file)
		break
	case ".json":
		cfg, err = config.FromJSON(file)
		break
	default:
		return nnp.ProducerConfig{}, fmt.Errorf("parse config: provided file is not supported, supported files are .yml, .yaml, and .json; file path: %s", path)
	}

	if err != nil {
		return nnp.ProducerConfig{}, err
	}

	return converter.ProducerConfigFrom(cfg)
}

func sampleConfig() nnp.ProducerConfig {
	r := strings.NewReader(cfgYAML)
	c, _ := config.FromYAML(r)
	cfg, _ := converter.ProducerConfigFrom(c)
	return cfg
}
