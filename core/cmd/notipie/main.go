package main

import (
	"github.com/jazzsewera/notipie/core/internal/impl"
)

func main() {
	//hub := net.NewHub()

	endpoint := impl.Endpoint{}

	endpoint.Setup()
	endpoint.Run()
}
