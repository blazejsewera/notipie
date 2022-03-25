package test

import (
	"github.com/blazejsewera/notipie/core/internal/infra"
	"testing"
)

func initCore(t testing.TB) {
	t.Helper()
	appCtx := new(infra.AppContext)
	appCtx.Init(config)
	go appCtx.Start()
	t.Log("initCore: started core")
}
