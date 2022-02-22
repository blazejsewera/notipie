package main

import notipie "github.com/jazzsewera/notipie/core/internal/infra"

func main() {
	config := notipie.DefaultConfig()
	appCtx := notipie.AppContext{}
	appCtx.Init(config)
	appCtx.Start()
}
