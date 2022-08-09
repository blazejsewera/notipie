NOTIPIE=github.com/blazejsewera/notipie/core/cmd/notipie

PREFIX=> core >

NOTIPIE_CONFIG_FILENAME=notipie.config.json
NOTIPIE_CONFIG_EXAMPLE_FILENAME=notipie.config.example.json

DIST_DIR=dist

.PHONY: clean test

all: clean sync lint build test

build:
	@go build $(NOTIPIE)
	@echo "$(PREFIX) binary ./notipie built"

copy-example-configs:
	@echo "$(PREFIX) additional configs for core copied"

docker-build:
	@docker build -f ./Dockerfile -t notipie-core .
	@echo "$(PREFIX) docker image built"

_build-binary-in-docker:
	@go build -o /notipie $(NOTIPIE)
	@echo "$(PREFIX) binary /notipie built for docker"

run:
	@go run $(NOTIPIE)

test:
	@go test ./...

test-race:
	@go test -race ./...

pre-commit: format
	@echo "$(PREFIX) pre-commit ok"

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
	@rm -f ./notipie
	@go clean
	@go clean -cache
	@go clean -testcache
	@echo "$(PREFIX) binaries removed, go cache cleaned"
