-- name: PostCreate :one
INSERT INTO posts (
    id,
    created_at,
    updated_at,
    title,
    url,
    description,
    published_at,
    feed_id
) VALUES ( 
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
    )
RETURNING *;

-- name: PostsGetByUser :many
SELECT p.*
FROM posts p
LEFT JOIN feed_follows ff ON ff.feed_id = p.feed_id
WHERE user_id = $1
ORDER BY p.published_at DESC
LIMIT $2; 
