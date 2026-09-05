# go-backend-practice

Small Go backend practice project using chi, pgx, sqlc, PostgreSQL, and Redis.

## Stack

- `chi` for HTTP routing
- `pgx` for PostgreSQL connections
- `sqlc` for type-safe SQL code generation
- `goose` style migration files
- `go-redis` for Redis

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
