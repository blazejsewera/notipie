package infra

import "github.com/jazzsewera/notipie/core/pkg/lib/log"

func Init(config Config) {
	initLogger(config.prod)
}

func initLogger(prod bool) {
	log.Init(prod)
}

func Start() {

}
