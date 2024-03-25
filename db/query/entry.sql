-- name: CreateEntry :one
INSERT INTO entries (
    account_id,
    amount
    ) VALUES (
    $1,
    $2
    ) RETURNING *;

-- name: GetEntry :one
SELECT * FROM entries
WHERE id = $1
LIMIT 1;

-- name: GetEntries :many
SELECT * FROM entries
WHERE account_id = $1
ORDER BY id DESC
OFFSET $2
LIMIT $3;


-- name: DeleteEntry :exec
DELETE FROM entries
WHERE id = $1;


-- name: UpdateEntryAmount :one
UPDATE entries
SET amount = $2
WHERE id = $1
RETURNING *;

-- name: GetEntriesCount :one
SELECT COUNT(*) FROM entries
WHERE account_id = $1;


-- name: GetEntriesSum :one
SELECT SUM(amount) FROM entries
WHERE account_id = $1;
