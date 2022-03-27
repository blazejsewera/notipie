# Core project architecture
The project on the top-level is structured over 4 directories:
- `cmd` -- entry point to the application (`main`)
- `internal` -- application-specific code
- `pkg` -- reusable utils, not specific to the application
- `test` -- black box integration test code

## The `internal` directory
Application-specific code is split into 5 directories representing the levels
of abstraction:
- `domain`
- `grid`
- `impl`
- `infra`
- `model`
