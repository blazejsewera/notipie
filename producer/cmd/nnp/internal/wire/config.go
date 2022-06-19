package wire

import (
	"fmt"
	"github.com/blazejsewera/notipie/core/pkg/lib/util"
	"github.com/blazejsewera/notipie/producer/pkg/lib/config"
	"github.com/blazejsewera/notipie/producer/pkg/lib/converter"
	"github.com/blazejsewera/notipie/producer/pkg/lib/nnp"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

const cfgYAML = `
endpointConfig:
  address: localhost
  port: 8080
  prefix: "/"
  root: ""
  push: push
`

func GetProducerConfig(basePath string, patch config.Config) (nnp.ProducerConfig, error) {
	base, err := getBaseConfig(basePath)
	if err != nil {
		return nnp.ProducerConfig{}, err
	}
	merged := mergeConfig(base, patch)
	return converter.ProducerConfigFrom(merged)
}

func PatchConfigOf(addr string, port int, appID string) config.Config {
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

func getBaseConfig(path string) (config.Config, error) {
	f := new(OsFilesystem)
	if path != "" {
		return parseConfigFromFile(f, path)
	} else if fileExists(DefaultProducerConfigFilePath) {
		return parseConfigFromFile(f, DefaultProducerConfigFilePath)
	}

	_, _ = fmt.Fprintln(os.Stderr, "warning: custom config not specified and default config does not exist, falling back to sample config")
	return sampleConfig(), nil
}

func parseConfigFromFile(f fs.FS, path string) (config.Config, error) {
	file, err := f.Open(path)
	if err != nil {
		return config.Config{}, err
	}
	defer func(file fs.File) {
		_ = file.Close()
	}(file)

	extension := filepath.Ext(path)
	var cfg config.Config

	switch extension {
	case ".yaml", ".yml":
		cfg, err = config.FromYAML(file)
	case ".json":
		cfg, err = config.FromJSON(file)
	default:
		return config.Config{}, fmt.Errorf("parse config: provided file is not supported, supported files are .yml, .yaml, and .json; file path: %s", path)
	}

	if err != nil {
		return config.Config{}, err
	}

	return cfg, nil
}

func sampleConfig() config.Config {
	r := strings.NewReader(cfgYAML)
	c, _ := config.FromYAML(r)
	return c
}

type OsFilesystem struct{}

func (f *OsFilesystem) Open(name string) (fs.File, error) {
	return os.Open(name)
}

//@impl
var _ fs.FS = new(OsFilesystem)
