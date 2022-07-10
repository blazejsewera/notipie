# Manual test utilities

## Notifications server
Produces sample notifications for the ui.

Run:
```sh
yarn start
```

Then open a ui with `make dev-ui` from the project root. The initial
notifications should be fetched by the ui. You can also press <kbd>Enter</kbd>
everytime you want to push a new notification.

## WS Client
Acts as a simple web socket client for the core.

Run:
```sh
yarn start
```

Then launch core with `make dev-core` from the project root. When you push a
notification to backend (using, e.g. `nnp` from the `producer` directory),
you should see the pushed notification on your terminal.

## Adding new test utilities
When you add such utilities with `yarn` as dependency manager,
you need to link `.yarnrc.yml`, and `.yarn/` from the project root.

```sh
ln -s ../../.yarnrc.yml .
ln -s ../../.yarn .
```
