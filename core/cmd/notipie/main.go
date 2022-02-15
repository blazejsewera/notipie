package main

import (
	"github.com/jazzsewera/notipie/core/internal/impl"
	"github.com/jazzsewera/notipie/core/internal/impl/net"
)

func main() {
	hub := net.NewHub()

	endpoint := impl.Endpoint{}

	endpoint.SetupFor(hub)
	endpoint.Run()
}
