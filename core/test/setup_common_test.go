package test

import (
	"fmt"
	"github.com/blazejsewera/notipie/core/internal/impl"
	"github.com/blazejsewera/notipie/core/internal/infra"
	"net/url"
)

const (
	isProd        = false
	address       = "localhost"
	prefix        = "/"
	root          = ""
	push          = "push"
	ws            = "ws"
	notifications = "notifications"

	httpScheme = "http"
	wsScheme   = "ws"
)

func host(port int) string {
	return fmt.Sprintf("%s:%d", address, port)
}

func config(port int) infra.Config {
	return infra.Config{
		Prod: isProd,
		EndpointConfig: impl.EndpointConfig{
			Address:       address,
			Port:          port,
			Prefix:        prefix,
			Root:          root,
			Push:          push,
			WebSocket:     ws,
			Notifications: notifications,
		},
	}
}

func rootURL(port int) url.URL {
	return url.URL{Scheme: httpScheme, Host: host(port), Path: getPath(root)}
}
func pushURL(port int) url.URL {
	return url.URL{Scheme: httpScheme, Host: host(port), Path: getPath(push)}
}
func wsURL(port int) url.URL {
	return url.URL{Scheme: wsScheme, Host: host(port), Path: getPath(ws)}
}
func notificationsURL(port int) url.URL {
	return url.URL{Scheme: httpScheme, Host: host(port), Path: getPath(notifications)}
}

func getPath(target string) string {
	return prefix + target
}
