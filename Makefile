# const

NVM:=$(shell nvm -v 2>/dev/null)
GO:=$(shell go version 2>/dev/null)
DOCKER:=$(shell docker -v 2>/dev/null)
DOCKER_COMPOSE:=$(shell docker compose version 2>/dev/null)

HOOKS_DIR=$(shell git rev-parse --show-toplevel)/.hooks
GIT_DIR=$(shell git rev-parse --absolute-git-dir)
GIT_HOOKS_DIR=$(GIT_DIR)/hooks
COMMIT_MSG_HOOK=commit-msg
PRE_COMMIT_HOOK=pre-commit

UI_DIR=ui
CORE_DIR=core
PRODUCER_DIR=producer
STACK_DIR=stack

TEST_DIR=test
TEST_NOTIFICATIONS_SERVER_DIR=notifications-server
TEST_WS_CLIENT_DIR=ws-client

CONFIG_FILENAME=notipie.config.json
EXAMPLE_CONFIG_FILENAME=notipie.config.example.json
UI_CONFIG_FILENAME=$(CONFIG_FILENAME)
EXAMPLE_UI_CONFIG_FILENAME=$(EXAMPLE_CONFIG_FILENAME)
CORE_CONFIG_FILENAME=$(CONFIG_FILENAME)
EXAMPLE_CORE_CONFIG_FILENAME=$(EXAMPLE_CONFIG_FILENAME)

.PHONY: clean


# install & no-params

all: print-env-help install copy-example-configs

print-env-help: print-env-help-ui print-env-help-core print-env-help-stack

print-env-help-ui:
ifneq ($(NVM_DIR),"")
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

print-env-help-stack: print-env-help-docker print-env-help-docker-compose

print-env-help-docker:
ifdef DOCKER
	@echo "> docker installed"
	@echo "> $(shell docker -v)"
else
$(warning docker not installed, install to build docker images)
endif

print-env-help-docker-compose:
ifdef DOCKER_COMPOSE
	@echo "> docker compose installed"
	@echo "> $(shell docker compose version)"
else
$(warning docker compose not installed, install to use docker-compose.yml)
endif

copy-example-configs: copy-example-config copy-configs
	@echo "> configs copied"

copy-example-config:
	@test -f $(CONFIG_FILENAME) && \
	echo "> config already exists" || (\
		cp $(EXAMPLE_CONFIG_FILENAME) $(CONFIG_FILENAME) && \
		echo "> example config copied ($(EXAMPLE_CONFIG_FILENAME) => $(CONFIG_FILENAME))" \
	)

copy-configs: copy-config-ui copy-config-core

copy-config-ui:
	@test -f $(UI_DIR)/$(UI_CONFIG_FILENAME) && \
	echo "> config for ui already exists" || (\
		cp $(CONFIG_FILENAME) $(UI_DIR)/$(UI_CONFIG_FILENAME) && \
		echo "> config for ui copied ($(CONFIG_FILENAME) => $(UI_DIR)/$(UI_CONFIG_FILENAME))" \
	)

copy-config-core:
	@test -f $(CORE_DIR)/$(CORE_CONFIG_FILENAME) && \
	echo "> config for core already exists" || (\
		cp $(CONFIG_FILENAME) $(CORE_DIR)/$(CORE_CONFIG_FILENAME) && \
		echo "> config for core copied ($(CONFIG_FILENAME) => $(CORE_DIR)/$(CORE_CONFIG_FILENAME))" \
	)
	@cd $(CORE_DIR); \
		$(MAKE) copy-example-configs

uc: update-configs

update-configs: copy-example-config update-config-ui update-config-core
	@echo "> configs updated ($(CONFIG_FILENAME) => $(UI_DIR)/$(UI_CONFIG_FILENAME), $(CONFIG_FILENAME) => $(CORE_DIR)/$(CORE_CONFIG_FILENAME))"

update-config-ui:
	@cp $(CONFIG_FILENAME) $(UI_DIR)/$(UI_CONFIG_FILENAME)
	@echo "> config updated ($(CONFIG_FILENAME) => $(UI_DIR)/$(UI_CONFIG_FILENAME))"

update-config-core:
	@cp $(CONFIG_FILENAME) $(CORE_DIR)/$(CORE_CONFIG_FILENAME)
	@echo "> config updated ($(CONFIG_FILENAME) => $(CORE_DIR)/$(CORE_CONFIG_FILENAME))"

install: install-hooks install-ui install-core install-producer install-test
	@echo "> workspace ready"

install-hooks: clean-hooks
	@ln -s $(HOOKS_DIR)/$(PRE_COMMIT_HOOK) $(GIT_HOOKS_DIR)/$(PRE_COMMIT_HOOK)
	@ln -s $(HOOKS_DIR)/$(COMMIT_MSG_HOOK) $(GIT_HOOKS_DIR)/$(COMMIT_MSG_HOOK)

clean-hooks:
	@rm -f $(GIT_HOOKS_DIR)/*

install-ui: pre-install-ui
	@cd $(UI_DIR); \
	yarn
	@echo "> workspace for ui synced"

pre-install-ui:
	@cp -r .yarnrc.yml .yarn $(UI_DIR)
	@echo "> configured yarn for ui"

install-core:
	@cd $(CORE_DIR); \
	$(MAKE) sync; \
	$(MAKE) tidy
	@echo "> workspace for core synced"

install-producer:
	@cd $(PRODUCER_DIR); \
	$(MAKE) sync; \
	$(MAKE) tidy
	@echo "> workspace for producer synced"

install-test: install-test-notifications-server install-test-ws-client

install-test-notifications-server: pre-install-test-notifications-server
	@cd $(TEST_DIR)/$(TEST_NOTIFICATIONS_SERVER_DIR); \
	yarn
	@echo "> workspace for $(TEST_DIR)/$(TEST_NOTIFICATIONS_SERVER_DIR) synced"

pre-install-test-notifications-server:
	@cp -r .yarnrc.yml .yarn $(TEST_DIR)/$(TEST_NOTIFICATIONS_SERVER_DIR)
	@echo "> configured yarn for notifications-server"

install-test-ws-client: pre-install-test-ws-client
	@cd $(TEST_DIR)/$(TEST_WS_CLIENT_DIR); \
	yarn
	@echo "> workspace for $(TEST_DIR)/$(TEST_WS_CLIENT_DIR) synced"

pre-install-test-ws-client:
	@cp -r .yarnrc.yml .yarn $(TEST_DIR)/$(TEST_WS_CLIENT_DIR)
	@echo "> configured yarn for ws-client"


# clean

nuke: clean remove-configs
	@cd $(UI_DIR); \
	rm -rf node_modules; \
	rm -rf .yarn; \
	rm -f .yarnrc.yml
	@cd $(TEST_DIR)/$(TEST_NOTIFICATIONS_SERVER_DIR); \
	rm -rf node_modules; \
	rm -rf .yarn; \
	rm -f .yarnrc.yml
	@cd $(TEST_DIR)/$(TEST_WS_CLIENT_DIR); \
	rm -rf node_modules; \
	rm -rf .yarn; \
	rm -f .yarnrc.yml
	@echo "> nuked all node_modules and yarn configs"

rmc: remove-configs

remove-configs: remove-configs-ui remove-configs-core
	@rm -f $(CONFIG_FILENAME)
	@echo "> configs removed"

remove-configs-ui:
	@rm -f $(UI_DIR)/$(UI_CONFIG_FILENAME)
	@echo "> config for ui removed"

remove-configs-core:
	@rm -f $(CORE_DIR)/$(CORE_CONFIG_FILENAME)
	@echo "> config for core removed"

clean: clean-ui clean-core clean-test
	@echo "> cleaned"

clean-ui:
	@cd $(UI_DIR); \
	rm -rf dist; \
	rm -rf storybook-static
	@echo "> removed dist and storybook-static from ui"

clean-core:
	@cd $(CORE_DIR); \
	$(MAKE) clean
	@echo "> cleaned core"

clean-test: clean-test-notifications-server clean-test-ws-client

clean-test-notifications-server:
	@cd $(TEST_DIR)/$(TEST_NOTIFICATIONS_SERVER_DIR); \
	rm -rf build

clean-test-ws-client:
	@cd $(TEST_DIR)/$(TEST_WS_CLIENT_DIR); \
	rm -rf build


# build

b: build  # run with -j2 or more for performance
bui: build-ui
bstorybook: build-storybook
bcore: build-core
bproducer: build-producer
bt: build-test

build: build-ui build-storybook build-core build-producer  # run with -j2 or more for performance
	@echo "> built"

build-ui:
	@cd $(UI_DIR); \
	yarn build
	@echo "> built dist in ui"

build-storybook:
	@cd $(UI_DIR); \
	yarn build-storybook
	@echo "> built storybook in ui"

build-core:
	@cd $(CORE_DIR); \
	$(MAKE) build
	@echo "> built binary in core"

build-producer:
	@cd $(PRODUCER_DIR); \
	$(MAKE) build
	@echo "> built binary in producer"

build-test: build-test-notifications-server build-test-ws-client

build-test-notifications-server:
	@cd $(TEST_DIR)/$(TEST_NOTIFICATIONS_SERVER_DIR); \
	yarn build

build-test-ws-client:
	@cd $(TEST_DIR)/$(TEST_WS_CLIENT_DIR); \
	yarn build


# docker

docker-build: docker-build-ui docker-build-core

docker-build-ui: build-ui
	@cd $(UI_DIR); \
	yarn docker-build
	@echo "> docker image notipie-ui built"

docker-build-core:
	@cd $(CORE_DIR); \
	$(MAKE) docker-build
	@echo "> docker image notipie-core built"

docker-compose-up:
	@cd $(STACK_DIR); \
	docker compose up

docker-compose-up-d:
	@cd $(STACK_DIR); \
	docker compose up -d

docker-compose-down:
	@cd $(STACK_DIR); \
	docker compose down


# dev

d: dev  # run with -j2 or more
dui: dev-ui
dcore: dev-core

dev: print-multithread-notice dev-ui dev-core  # run with -j2 or more

print-multithread-notice:
	@echo "> check if you are running make with -j2 or more"

dev-ui:
	@cd $(UI_DIR); \
	yarn dev

dev-core:
	@cd $(CORE_DIR); \
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
	@cd $(UI_DIR); \
	yarn test
	@echo "> completed tests in ui"

test-core:
	@cd $(CORE_DIR); \
	$(MAKE) test
	@echo "> completed tests in core"

test-producer:
	@cd $(PRODUCER_DIR); \
	$(MAKE) test
	@echo "> completed tests in producer"

test-race: test-core-race test-producer-race
	@echo "> tests with race detection completed"

test-core-race:
	@cd $(CORE_DIR); \
	$(MAKE) test-race
	@echo "> completed tests in core with race detection"

test-producer-race:
	@cd $(PRODUCER_DIR); \
	$(MAKE) test-race
	@echo "> completed tests in producer with race detection"


# lint

lint-fix: lint-ui-fix lint-core lint-producer
	@echo "> linted and fixed"

lint-ui-fix:
	@cd $(UI_DIR); \
	yarn lint:fix
	@echo "> linted and fixed ui"

lint: lint-ui lint-core lint-producer
	@echo "> linted"

lint-ui:
	@cd $(UI_DIR); \
	yarn lint
	@echo "> linted ui"

lint-ui-staged:
	@cd $(UI_DIR); \
	yarn lint:staged
	@echo "> linted staged ui"

lint-core:
	@cd $(CORE_DIR); \
	$(MAKE) lint
	@echo "> linted core"

lint-producer:
	@cd $(PRODUCER_DIR); \
	$(MAKE) lint
	@echo "> linted producer"


# format

format: format-ui format-core
	@echo "> formatted"

format-ui:
	@cd $(UI_DIR); \
	yarn format
	@echo "> formatted ui"

format-core:
	@cd $(CORE_DIR); \
	$(MAKE) format
	@echo "> formatted core"


# pre-commit check

pre-commit: format-core lint-ui-staged
	@echo "> pre-commit check done"
