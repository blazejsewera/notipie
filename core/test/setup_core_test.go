package test

import (
	"github.com/blazejsewera/notipie/core/internal/infra"
	"testing"
)

func initCore(t testing.TB, port int) {
	t.Helper()
	appCtx := new(infra.AppContext)
	appCtx.Init(config(port))
	go appCtx.Start()
	t.Log("initCore: started core")
}
