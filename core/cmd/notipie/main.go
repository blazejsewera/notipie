package main

import notipie "github.com/blazejsewera/notipie/core/internal/infra"

func main() {
	config := notipie.DefaultConfig()
	appCtx := notipie.AppContext{}
	appCtx.Init(config)
	appCtx.Start()
}
