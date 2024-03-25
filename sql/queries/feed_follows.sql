-- name: CreateFeedFollow :one
INSERT INTO feed_follows (
    id,
    feed_id,
    user_id,
    created_at,
    updated_at
)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: DeleteFeedFollow :execrows
DELETE FROM feed_follows
WHERE id=$1
RETURNING *;

-- name: GetAllFeedFollows :many
SELECT *
    FROM feed_follows
    WHERE user_id=$1;
