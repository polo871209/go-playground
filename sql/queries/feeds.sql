-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: GetNextFeedsToFetch :many
SELECT * FROM feeds
ORDER BY last_fatched_at ASC NULLS FIRST
LIMIT $1;

-- name: MarkFeedAsFetched :one
UPDATE feeds
SET last_fatched_at = NOW(),
updated_at = NOW()
WHERE id = $1
RETURNING *;