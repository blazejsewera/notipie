.PHONY: clean

all: clean sync tidy lint build test

build:
	go build github.com/jazzsewera/notipie/core/cmd/notipie

run:
	go run github.com/jazzsewera/notipie/core/cmd/notipie

test:
	go test ./...

lint:
	go vet ./...

format:
	go fmt ./...

sync:
	go mod download

tidy:
	go mod tidy

clean:
	rm -f ./notipie
	go clean
