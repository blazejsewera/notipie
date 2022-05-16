.PHONY: clean test

all: clean sync tidy lint build test

build:
	go build github.com/blazejsewera/notipie/producer/cmd/nnp

run:
	go run github.com/blazejsewera/notipie/producer/cmd/nnp

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
	rm -f ./nnp
	go clean
