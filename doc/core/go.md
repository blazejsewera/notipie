# Go in core
Go is quickly gaining popularity among developers, with its great tooling,
focus on microservices, and state-of-the-art standard library.

## Motivation
When choosing the right language for the project, I focused on finding the
right tool for the application and developer experience.

I wanted `notipie` to be a high-performance microservice, so I didn't take
interpreted languages like Python or JavaScript into consideration.
I mostly considered Java, Kotlin, Rust, and Go.

Java, although popular, doesn't have the greatest developer experience. Things
like `equals` and `hashCode` are unnecessary bloat in the code. Project
Lombok fixes some of them, but the tooling is limited to IntelliJ, you have to
download a lot of libraries for dealing with JSON, create your own code style
guide, and perform a fair bit of setup.

Kotlin, far better than Java, but also locked-in to IntelliJ with tooling, was
an interesting option for me, but not ideal.

Rust was too low-level for my application. Explicit memory management,
although performant, was simply too verbose and work-intensive for my use
case.

Go was a perfect option. A plethora of great tooling, like first-party Go
plugin for VSCode, GoLand from JetBrains, community plugins for Neovim, all
working great and providing a good developer experience.
Furthermore, extraordinary performance of the tooling itself, with tests
running in under a second, super-fast compiler, one of the best standard
libraries I've seen, and overall simplicity of the language, made the choice
obvious.

## How did Go make the development easier

### Built-in language features
The feature that helped the most during development was channels and
_goroutines_, coroutines automatically managed by the Go runtime. The idea
behind those was very simple to understand, and working with concurrent
programming was a little bit easier.

### Standard library
Standard `testing` package provides a unified, and simple tooling for testing.
I didn't have to think anything about test setup. No custom scripts,
third-party libraries, or IDE setup. All I needed to do was to name a file
with `_test.go` suffix, write a function starting with `Test`, and run
`go test ./...`. Both VSCode with Go plugin, and GoLand automatically picked
up the test setup and I was ready to develop with TDD.

Standard `net/http` package provides everything needed for setting up REST
endpoints. Although I used [Gin](https://gin-gonic.com) for this, due to a
simpler interface, I used status codes and HTTP client implementation from
`net/http`.

### Third-party libraries
Gin was great for writing REST endpoints, with `gin.Context` having easy
access to standard-library-compatible fields, making it easily pluggable to
other third-party libraries, like
[Gorilla WebSockets](https://github.com/gorilla/websocket).

[Zap](https://github.com/uber-go/zap) provided an excellent and performant way
to log things in the backend. Structured logging, automatic serialization to
JSON in production mode, and human-readable format in debug mode, paired with
low or zero-allocation overhead, made it a perfect choice for logging in a
microservice.
