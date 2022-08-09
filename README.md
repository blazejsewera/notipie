# Notipie
A notification aggregator.

[![Notipie CI](https://github.com/blazejsewera/notipie/actions/workflows/ci.yml/badge.svg)](https://github.com/blazejsewera/notipie/actions/workflows/ci.yml)

## Quick links
- [board](https://github.com/orgs/blazejsewera/projects/1)
- [issues](https://github.com/blazejsewera/notipie/issues)
  - [new issue](https://github.com/blazejsewera/notipie/issues/new)

## Prerequisites
You will need:
- `go` version `1.18` or higher,
- `make`,
- `yarn` with `node.js` version `lts/gallium = v16`.

## Quick start
Just run `make` from the project root. All dependencies for all modules should
automatically download, and example configs should be copied to `core` and `ui`.

Read the output of this command, it will inform you about your nvm and Go
installation.

To quickly run the application for development, execute `make -j2 dev`.

### Recommended environment setup
#### Go
For Linux, use a package manager to download Go `1.18` or higher, e.g.,
`sudo apt install golang-go`, or `sudo pacman -S go`, depending on your distro.

For macOS, install Go with [Homebrew](https://formulae.brew.sh/formula/go).

For Windows, you can use [Scoop](https://scoop.sh/#/apps?q=go), or
[Chocolatey](https://community.chocolatey.org/packages/go).

Alternatively, [download it from go.dev](https://go.dev/dl/) and install
manually.

#### nvm
For Linux and macOS, install [nvm for POSIX](https://github.com/nvm-sh/nvm).

For Windows, install [nvm for Windows](https://github.com/coreybutler/nvm-windows).

#### Setup Node.js with nvm
Simply run `nvm install` and `nvm use` from the project root. It will
automatically install the latest `lts/gallium`.

You can also install the latest `v16` manually, downloading it from
[this repository](https://nodejs.org/dist/latest-gallium/).

## Project components
- `core` -- backend
- `ui` -- frontend

---------------------------------------------------------------------------------

`{}` and `<>` with ❤️ by Blazej Sewera&emsp;┃ [www](https://www.sewera.dev) ┆
[gh](https://github.com/blazejsewera) ┆ [gh:pv](https://github.com/sewera) ┃
