NNP=github.com/blazejsewera/notipie/producer/cmd/nnp

PREFIX=> producer >

.PHONY: clean test

all: clean sync lint build test

build:
	@go build $(NNP)
	@echo "$(PREFIX) binary ./nnp built"

run:
	@go run $(NNP)

test:
	@go test ./...

test-race:
	@go test -race ./...

lint:
	@go vet ./...

format:
	@go fmt ./...

sync:
	@go mod download
	@go mod verify

tidy:
	@go mod tidy

clean:
	@rm -f ./nnp
	@go clean
	@go clean -cache
	@go clean -testcache
	@echo "$(PREFIX) binaries removed, go cache cleaned"
