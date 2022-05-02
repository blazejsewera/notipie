# const

UI_DIR=ui
CORE_DIR=core

CONFIG_FILENAME=notipie.config.json
EXAMPLE_CONFIG_FILENAME=notipie.config.example.json
UI_CONFIG_FILENAME=${CONFIG_FILENAME}
EXAMPLE_UI_CONFIG_FILENAME=${EXAMPLE_CONFIG_FILENAME}
CORE_CONFIG_FILENAME=${CONFIG_FILENAME}
EXAMPLE_CORE_CONFIG_FILENAME=${EXAMPLE_CONFIG_FILENAME}

.PHONY: clean


# install & no-params

all: install copy-example-configs

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

install: install-workspace install-core
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


# clean

nuke: clean remove-configs
	@rm -rf node_modules
	@cd ${UI_DIR}; \
	rm -rf node_modules
	@echo "> nuked all node_modules"

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
	rm -rf dist
	@echo "> removed dist from ui"

clean-core:
	@cd ${CORE_DIR}; \
	rm -rf notipie
	@echo "> removed binary from core"


# build

build: build-ui build-core
	@echo "> built"

build-ui:
	@cd ${UI_DIR}; \
	yarn build
	@echo "> built dist in ui"

build-core:
	@cd ${CORE_DIR}; \
	$(MAKE) build
	@echo "> built binary in core"


# dev

dev: dev-ui dev-core  # run with -j2

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

test: test-ui test-core
	@echo "> tests completed"

test-ui:
	@cd ${UI_DIR}; \
	yarn test
	@echo "> completed tests in ui"

test-core:
	@cd ${CORE_DIR}; \
	$(MAKE) test
	@echo "> completed tests in core"


# lint

lint-fix: lint-ui-fix lint-core
	@echo "> linted and fixed"

lint-ui-fix:
	@cd ${UI_DIR}; \
	yarn lint:fix
	@echo "> linted and fixed ui"

lint: lint-ui lint-core
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
