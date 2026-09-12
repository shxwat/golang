# Go Learning Repository

A collection of small, independent Go programs created while learning the
language and backend development. The examples cover fundamentals, HTTP APIs,
concurrency, error handling, interfaces, pointers, and a few mini projects.

## Topics

- Go basics, functions, maps, structs, pointers, and interfaces
- Error handling and wrapping
- Goroutines, channels, wait groups, and mutexes
- HTTP handlers and JSON APIs
- PostgreSQL integration
- Small concurrency projects such as a rate limiter, worker pool, link checker,
  job processor, and download manager
- A separate step-by-step backend series in [`go-backend`](./go-backend)

## Requirements

- Go 1.26.6 or newer
- PostgreSQL only for `apis/postgres_user_api.go`

## Running an example

Most topic directories contain multiple standalone programs. Run one file at a
time from the repository root:

```sh
go run basic/hello_world.go
go run channels/buffered_channel.go
go run Rate_Limiter/main.go
```

The `go-backend` directory is its own Go module. Follow its
[`README.md`](./go-backend/README.md) for those examples.

## PostgreSQL example

The PostgreSQL API reads its connection string from `DATABASE_URL`; credentials
are not stored in the source code. Start PostgreSQL, then set the variable before
running the example:

```sh
export DATABASE_URL='postgres://postgres:your_password@localhost:5433/postgres?sslmode=disable'
go run apis/postgres_user_api.go
```

See [`.env.example`](./.env.example) for the expected format. Keep actual `.env`
files local; they are excluded by `.gitignore`.
