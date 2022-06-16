package converter_test

import (
	"github.com/blazejsewera/notipie/producer/pkg/lib/config"
	"github.com/blazejsewera/notipie/producer/pkg/lib/converter"
	"github.com/blazejsewera/notipie/producer/pkg/lib/nnp"
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

func TestConfigConverter(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		// given
		source := config.Config{
			EndpointConfig: config.EndpointConfig{
				Address: "localhost",
				Port:    1234,
				Prefix:  "/",
				Root:    "",
				Push:    "push",
			},
			AppID: "AppID",
		}
		expected := nnp.ProducerConfig{
			AppID: "AppID",
			Endpoint: nnp.ProducerEndpointConfig{
				RootURL: url.URL{
					Scheme: "http",
					Host:   "localhost:1234",
					Path:   "/",
				},
				PushURL: url.URL{
					Scheme: "http",
					Host:   "localhost:1234",
					Path:   "/push",
				},
			},
		}

		// when
		actual, err := converter.ProducerConfigFrom(source)

		// then
		if assert.NoError(t, err) {
			assert.Equal(t, expected, actual)
		}
	})

	t.Run("invalid", func(t *testing.T) {
		source := config.Config{}

		// when
		_, err := converter.ProducerConfigFrom(source)

		// then
		assert.Error(t, err)
	})
}
