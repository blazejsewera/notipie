.PHONY: clean test

all: clean sync tidy lint build test

build:
	go build github.com/blazejsewera/notipie/core/cmd/notipie

run:
	go run github.com/blazejsewera/notipie/core/cmd/notipie

test:
	go test ./...

test-race:
	go test -race ./...

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
