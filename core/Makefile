NOTIPIE=github.com/blazejsewera/notipie/core/cmd/notipie

.PHONY: clean test

all: clean sync lint build test

build:
	@go build $(NOTIPIE)
	@echo "> dev binary ./notipie built for current platform"

build-dist:
	@mkdir -p dist
	@cp ./notipie.config.json ./dist/notipie.config.json
	@CGO_ENABLED=0 go build -o ./dist/notipie $(NOTIPIE)
	@echo "> dist statically-linked binary built"
	@echo "> can be used with docker `FROM scratch`"
	@echo "> binary and notipie.config.json copied to dist dir"

docker: build-dist
	@docker build -f ./Dockerfile -t notipie-core ./dist
	@echo "> docker image built"

run:
	@go run $(NOTIPIE)

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

tidy:
	@go mod tidy

clean:
	@rm -f ./notipie
	@rm -rf ./dist
	@go clean
	@go clean -cache
	@go clean -testcache
	@echo "> binaries removed, go cache cleaned"
