package main

import (
	"github.com/blazejsewera/notipie/producer/cmd/nnp/internal/cli"
)

func main() {
	cli.Setup()
	cli.Execute()
}
