package wire

import (
	"github.com/blazejsewera/notipie/producer/pkg/lib/config"
	"github.com/blazejsewera/notipie/producer/pkg/lib/converter"
	"github.com/blazejsewera/notipie/producer/pkg/lib/nnp"
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

func getProducerConfig() (nnp.ProducerConfig, error) {
	r := strings.NewReader(cfgYAML)
	c, err := config.FromYAML(r)
	if err != nil {
		return nnp.ProducerConfig{}, err
	}

	cfg, err := converter.ProducerConfigFrom(c)
	if err != nil {
		return nnp.ProducerConfig{}, err
	}

	return cfg, nil
}
