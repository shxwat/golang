# Go Backend Practice

Each topic has its own folder and `main.go`. Completed topics run independently.
The shared `go.mod` stays here in `go-backend`; do not run `go mod init` for each topic.

| Folder | Topic | Endpoint |
| --- | --- | --- |
| `01-http-server` | Basic HTTP server and port configuration | `/` |
| `02-structs-json` | Structs and a JSON response | `/users` |
| `03-post-request` | Read JSON data from a POST request | `GET /users`, `POST /users` |
| `04-create-user` | Create a user and append it to an in-memory list | `GET /users`, `POST /users` |
| `05-path-params` | Read a user ID from the URL path | `GET /users`, `POST /users`, `GET /users/{id}` |
| `06-put-delete` | Practice updating and deleting users (empty file; implementation pending) | Planned: `PUT /users/{id}`, `DELETE /users/{id}` |

From the `go-backend` directory, run one topic:

```sh
go run ./01-http-server
```

Stop it with Ctrl+C before running the next topic:

```sh
go run ./02-structs-json
```

After implementing the PUT and DELETE topic, run it with `go run ./06-put-delete`.

For a new topic, create a folder such as `07-topic-name` and put its
`main.go` inside. Run it with `go run ./07-topic-name` from this directory.
