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
	port          = 5150
	prefix        = "/"
	root          = ""
	push          = "push"
	ws            = "ws"
	notifications = "notifications"

	httpScheme = "http"
	wsScheme   = "ws"
)

var host = fmt.Sprintf("%s:%d", address, port)

var config = infra.Config{
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

var (
	rootURL          = url.URL{Scheme: httpScheme, Host: host, Path: getPath(root)}
	pushURL          = url.URL{Scheme: httpScheme, Host: host, Path: getPath(push)}
	wsURL            = url.URL{Scheme: wsScheme, Host: host, Path: getPath(ws)}
	notificationsURL = url.URL{Scheme: httpScheme, Host: host, Path: getPath(notifications)}
)

func getPath(target string) string {
	return prefix + target
}
