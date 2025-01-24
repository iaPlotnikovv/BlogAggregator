-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;
-- name: ExistsUser :one
SELECT EXISTS(SELECT name FROM users WHERE name = $1);
-- name: ResetUsers :exec
DELETE FROM users;
-- name: GetUsers :many
SELECT name FROM users;