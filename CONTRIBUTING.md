# Contributing to billtop

## Development Setup

Requirements:
- Go 1.23+
- Docker (for container builds)

```bash
git clone https://github.com/jopnelli/billtop.git
cd billtop
go mod download
go build ./cmd/billtop
./billtop version
```

## Running Tests

```bash
go test -race ./...
```

## Linting

```bash
# Install golangci-lint: https://golangci-lint.run/welcome/install/
golangci-lint run
```

## Making Changes

1. Fork the repository
2. Create a branch from `main`
3. Make your changes
4. Ensure tests pass and lint is clean
5. Open a pull request

### Commit Messages

We use [Conventional Commits](https://www.conventionalcommits.org/):

```
feat(report): add --format json flag
fix(bigquery): handle empty credits array
docs: update setup instructions
test(anomaly): add edge case for zero baseline
```

### Code Style

- Follow the patterns in existing code
- Domain logic goes in `internal/domain/`, never imports adapter packages
- Adapters implement interfaces from `internal/port/`
- CLI commands and HTTP handlers are thin — delegate to domain services
- No global mutable state
- Every function doing I/O takes `context.Context` as first argument
- Use `log/slog` for logging, `fmt.Errorf("doing x: %w", err)` for error wrapping

## Architecture

```
internal/
  domain/    Pure business logic (no external deps)
  port/      Interfaces the domain needs
  adapter/   Infrastructure implementations
  server/    HTTP API
  cli/       CLI commands
```

See IMPLEMENTATION.md (not tracked in git) for detailed architecture decisions.
