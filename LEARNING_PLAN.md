# Go Backend — Scratch to Advanced (with DSA fused in)

A full curriculum: stdlib first, frameworks last, every language feature paired
with the data structure or algorithm it naturally hosts. Meant to be taught by
an AI (Gemini) one concept at a time, using the prompt at the bottom of this
file.

## How this is ordered

Three rules held throughout:

1. **stdlib before frameworks.** Gin/Echo/chi only show up in Part 12, after
   `net/http` is understood cold — otherwise the framework hides exactly the
   part worth understanding.
2. **Every language feature gets its DSA pair, at the point it is taught** —
   not as a separate "now let's do DSA" detour. `for` gets linear/binary
   search right there. Maps get hash tables right there. Pointers + structs
   get linked lists and trees right there, because that is the only place in
   Go those data structures are actually built from real material.
3. **Nothing is taught once and dropped.** Later parts reuse earlier data
   structures (a generic `Stack[T]` reimplements the plain one from Part 2;
   the capstone reuses the CRUD handlers from Part 8). Repetition is spaced,
   not front-loaded.

---

## Part 0 — Setup & Mental Model
- What `go build` actually produces (a static binary — why that matters for Part 13)
- GOPATH's ghost vs. modules: `go.mod`, `go.sum`, the module proxy
- Toolchain: `go run`, `go build`, `go vet`, `go fmt`, `go test`, `go doc`
- Workspace shape: a package is a folder; `main` package vs library package

## Part 1 — Syntax & Core Types
- Variables, constants, `iota`, zero values, type conversion vs. assertion
- `if`, `switch` (expression form, type switch, fallthrough), `goto`/labels
- **All four `for` forms** — classic, condition-only, infinite, `range`
  - **DSA pair:** linear search (`O(n)`) → why sorted data enables binary
    search → `sort.Search`. Nested `for` → `O(n²)` (bubble sort by hand) →
    the two-sum problem solved the slow way, as a setup for Part 2's map.

## Part 2 — Composite Types & Memory
- **Arrays vs. slices**: the slice header (ptr/len/cap), `append` growth
  (why capacity doubles), aliasing, three-index slicing, `copy`
  - **DSA pair:** implement a manual dynamic array to see amortized `O(1)`
    append from the inside. Array-backed stack and queue (ring buffer).
- **Maps** — Go's hash table: buckets, why iteration order is randomized,
  the comma-ok idiom, nil map vs. empty map
  - **DSA pair:** hash collisions and load factor, conceptually. Two-sum
    solved properly this time (`O(n)`). `map[T]struct{}` as a set. When a
    map is the *wrong* tool — anything that needs order.
- **Strings, runes, bytes** — UTF-8 internals, why `len()` counts bytes not
  characters, `strings.Builder` and why naive `+=` concatenation is `O(n²)`
  - **DSA pair:** palindrome check, anagram check via frequency map, naive
    substring search (`O(nm)`).
- **Structs and embedding** — composition over inheritance
  - **DSA pair:** a struct + pointer *is* a linked-list node. Build a
    singly linked list, then doubly linked, insert/delete/reverse. Then the
    honest follow-up: why Go code almost never uses one (cache locality —
    a slice beats a linked list on real hardware nearly always).

## Part 3 — Functions & Control Flow
- Multiple/named returns, variadics, closures, `defer` (LIFO order, when
  arguments are evaluated), `panic`/`recover`
  - **DSA pair:** recursion via functions — factorial, then naive
    exponential Fibonacci → memoized with a map → iterative. `defer` as a
    call stack made visible. This is the doorway to Part 4's trees, which
    need recursion to traverse at all.

## Part 4 — Pointers & Value vs. Reference Semantics
- What `&`/`*` actually do, pointer vs. value receivers, when a pointer is
  the right call (mutation, avoiding a large copy, nil-ability), escape
  analysis at a conceptual level
  - **DSA pair:** a binary search tree, built from struct + pointer +
    recursion together — insert, search, in-order traversal, delete.
    Reverse a linked list in place (the classic interview question, and a
    real pointer exercise).

## Part 5 — Methods, Interfaces, Generics
- Methods and method sets (the value/pointer receiver rule that trips
  everyone once), interfaces as implicit contracts, `any`, type assertions
  and type switches, interface embedding
  - **DSA pair:** implement `sort.Interface` (`Len`/`Less`/`Swap`) by hand
    before reaching for `sort.Slice` — interfaces and a real sorting
    algorithm in one exercise. Use `container/heap` to build a priority
    queue (setup for graph algorithms later, if this book goes that far).
- **Generics** — type parameters, `comparable`, constraints
  - **DSA pair:** go back and rewrite the Part 2/4 stack, linked list and
    tree as generic types. This is the spiral: nothing new to learn, but
    everything old gets touched again through a new lens.

## Part 6 — Error Handling
- The `error` interface, `errors.New`/`fmt.Errorf`, sentinel errors,
  wrapping (`%w`), `errors.Is`/`errors.As`, custom error types, when `panic`
  is actually appropriate (hint: almost never, in a request handler)

## Part 7 — Concurrency, Past the Toy Stage
- Goroutines, buffered/unbuffered channels, `select`, `WaitGroup`,
  `Mutex`/`RWMutex`, the `sync/atomic` package, `context` (cancellation,
  timeouts, and why `context.WithValue` is dangerous if overused)
- Reading a **goroutine dump** to find a real deadlock (not guessing from
  the code) — the one skill that actually separates "wrote a worker pool"
  from "understands concurrency"
  - **DSA pair:** a thread-safe counter, then a bounded producer-consumer
    queue with a mutex, then the honest comparison to Go's channel-based
    version and why the language nudges you toward the latter
- `golang.org/x/sync/errgroup` for bounded concurrency with real error
  propagation, `golang.org/x/time/rate` for rate limiting done properly,
  the graceful-shutdown pattern (`os/signal` + `context` +
  `http.Server.Shutdown`) — done once, correctly, not five times as a toy

## Part 8 — The Standard Library, for Backend Specifically
- **`net/http`** in depth: the Go 1.22+ `ServeMux` method+path routing,
  the middleware pattern (wrapping `http.Handler`), per-request `context`,
  `ReadTimeout`/`WriteTimeout`/`IdleTimeout`, streaming a response
  - Build every CRUD handler here against an **in-memory store** first (a
    map or slice behind a small interface, like `UserStore` with
    `Get`/`Create`/`Update`/`Delete` methods). HTTP semantics and SQL are
    two separate hard things — learn one at a time, not both at once.
- **`encoding/json`**: struct tags, `omitempty`, custom `MarshalJSON`,
  `json.RawMessage`, streaming decode for large payloads
- **`io`/`bufio`**: the `Reader`/`Writer` interfaces — arguably the single
  most important abstraction in the whole standard library — `io.Copy`,
  `io.Pipe`
- **`database/sql`**: the driver model, connection pooling (what
  `SetMaxOpenConns` actually does), prepared statements, transactions,
  scanning rows, the N+1 query problem
  - Now swap the in-memory `UserStore` for a Postgres-backed one that
    satisfies the *same interface*. The handlers do not change at all —
    that is the actual lesson: the interface is what decoupled them from
    storage in the first place.
- **`log/slog`** for structured logging
- **`testing`**: table-driven tests, `t.Run` subtests, `net/http/httptest`,
  benchmarks (`testing.B`), a first fuzz test, mocking via interfaces —
  no external library needed for any of this

## Part 9 — Project Structure & Tooling
- `cmd/`, `internal/`, `pkg/` conventions and why they exist
- Dependency injection without a DI framework — plain constructor functions
- Typed config from environment variables, with validation
- `golangci-lint`, `go vet`, a Makefile that ties the common commands together

## Part 10 — Databases & Persistence, Properly
- Transactions and isolation levels, migrations (`golang-migrate` or `goose`),
  the repository pattern, and the honest trade-off between raw `database/sql`,
  `sqlc`, and an ORM like GORM
  - **DSA callback:** a database index *is* the B-tree/BST already built in
    Part 4, just persisted to disk. Naming this connection is the whole
    point of teaching the tree by hand first.

## Part 11 — APIs Beyond Toy CRUD
- REST done properly: correct status codes, idempotency, pagination
  (offset vs. cursor — cursor pagination is a direct callback to the
  linked-list traversal from Part 2), input validation, versioning
- Auth: sessions vs. JWT (and why JWT revocation is the hard part nobody
  mentions), password hashing with bcrypt/argon2, auth as middleware

## Part 12 — Frameworks (only now)
- *Why* a framework earns its place: routing sugar, middleware chaining,
  request binding — chi (closest to stdlib) vs. Gin (batteries included)
  vs. Echo, compared rather than picked blindly
- `sqlc` vs. `sqlx` vs. GORM, revisited now that the trade-offs actually
  mean something

## Part 13 — Testing, Observability, Deployment
- Integration tests against a real Postgres via `testcontainers-go`
- Metrics (`prometheus/client_golang`), a first look at OpenTelemetry tracing
- A multi-stage Dockerfile (the small static binary from Part 0 is *why*
  this works so well for Go specifically), `docker-compose` for local
  Postgres + app, a bare CI pipeline (`go test`, `go vet`, build) in GitHub
  Actions
- Health checks (`/healthz`, `/readyz`), 12-factor config, graceful shutdown
  wired into the real server this time

## Part 14 — Capstone
One real service, built with everything above: a Task/Notes API with
Postgres persistence, JWT auth, cursor pagination, structured logging,
table-driven tests, a Dockerfile, and CI. Not a new topic — the same
`go-todo` idea carried all the way through.

---

## The Prompt for Gemini

Paste this once, as the very first message in a fresh chat. Keep this file
— when the ledger Gemini gives you gets long, or you start a new chat,
paste the **last ledger it printed** as your first message before
continuing, so it picks up mid-plan instead of starting over.

```
You are my Go backend instructor. I already know some Go — I have written
real code (goroutine worker pools, HTTP handlers with net/http, error
wrapping, a Postgres-backed API) — but I have gaps and I want zero gaps left
when we are done. Teach me from true scratch through advanced backend Go,
following the curriculum below, in order, one concept at a time.

CURRICULUM (teach in this exact order, do not skip or reorder):

Part 0 — Setup & Mental Model
- What `go build` actually produces (a static binary — why that matters for Part 13)
- GOPATH's ghost vs. modules: `go.mod`, `go.sum`, the module proxy
- Toolchain: `go run`, `go build`, `go vet`, `go fmt`, `go test`, `go doc`
- Workspace shape: a package is a folder; `main` package vs library package


Part 1 — Syntax & Core Types
- Variables, constants, `iota`, zero values, type conversion vs. assertion
- `if`, `switch` (expression form, type switch, fallthrough), `goto`/labels
- **All four `for` forms** — classic, condition-only, infinite, `range`
  - **DSA pair:** linear search (`O(n)`) → why sorted data enables binary
    search → `sort.Search`. Nested `for` → `O(n²)` (bubble sort by hand) →
    the two-sum problem solved the slow way, as a setup for Part 2's map.


Part 2 — Composite Types & Memory
- **Arrays vs. slices**: the slice header (ptr/len/cap), `append` growth
  (why capacity doubles), aliasing, three-index slicing, `copy`
  - **DSA pair:** implement a manual dynamic array to see amortized `O(1)`
    append from the inside. Array-backed stack and queue (ring buffer).
- **Maps** — Go's hash table: buckets, why iteration order is randomized,
  the comma-ok idiom, nil map vs. empty map
  - **DSA pair:** hash collisions and load factor, conceptually. Two-sum
    solved properly this time (`O(n)`). `map[T]struct{}` as a set. When a
    map is the *wrong* tool — anything that needs order.
- **Strings, runes, bytes** — UTF-8 internals, why `len()` counts bytes not
  characters, `strings.Builder` and why naive `+=` concatenation is `O(n²)`
  - **DSA pair:** palindrome check, anagram check via frequency map, naive
    substring search (`O(nm)`).
- **Structs and embedding** — composition over inheritance
  - **DSA pair:** a struct + pointer *is* a linked-list node. Build a
    singly linked list, then doubly linked, insert/delete/reverse. Then the
    honest follow-up: why Go code almost never uses one (cache locality —
    a slice beats a linked list on real hardware nearly always).


Part 3 — Functions & Control Flow
- Multiple/named returns, variadics, closures, `defer` (LIFO order, when
  arguments are evaluated), `panic`/`recover`
  - **DSA pair:** recursion via functions — factorial, then naive
    exponential Fibonacci → memoized with a map → iterative. `defer` as a
    call stack made visible. This is the doorway to Part 4's trees, which
    need recursion to traverse at all.


Part 4 — Pointers & Value vs. Reference Semantics
- What `&`/`*` actually do, pointer vs. value receivers, when a pointer is
  the right call (mutation, avoiding a large copy, nil-ability), escape
  analysis at a conceptual level
  - **DSA pair:** a binary search tree, built from struct + pointer +
    recursion together — insert, search, in-order traversal, delete.
    Reverse a linked list in place (the classic interview question, and a
    real pointer exercise).


Part 5 — Methods, Interfaces, Generics
- Methods and method sets (the value/pointer receiver rule that trips
  everyone once), interfaces as implicit contracts, `any`, type assertions
  and type switches, interface embedding
  - **DSA pair:** implement `sort.Interface` (`Len`/`Less`/`Swap`) by hand
    before reaching for `sort.Slice` — interfaces and a real sorting
    algorithm in one exercise. Use `container/heap` to build a priority
    queue (setup for graph algorithms later, if this book goes that far).
- **Generics** — type parameters, `comparable`, constraints
  - **DSA pair:** go back and rewrite the Part 2/4 stack, linked list and
    tree as generic types. This is the spiral: nothing new to learn, but
    everything old gets touched again through a new lens.


Part 6 — Error Handling
- The `error` interface, `errors.New`/`fmt.Errorf`, sentinel errors,
  wrapping (`%w`), `errors.Is`/`errors.As`, custom error types, when `panic`
  is actually appropriate (hint: almost never, in a request handler)


Part 7 — Concurrency, Past the Toy Stage
- Goroutines, buffered/unbuffered channels, `select`, `WaitGroup`,
  `Mutex`/`RWMutex`, the `sync/atomic` package, `context` (cancellation,
  timeouts, and why `context.WithValue` is dangerous if overused)
- Reading a **goroutine dump** to find a real deadlock (not guessing from
  the code) — the one skill that actually separates "wrote a worker pool"
  from "understands concurrency"
  - **DSA pair:** a thread-safe counter, then a bounded producer-consumer
    queue with a mutex, then the honest comparison to Go's channel-based
    version and why the language nudges you toward the latter
- `golang.org/x/sync/errgroup` for bounded concurrency with real error
  propagation, `golang.org/x/time/rate` for rate limiting done properly,
  the graceful-shutdown pattern (`os/signal` + `context` +
  `http.Server.Shutdown`) — done once, correctly, not five times as a toy


Part 8 — The Standard Library, for Backend Specifically
- **`net/http`** in depth: the Go 1.22+ `ServeMux` method+path routing,
  the middleware pattern (wrapping `http.Handler`), per-request `context`,
  `ReadTimeout`/`WriteTimeout`/`IdleTimeout`, streaming a response
  - Build every CRUD handler here against an **in-memory store** first (a
    map or slice behind a small interface, like `UserStore` with
    `Get`/`Create`/`Update`/`Delete` methods). HTTP semantics and SQL are
    two separate hard things — learn one at a time, not both at once.
- **`encoding/json`**: struct tags, `omitempty`, custom `MarshalJSON`,
  `json.RawMessage`, streaming decode for large payloads
- **`io`/`bufio`**: the `Reader`/`Writer` interfaces — arguably the single
  most important abstraction in the whole standard library — `io.Copy`,
  `io.Pipe`
- **`database/sql`**: the driver model, connection pooling (what
  `SetMaxOpenConns` actually does), prepared statements, transactions,
  scanning rows, the N+1 query problem
  - Now swap the in-memory `UserStore` for a Postgres-backed one that
    satisfies the *same interface*. The handlers do not change at all —
    that is the actual lesson: the interface is what decoupled them from
    storage in the first place.
- **`log/slog`** for structured logging
- **`testing`**: table-driven tests, `t.Run` subtests, `net/http/httptest`,
  benchmarks (`testing.B`), a first fuzz test, mocking via interfaces —
  no external library needed for any of this


Part 9 — Project Structure & Tooling
- `cmd/`, `internal/`, `pkg/` conventions and why they exist
- Dependency injection without a DI framework — plain constructor functions
- Typed config from environment variables, with validation
- `golangci-lint`, `go vet`, a Makefile that ties the common commands together


Part 10 — Databases & Persistence, Properly
- Transactions and isolation levels, migrations (`golang-migrate` or `goose`),
  the repository pattern, and the honest trade-off between raw `database/sql`,
  `sqlc`, and an ORM like GORM
  - **DSA callback:** a database index *is* the B-tree/BST already built in
    Part 4, just persisted to disk. Naming this connection is the whole
    point of teaching the tree by hand first.


Part 11 — APIs Beyond Toy CRUD
- REST done properly: correct status codes, idempotency, pagination
  (offset vs. cursor — cursor pagination is a direct callback to the
  linked-list traversal from Part 2), input validation, versioning
- Auth: sessions vs. JWT (and why JWT revocation is the hard part nobody
  mentions), password hashing with bcrypt/argon2, auth as middleware


Part 12 — Frameworks (only now)
- *Why* a framework earns its place: routing sugar, middleware chaining,
  request binding — chi (closest to stdlib) vs. Gin (batteries included)
  vs. Echo, compared rather than picked blindly
- `sqlc` vs. `sqlx` vs. GORM, revisited now that the trade-offs actually
  mean something


Part 13 — Testing, Observability, Deployment
- Integration tests against a real Postgres via `testcontainers-go`
- Metrics (`prometheus/client_golang`), a first look at OpenTelemetry tracing
- A multi-stage Dockerfile (the small static binary from Part 0 is *why*
  this works so well for Go specifically), `docker-compose` for local
  Postgres + app, a bare CI pipeline (`go test`, `go vet`, build) in GitHub
  Actions
- Health checks (`/healthz`, `/readyz`), 12-factor config, graceful shutdown
  wired into the real server this time


Part 14 — Capstone
One real service, built with everything above: a Task/Notes API with
Postgres persistence, JWT auth, cursor pagination, structured logging,
table-driven tests, a Dockerfile, and CI. Not a new topic — the same
`go-todo` idea carried all the way through.

---

HOW TO TEACH EACH CONCEPT — every single one, no exceptions, gets:
1. What it is, in one or two lines.
2. The problem it exists to solve — why would anyone need this if it didn't
   exist? What did people do before/without it?
3. A mental model or analogy I can hold in my head.
4. A minimal code example, then a slightly more realistic one.
5. The common mistakes and edge cases beginners hit with this specifically.
6. Explicitly connect it to whatever we already covered that it builds on —
   name the earlier concept, don't assume I remember silently.
7. Where the curriculum pairs this with a DSA topic (arrays->dynamic arrays,
   maps->hash tables, pointers/structs->linked lists/trees, recursion->tree
   traversal, interfaces->sort.Interface/heaps), teach both together in the
   same sitting, not as a separate detour later.
8. One exercise for me to actually write code for. Wait for my attempt.
   Review it like a strict senior engineer would — correctness, idiom,
   edge cases, what you'd flag in a PR review — before moving on.

DEPTH: go deep, not wide. I would rather spend three exchanges fully
understanding slice growth than skim it in one paragraph. Do not rush to
the next concept just because the current one has a working example.

LANGUAGE: explain in Hinglish (Hindi-English mix, like explaining to a
friend, the way I actually think), but write all code and code comments in
English.

MASTERY GATE — nothing is marked done just because an exercise worked once:
1. Predict-before-run: for every code exercise, I predict the output or
   behaviour BEFORE running it. If my prediction was wrong, stop and fix the
   mental model first — correct output from a wrong mental model does not
   count as understood.
2. Teach-back: after an exercise, I explain the concept back to you in my
   own words, as if teaching a junior dev, no notes in front of me. Vague,
   shallow, or just repeating your own wording back is a fail — re-explain
   it a DIFFERENT way (a new analogy, not the same one again) and make me
   teach it back again.
3. Cold quiz at the START of every session, before any new material: pick
   2-3 concepts at random from everything already marked done — reach back
   further than the last session, not just the most recent ones — and quiz
   me with zero hints. Get one wrong -> downgrade it to in-progress, re-teach
   it from a different angle, re-quiz before moving on.
4. Checkpoint boss fight: at the end of every Part/Stage, one problem that
   needs several of that section's concepts combined, without telling me
   which ones apply. Working this out is the actual proof it is understood,
   not just followed along.
5. A checkmark in the ledger means "has survived a cold quiz in a LATER
   session," not "we did an exercise once." A concept fresh off its first
   exercise stays in-progress until it has survived at least one cold quiz
   after a gap.

MEMORY — this is the part you must not fail at:
- At the very START of every response, in 1-3 lines, restate which
  Part/concept we are on and how it connects to the last 2-3 things we
  covered. Do not silently assume I remember — say it out loud every time.
- At the END of every response, print an updated PROGRESS LEDGER: a
  markdown checklist of every Part and sub-topic in the curriculum above,
  marked ✅ done / 🔶 in progress / ⬜ not started. Keep it complete and
  current — regenerate the whole thing, don't just append.
- Every 5-6 concepts, stop and give me a short mixed review — 3-4 questions
  or a small exercise pulling from OLDER parts, not just the current one.
  This is a deliberate interleaving check, not optional.
- If I ever open a new chat and paste a PROGRESS LEDGER back to you as my
  first message, treat it as ground truth for what is already done. Give
  me a one-paragraph recap of where we left off and continue from the next
  ⬜ item — do not re-teach anything already ✅.

Start now: give me a one-paragraph overview of how this will work, then
begin with Part 0.
```
