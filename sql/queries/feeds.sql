-- name: CreateFeed :one
INSERT INTO feeds (
    id, 
    created_at, 
    updated_at,
    name,
    url,
    user_id
)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: FeedsGetAll :many
SELECT * FROM feeds;


-- name: GetNextFeedsToFetch :many
SELECT *
    FROM feeds
    ORDER BY LAST_FETCHED_AT DESC NULLS FIRST
    LIMIT $1;

-- name: MarkFeedFetched :one
UPDATE feeds
SET LAST_FETCHED_AT = NOW()::TIMESTAMP, updated_at = NOW()::TIMESTAMP
WHERE id = $1
RETURNING *;
