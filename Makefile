.PHONY: clean

# install & no-params

all: install

install: install-workspace install-core
	@echo "> workspace ready"

install-workspace:
	@yarn
	@echo "> workspace for ui and project-level deps synced"
	@echo "> husky installed for git"

install-core:
	@cd core; \
	$(MAKE) sync; \
	$(MAKE) tidy
	@echo "> workspace for core synced"


# clean

nuke: clean
	@rm -rf node_modules
	@cd ui; \
	rm -rf node_modules
	@echo "> nuked all node_modules"

clean: clean-ui clean-core
	@echo "> cleaned"

clean-ui:
	@cd ui; \
	rm -rf dist
	@echo "> removed dist from ui"

clean-core:
	@cd core; \
	rm -rf notipie
	@echo "> removed binary from core"


# build

build: build-ui build-core
	@echo "> built"

build-ui:
	@cd ui; \
	yarn build
	@echo "> built dist in ui"

build-core:
	@cd core; \
	$(MAKE) build
	@echo "> built binary in core"


# dev

dev-ui:
	@cd ui; \
	yarn dev

dev-core:
	@cd core; \
	$(MAKE) run


# test

test: test-ui test-core
	@echo "> tests completed"

test-ui:
	@cd ui; \
	yarn test
	@echo "> completed tests in ui"

test-core:
	@cd core; \
	$(MAKE) test
	@echo "> completed tests in core"


# lint

lint: lint-ui lint-core
	@echo "> linted"

lint-ui:
	@cd ui; \
	yarn lint
	@echo "> linted ui"

lint-core:
	@cd core; \
	$(MAKE) lint
	@echo "> linted core"


# format

format: format-ui format-core
	@echo "> formatted"

format-ui:
	@cd ui; \
	yarn format
	@echo "> formatted ui"

format-core:
	@cd core; \
	$(MAKE) format
	@echo "> formatted core"
