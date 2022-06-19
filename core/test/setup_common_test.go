package test

import (
	"github.com/blazejsewera/notipie/core/pkg/lib/api"
	"net/url"

	"github.com/blazejsewera/notipie/core/internal/impl"
	"github.com/blazejsewera/notipie/core/internal/infra"
)

const (
	isProd  = false
	address = "localhost"
)

func config(port int) infra.Config {
	return infra.Config{
		Prod: isProd,
		EndpointConfig: impl.EndpointConfig{
			Address: address,
			Port:    port,
		},
	}
}

func pushURL(port int) url.URL {
	return getURL(port, api.Push)
}
func wsURL(port int) url.URL {
	return getURL(port, api.WebSocket)
}
func notificationsURL(port int) url.URL {
	return getURL(port, api.Notifications)
}

func getURL(port int, path api.Path) url.URL {
	return api.GetURL(api.GetHost(address, port), path)
}
