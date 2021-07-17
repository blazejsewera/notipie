all: clean tidy lint build test

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

tidy:
	go mod tidy

clean:
	go clean
