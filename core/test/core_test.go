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

var config = infra.Config{
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

func initCore(t testing.TB) {
	t.Helper()
	appCtx := new(infra.AppContext)
	appCtx.Init(config)
	go appCtx.Start()
	t.Log("initCore: started core")
}
