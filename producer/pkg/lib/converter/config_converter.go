package converter

import (
	"errors"
	"github.com/blazejsewera/notipie/core/pkg/lib/api"
	"github.com/blazejsewera/notipie/producer/pkg/lib/config"
	"github.com/blazejsewera/notipie/producer/pkg/lib/nnp"
)

func ProducerConfigFrom(serializable config.Config) (nnp.ProducerConfig, error) {
	if serializable.Address == "" {
		return nnp.ProducerConfig{}, errors.New("convert config: address was empty")
	}

	host := api.GetHost(serializable.Address, serializable.Port)

	return nnp.ProducerConfig{
		AppID: serializable.AppID,
		Endpoint: nnp.ProducerEndpointConfig{
			RootURL: api.GetURL(host, api.Root),
			PushURL: api.GetURL(host, api.Push),
		},
	}, nil
}
