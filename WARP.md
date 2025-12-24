# WARP.md

This file provides guidance to WARP (warp.dev) when working with code in this repository.

## Repository overview

- This repository is a Go module: `github.com/kvervandi/task-api`.
- The Go toolchain version in `go.mod` is `go 1.25.5`.
- At the time of writing, there are no Go source files, tests, or additional configuration files checked into the repository, so most architecture and workflows are not yet defined.

Future agents should re-scan the repository structure (e.g., `cmd/`, `internal/`, `pkg/`, tests, and configs) before making assumptions about architecture or commands.

## Commands

Because this is a standard Go module and no custom tooling (Makefile, task runner, etc.) is present, use the built-in Go tooling by default. These commands assume a typical Go project layout and may need adjustment once real packages and commands exist.

- **Run all tests in the module**
  - `go test ./...`
- **Run tests for a specific package**
  - `go test ./path/to/package`
- **Run a single test by name**
  - `go test ./path/to/package -run TestName`
- **Build all packages in the module**
  - `go build ./...`

If a `cmd/` directory or other entrypoint packages are added later, prefer `go build ./cmd/<name>` or `go run ./cmd/<name>` for building/running those binaries.

## Architecture and structure

Currently, the only tracked file is `go.mod`, so there is no established application structure yet (no `main` packages, libraries, or tests).

When additional code is added, future agents should:
- Identify top-level entrypoints (e.g., under `cmd/` or a root `main.go`).
- Map any domain packages (e.g., `internal/tasks`, `internal/api`, etc.) and how they depend on each other.
- Note how configuration is loaded (environment variables, config files, flags) if/when such logic appears.

Until more files exist, avoid assuming any particular architecture beyond a single Go module named `github.com/kvervandi/task-api`.