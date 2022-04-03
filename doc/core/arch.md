# Core project architecture

The project on the top-level is structured over 4 directories:

- `cmd` -- entry point to the application (`main`)
- `internal` -- application-specific code
- `pkg` -- reusable utils, not specific to the application
- `test` -- black box integration test code

## The `internal` directory

Application-specific code is split into 5 directories representing the levels
of abstraction:

- `domain` -- business logic, defines data structures and communication of
  domain objects on the highest level of abstraction,
- `grid` -- lower level of abstraction than domain, defines proxies that convert
  network models into domain models, creates and organizes domain objects into a
  grid, in which those objects can communicate,
- `impl` -- implements network endpoints, WebSockets, and persistence,
- `infra` -- configures the application and sets up the context for DI,
- `model` -- defines network data models.
