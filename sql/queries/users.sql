-- name: CreateUser :one
INSERT INTO users (email, name)
VALUES ($1, $2)
RETURNING id, email, name, created_at;

-- name: GetUser :one
SELECT id, email, name, created_at
FROM users
WHERE id = $1;

-- name: ListUsers :many
SELECT id, email, name, created_at
FROM users
ORDER BY id DESC
LIMIT $1 OFFSET $2;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
