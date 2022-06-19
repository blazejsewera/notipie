package api_test

import (
	"github.com/blazejsewera/notipie/core/pkg/lib/api"
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

func TestUtil(t *testing.T) {
	// given
	addr := "localhost"
	port := 5150

	t.Run("host", func(t *testing.T) {
		// when
		host := api.GetHost(addr, port)

		// then
		assert.Equal(t, "localhost:5150", host)
	})

	t.Run("path", func(t *testing.T) {
		// when
		path := api.GetPath(api.Push)

		// then
		assert.Equal(t, "/push", path)
	})

	// given
	host := api.GetHost(addr, port)

	t.Run("http url string", func(t *testing.T) {
		// when
		actual := api.GetURLStr(host, api.Push)

		// then
		assert.Equal(t, "http://localhost:5150/push", actual)
	})

	t.Run("ws url string", func(t *testing.T) {
		// when
		actual := api.GetURLStr(host, api.WebSocket)

		// then
		assert.Equal(t, "ws://localhost:5150/ws", actual)
	})

	t.Run("http url", func(t *testing.T) {
		// given
		expected := url.URL{Scheme: api.HttpSchema, Host: host, Path: "/" + api.Notifications.Path}

		// when
		actual := api.GetURL(host, api.Notifications)

		// then
		assert.Equal(t, expected, actual)
	})
}
