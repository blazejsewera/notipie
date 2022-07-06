# const

NVM:=$(shell nvm -v 2>/dev/null)
GO:=$(shell go version 2>/dev/null)

UI_DIR=ui
CORE_DIR=core
PRODUCER_DIR=producer
STACK_DIR=stack

CONFIG_FILENAME=notipie.config.json
EXAMPLE_CONFIG_FILENAME=notipie.config.example.json
UI_CONFIG_FILENAME=${CONFIG_FILENAME}
EXAMPLE_UI_CONFIG_FILENAME=${EXAMPLE_CONFIG_FILENAME}
CORE_CONFIG_FILENAME=${CONFIG_FILENAME}
EXAMPLE_CORE_CONFIG_FILENAME=${EXAMPLE_CONFIG_FILENAME}

.PHONY: clean


# install & no-params

all: print-env-help install copy-example-configs

print-env-help: print-env-help-ui print-env-help-core

print-env-help-ui:
ifneq (${NVM_DIR},"")
	@echo "> nvm installed, run 'nvm install' and 'nvm use' in your shell"
else ifneq ("$(wildcard $(HOME)/.nvm/nvm.sh)","")
	@echo "> nvm installed, run 'nvm install' and 'nvm use' in your shell"
else
$(error nvm not installed, consider installing it or install latest LTS node.js manually)
endif

print-env-help-core:
ifdef GO
	@echo "> go installed"
	@echo "> $(shell go version)"
else
$(error go not installed, install version specified in 'core/go.mod' or later)
endif

copy-example-configs: copy-example-config copy-configs
	@echo "> configs copied"

copy-example-config:
	@test -f ${CONFIG_FILENAME} && \
	echo "> config already exists" || (\
		cp ${EXAMPLE_CONFIG_FILENAME} ${CONFIG_FILENAME} && \
		echo "> example config copied (${EXAMPLE_CONFIG_FILENAME} => ${CONFIG_FILENAME})" \
	)

copy-configs: copy-config-ui copy-config-core

copy-config-ui:
	@test -f ${UI_DIR}/${UI_CONFIG_FILENAME} && \
	echo "> config for ui already exists" || (\
		cp ${CONFIG_FILENAME} ${UI_DIR}/${UI_CONFIG_FILENAME} && \
		echo "> config for ui copied (${CONFIG_FILENAME} => ${UI_DIR}/${UI_CONFIG_FILENAME})" \
	)

copy-config-core:
	@test -f ${CORE_DIR}/${CORE_CONFIG_FILENAME} && \
	echo "> config for core already exists" || (\
		cp ${CONFIG_FILENAME} ${CORE_DIR}/${CORE_CONFIG_FILENAME} && \
		echo "> config for core copied (${CONFIG_FILENAME} => ${CORE_DIR}/${CORE_CONFIG_FILENAME})" \
	)

uc: update-configs

update-configs: copy-example-config update-config-ui update-config-core
	@echo "> configs updated (${CONFIG_FILENAME} => ${UI_DIR}/${UI_CONFIG_FILENAME}, ${CONFIG_FILENAME} => ${CORE_DIR}/${CORE_CONFIG_FILENAME})"

update-config-ui:
	@cp ${CONFIG_FILENAME} ${UI_DIR}/${UI_CONFIG_FILENAME}
	@echo "> config updated (${CONFIG_FILENAME} => ${UI_DIR}/${UI_CONFIG_FILENAME})"

update-config-core:
	@cp ${CONFIG_FILENAME} ${CORE_DIR}/${CORE_CONFIG_FILENAME}
	@echo "> config updated (${CONFIG_FILENAME} => ${CORE_DIR}/${CORE_CONFIG_FILENAME})"

install: install-workspace install-core install-producer
	@echo "> workspace ready"

install-workspace:
	@yarn
	@echo "> workspace for ui and project-level deps synced"
	@echo "> husky installed for git"

install-core:
	@cd ${CORE_DIR}; \
	$(MAKE) sync; \
	$(MAKE) tidy
	@echo "> workspace for core synced"

install-producer:
	@cd ${PRODUCER_DIR}; \
	$(MAKE) sync; \
	$(MAKE) tidy
	@echo "> workspace for producer synced"


# clean

nuke: clean remove-configs
	@rm -rf .yarn/cache
	@rm -rf node_modules
	@cd ${UI_DIR}; \
	rm -rf node_modules
	@echo "> nuked all node_modules and yarn cache"

rmc: remove-configs

remove-configs: remove-configs-ui remove-configs-core
	@rm -f ${CONFIG_FILENAME}
	@echo "> configs removed"

remove-configs-ui:
	@rm -f ${UI_DIR}/${UI_CONFIG_FILENAME}
	@echo "> config for ui removed"

remove-configs-core:
	@rm -f ${CORE_DIR}/${CORE_CONFIG_FILENAME}
	@echo "> config for core removed"

clean: clean-ui clean-core
	@echo "> cleaned"

clean-ui:
	@cd ${UI_DIR}; \
	rm -rf dist; \
	rm -rf storybook-static
	@echo "> removed dist and storybook-static from ui"

clean-core:
	@cd ${CORE_DIR}; \
	$(MAKE) clean
	@echo "> cleaned core"


# build

b: build  # run with -j2 or more for performance
bui: build-ui
bstorybook: build-storybook
bcore: build-core
bproducer: build-producer

build: build-ui build-storybook build-core build-producer  # run with -j2 or more for performance
	@echo "> built"

build-ui:
	@cd ${UI_DIR}; \
	yarn build
	@echo "> built dist in ui"

build-storybook:
	@cd ${UI_DIR}; \
	yarn build-storybook

build-core:
	@cd ${CORE_DIR}; \
	$(MAKE) build
	@echo "> built binary in core"

build-producer:
	@cd ${PRODUCER_DIR}; \
	$(MAKE) build
	@echo "> built binary in producer"


# docker

docker-ui: build-ui
	@cd ${UI_DIR}; \
	yarn docker
	@echo "> docker image notipie-ui built"

docker-core:
	@cd ${CORE_DIR}; \
	$(MAKE) docker
	@echo "> docker image notipie-core built"

docker-up: docker-ui docker-core
	@cd ${STACK_DIR}; \
	docker compose up

docker-down:
	@cd ${STACK_DIR}; \
	docker compose down


# dev

d: dev  # run with -j2 or more
dui: dev-ui
dcore: dev-core

dev: print-multithread-notice dev-ui dev-core  # run with -j2 or more

print-multithread-notice:
	@echo "> check if you are running make with -j2 or more"

dev-ui:
	@cd ${UI_DIR}; \
	yarn dev

dev-core:
	@cd ${CORE_DIR}; \
	$(MAKE) run

dev-manual-test:
	@cd test/manual; \
	yarn start

# test

t: test  # run with -j2 or more for performance
tui: test-ui
tcore: test-core
tproducer: test-producer
tr: test-race

test: test-ui test-core test-producer  # run with -j2 or more for performance
	@echo "> tests completed"

test-ui:
	@cd ${UI_DIR}; \
	yarn test
	@echo "> completed tests in ui"

test-core:
	@cd ${CORE_DIR}; \
	$(MAKE) test
	@echo "> completed tests in core"

test-producer:
	@cd ${PRODUCER_DIR}; \
	$(MAKE) test
	@echo "> completed tests in producer"

test-race: test-core-race test-producer-race
	@echo "> tests with race detection completed"

test-core-race:
	@cd ${CORE_DIR}; \
	$(MAKE) test-race
	@echo "> completed tests in core with race detection"

test-producer-race:
	@cd ${PRODUCER_DIR}; \
	$(MAKE) test-race
	@echo "> completed tests in producer with race detection"


# lint

lint-fix: lint-ui-fix lint-core lint-producer
	@echo "> linted and fixed"

lint-ui-fix:
	@cd ${UI_DIR}; \
	yarn lint:fix
	@echo "> linted and fixed ui"

lint: lint-ui lint-core lint-producer
	@echo "> linted"

lint-ui:
	@cd ${UI_DIR}; \
	yarn lint
	@echo "> linted ui"

lint-ui-staged:
	@cd ${UI_DIR}; \
	yarn lint:staged
	@echo "> linted staged ui"

lint-core:
	@cd ${CORE_DIR}; \
	$(MAKE) lint
	@echo "> linted core"

lint-producer:
	@cd ${PRODUCER_DIR}; \
	$(MAKE) lint
	@echo "> linted producer"


# format

format: format-ui format-core
	@echo "> formatted"

format-ui:
	@cd ${UI_DIR}; \
	yarn format
	@echo "> formatted ui"

format-core:
	@cd ${CORE_DIR}; \
	$(MAKE) format
	@echo "> formatted core"


# pre-commit check

pre-commit: format-core lint-ui-staged
	@echo "> pre-commit check done"
