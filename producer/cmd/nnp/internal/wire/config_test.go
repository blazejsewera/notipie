package wire

import (
	"github.com/blazejsewera/notipie/producer/pkg/lib/config"
	"github.com/stretchr/testify/assert"
	"testing"
	"testing/fstest"
)

func TestGetProducerNotification(t *testing.T) {
	// given
	path := "config_file.yaml"
	f := fstest.MapFS{}
	file := &fstest.MapFile{
		Data: []byte(cfgYAML),
		Mode: 0644,
		Sys:  nil,
	}
	f[path] = file

	expected := config.Config{
		EndpointConfig: config.EndpointConfig{
			Address: "localhost",
			Port:    8080,
		},
	}

	// when
	actual, err := parseConfigFromFile(f, path)

	// then
	if assert.NoError(t, err) {
		assert.Equal(t, expected, actual)
	}
}
