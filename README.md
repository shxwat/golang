# Go Learning Repository

A hands-on collection of standalone Go programs covering language fundamentals,
backend APIs, error handling, concurrency patterns, and small worker-based
systems. Each example focuses on one concept so it can be read and run without a
large application setup.

## What This Repository Covers

- Go syntax, variables, arrays, slices, maps, loops, and `defer`
- Functions, variadic arguments, and multiple return values
- Structs, methods, pointers, and interfaces
- Idiomatic error creation, comparison, propagation, and wrapping
- Goroutines, channels, `select`, wait groups, mutexes, and contexts
- HTTP servers, routing, path values, query parameters, and JSON
- In-memory CRUD APIs and PostgreSQL-backed persistence
- Worker pools, timeouts, rate limiting, graceful shutdown, and concurrent jobs

## Repository Structure

| Path | Contents |
| --- | --- |
| `basic/` | Core syntax, data types, control flow, arrays, and slices |
| `functions/` | Parameters, return values, calculations, and variadic functions |
| `maps/` | Map creation, lookup, update, and pricing/age examples |
| `structs/` | Custom data types and functions that work with structs |
| `pointers/` | Pointer-based mutation, methods, carts, wallets, and game examples |
| `interfaces/` | Interface contracts for vehicles, devices, storage, and charging |
| `errors/` | Validation, sentinel errors, wrapping, and error-chain examples |
| `goroutines/` | Concurrent tasks using goroutines, wait groups, and mutexes |
| `channels/` | Buffered/unbuffered channels, `select`, hand-offs, and URL checks |
| `apis/` | Standalone HTTP and JSON API examples |
| `go-backend/` | A separate, progressive backend-learning module |
| Project folders | Larger concurrency exercises, each with its own `main.go` |

> [!IMPORTANT]
> Topic folders such as `basic`, `functions`, `apis`, `channels`, and
> `goroutines` contain multiple independent programs with separate `main()`
> functions. Run a specific file from these folders, not the entire folder.

## Standalone Projects

These directories contain one runnable program each:

| Project | What it demonstrates | Run command |
| --- | --- | --- |
| `Background_Job_Processor` | Buffered jobs, wait groups, OS signals, and graceful shutdown | `go run ./Background_Job_Processor` |
| `Flight_Price_Aggregator` | Concurrent price requests with a shared context timeout | `go run ./Flight_Price_Aggregator` |
| `Internet_Download_Manager` | Parallel part downloads and mutex-protected progress | `go run ./Internet_Download_Manager` |
| `Link_Checker` | Three-worker URL health checker with HTTP status and timeout reporting | `go run ./Link_Checker` |
| `Rate_Limiter` | Token-style request limiting with a buffered channel and ticker | `go run ./Rate_Limiter` |
| `The_Code_Evaluator_Engine` | Worker-based evaluation, per-job timeouts, and synchronized totals | `go run ./The_Code_Evaluator_Engine` |
| `workerPool` | Three workers processing a queue of image IDs | `go run ./workerPool` |
| `mini_project` | A small SaaS-style background job processor | `go run ./mini_project` |
| `LeaderBoard_Engine` | Placeholder for a future leaderboard implementation | `go run ./LeaderBoard_Engine` |

`Link_Checker` makes real outbound HTTP requests, so its results depend on your
internet connection and the current availability of the listed websites.

## Topic Examples

### Fundamentals

- `basic/`: hello world, variables, data types, arrays, slices, loops, `defer`,
  comma-ok checks, time-based `switch`, and small profile examples.
- `functions/`: arithmetic helpers, salary calculations, value parameters,
  developer data, multiple returns, and variadic arguments.
- `maps/`: age and pricing lookups.
- `structs/`: products, users, and inventory represented with custom types.
- `pointers/`: mutation through pointers, pointer receiver methods, cart totals,
  wallet balance, RAM upgrades, and game damage/weapon examples.
- `interfaces/`: behavior shared across storage providers, vehicles, and charging
  devices.

Run any one of these files from the repository root:

```sh
go run basic/hello_world.go
go run functions/calculate_salary.go
go run pointers/shopping_cart.go
go run interfaces/vehicle_start.go
```

### Error Handling

The `errors/` examples cover:

- Input validation for age, order totals, withdrawals, and pizza quantities
- Returning errors from functions
- Comparing known errors with `errors.Is`
- Adding context with `%w`
- Propagating database, file, inventory, payment, and user lookup failures

Example:

```sh
go run errors/database_error_wrapping.go
```

### Concurrency

The `goroutines/` and `channels/` directories progress from basic concurrent
functions to synchronization and communication patterns:

- Starting work with `go`
- Waiting for tasks with `sync.WaitGroup`
- Protecting shared state with `sync.Mutex`
- Sending data through buffered and unbuffered channels
- Closing and ranging over channels
- Coordinating multiple events with `select`
- Comparing sequential and concurrent URL checks
- Understanding an intentional unbuffered-channel deadlock example

Examples:

```sh
go run goroutines/mutex_protected_balance.go
go run channels/food_delivery_select.go
go run channels/url_status_checker.go
```

## HTTP API Examples

Every file in `apis/` is a separate server. Stop the current server with
`Ctrl+C` before starting another one that uses the same port.

| File | Port | Routes / purpose |
| --- | ---: | --- |
| `basic_http_server.go` | `8080` | `GET /api` basic response |
| `hello_endpoint.go` | `8080` | `GET /hello` |
| `json_response.go` | `8080` | `GET /user` JSON response |
| `ping_endpoint.go` | `9000` | `GET /ping` health-style response |
| `status_and_search.go` | `9000` | `GET /status`, `GET /search?q=...` |
| `profile_search.go` | `9000` | `GET /profile`, `GET /search?q=...` |
| `profile_and_greeting.go` | `8000` | `GET /user`, `GET /greet?name=...` |
| `developer_path_values.go` | `9000` | `GET /developer/{name}/{role}` |
| `pokemon_path_value.go` | `9000` | `GET /pokemon/{name}` |
| `login_json.go` | `9000` | `POST /login` JSON request handling |
| `in_memory_user_crud.go` | `9000` | `GET /users`, `POST /users` |
| `user_crud_with_path_id.go` | `9000` | List, create, fetch, and delete users |
| `postgres_user_api.go` | `9000` | PostgreSQL-backed `GET /users` and `POST /users` |

Start an API and call it from another terminal:

```sh
go run apis/status_and_search.go
curl 'http://localhost:9000/search?q=golang'
```

## Progressive Backend Module

`go-backend/` is its own Go module containing a step-by-step HTTP API series:

| Step | Topic |
| --- | --- |
| `01-http-server` | Basic server and port configuration |
| `02-structs-json` | Structs and JSON responses |
| `03-post-request` | Reading JSON from POST requests |
| `04-create-user` | Creating users in memory |
| `05-path-params` | Fetching a user by path ID |
| `06-put-delete` | Updating and deleting users |

Run a step from inside that module:

```sh
cd go-backend
go run ./06-put-delete
```

More details are available in [`go-backend/README.md`](./go-backend/README.md).

## PostgreSQL API Setup

Only `apis/postgres_user_api.go` requires an external dependency and database.
The connection string is read from `DATABASE_URL`; real credentials are never
stored in source code.

1. Start a PostgreSQL server and create or select a database.
2. Export the connection string using your actual host, port, user, and password.
3. Run the API.

```sh
export DATABASE_URL='postgres://postgres:your_password@localhost:5433/postgres?sslmode=disable'
go run apis/postgres_user_api.go
```

The application connects to PostgreSQL, creates a `users` table when necessary,
and starts the server on port `9000`.

Create and list users:

```sh
curl -X POST http://localhost:9000/users \
  -H 'Content-Type: application/json' \
  -d '{"name":"Shashwat","role":"Developer"}'

curl http://localhost:9000/users
```

See [`.env.example`](./.env.example) for the expected variable format. Local
`.env` files are excluded by `.gitignore`. This project does not automatically
load `.env`; export the variable in your shell or use your preferred environment
loader.

## Requirements and Setup

- Go `1.26.6` or newer
- Git for cloning the repository
- PostgreSQL only for the database-backed API
- Internet access only for examples that call public URLs

Clone and enter the repository:

```sh
git clone https://github.com/shxwat/golang.git
cd golang
go mod download
```

The root `go.mod` declares the module and the PostgreSQL driver dependency.
`go.sum` records dependency checksums so that builds are reproducible. Both files
should remain committed.

## Notes

- Most programs are intentionally small learning exercises, not production
  services.
- Several API examples use the same port and cannot run simultaneously without
  changing their port numbers.
- In-memory API data resets whenever the program restarts.
- The link-checking examples depend on external websites and may produce
  different results between runs.
- `LeaderBoard_Engine` is currently a placeholder and remains to be implemented.

## Suggested Learning Order

1. `basic` and `functions`
2. `maps`, `structs`, `pointers`, and `interfaces`
3. `errors`
4. `goroutines` and `channels`
5. `apis`
6. `go-backend`
7. Standalone concurrency projects
