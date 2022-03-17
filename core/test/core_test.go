package test

import (
	"github.com/blazejsewera/notipie/core/internal/impl"
	"github.com/blazejsewera/notipie/core/internal/infra"
	"testing"
)

const (
	prefix = "/"
	root   = ""
	push   = "push"
	ws     = "ws"
)

func initCore(t testing.TB) infra.Config {
	t.Helper()
	config := infra.Config{
		Prod: false,
		EndpointConfig: impl.EndpointConfig{
			Address:   "localhost",
			Port:      5150,
			Prefix:    prefix,
			Root:      root,
			Push:      push,
			WebSocket: ws,
		},
	}
	appCtx := new(infra.AppContext)
	appCtx.Init(config)
	go appCtx.Start()
	t.Log("initCore: started core")
	return config
}
