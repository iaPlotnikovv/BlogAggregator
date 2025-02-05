-- name: CreateUser :one
INSERT INTO users (id, name)
VALUES (
    $1,
    $2
)
RETURNING *;
-- name: ExistsUser :one
SELECT EXISTS(SELECT name FROM users WHERE name = $1);
-- name: ResetUsers :exec
DELETE FROM users;
-- name: GetUsers :many
SELECT name FROM users;
-- name: GetCurrentUserID :one
SELECT id FROM users WHERE name = $1;