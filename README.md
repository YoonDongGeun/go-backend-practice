# go-backend-practice

Small Go backend practice project using chi, pgx, sqlc, PostgreSQL, and Redis.

## Stack

- `chi` for HTTP routing
- `pgx` for PostgreSQL connections
- `sqlc` for type-safe SQL code generation
- `goose` style migration files
- `go-redis` for Redis

## Architecture

The example keeps a Java-style 3-layer shape, but uses Go package boundaries:

```text
cmd/api/main.go
  -> internal/server router
  -> internal/user handler        # controller layer
  -> internal/user service        # business layer
  -> internal/user repository     # persistence interface + postgres implementation
  -> internal/store               # sqlc generated query code
  -> PostgreSQL
```

For new features, copy the `internal/user` shape:

- `dto.go` defines request/response DTOs and service command inputs.
- `entity.go` defines the feature's domain entity.
- `handler.go` owns HTTP parsing, status codes, and JSON responses.
- `service.go` owns business rules and depends on a repository interface.
- `repository.go` adapts the database implementation to the service.
- `mapper.go` converts database rows/entities/DTOs.
- `sql/queries/*.sql` contains SQL that `sqlc` compiles into `internal/store`.

## Run locally

```bash
docker compose up -d
go install github.com/pressly/goose/v3/cmd/goose@latest
goose -dir migrations postgres "postgres://postgres:postgres@localhost:5432/go_backend_practice?sslmode=disable" up
go run ./cmd/api
```

## API

```bash
curl http://localhost:8080/healthz

curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"email":"yoon@example.com","name":"Yoon"}'

curl http://localhost:8080/api/v1/users
curl http://localhost:8080/api/v1/users/1
```

## Generate SQL code

```bash
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
sqlc generate
```

## Tests

```bash
go test ./...
```

The sample tests show two common backend test styles:

- `internal/user/service_test.go` tests service behavior with a fake repository.
- `internal/user/handler_test.go` tests controller/handler behavior with `httptest`.
