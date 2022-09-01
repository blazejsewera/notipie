package test

import (
	"testing"
	"time"

	"github.com/blazejsewera/notipie/core/internal/infra"
)

func initCore(t testing.TB, port int) {
	t.Helper()
	appCtx := new(infra.AppContext)
	appCtx.Init(config(port))
	go appCtx.Start()
	time.Sleep(150 * time.Millisecond) // TODO: change to better application readiness check to stabilize CI pipeline
	t.Log("initCore: started core")
}
