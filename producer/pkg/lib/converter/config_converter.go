package converter

import (
	"errors"
	"github.com/blazejsewera/notipie/producer/pkg/lib/config"
	"github.com/blazejsewera/notipie/producer/pkg/lib/nnp"
	"net/url"
	"strconv"
)

func ProducerConfigFrom(serializable config.Config) (nnp.ProducerConfig, error) {
	if serializable.Address == "" {
		return nnp.ProducerConfig{}, errors.New("convert config: address was empty")
	}

	rawRootURL := rawRootURLOf(serializable)
	rootURL, err := url.Parse(rawRootURL)
	if err != nil {
		return nnp.ProducerConfig{}, err
	}

	rawPushURL := rawPushURLOf(serializable)
	pushURL, err := url.Parse(rawPushURL)
	if err != nil {
		return nnp.ProducerConfig{}, err
	}

	return nnp.ProducerConfig{
		AppID: serializable.AppID,
		Endpoint: nnp.ProducerEndpointConfig{
			RootURL: *rootURL,
			PushURL: *pushURL,
		},
	}, nil
}

func rawRootURLOf(c config.Config) string {
	return rawURLOf(c, c.Root)
}

func rawPushURLOf(c config.Config) string {
	return rawURLOf(c, c.Push)
}

func rawURLOf(c config.Config, path string) string {
	schema := "http://"
	return schema + c.Address + ":" + strconv.Itoa(c.Port) + c.Prefix + path
}
