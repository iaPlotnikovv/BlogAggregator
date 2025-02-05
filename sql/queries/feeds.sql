-- name: CreateFeed :one
INSERT INTO feeds (id, user_id, name, url)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: ListFeed :many
SELECT users.name, feeds.name, feeds.url
FROM feeds
INNER JOIN users
ON feeds.user_id = users.id;