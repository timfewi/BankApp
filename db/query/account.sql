-- name: CreateAccount :one 
INSERT INTO account (
    owner,
    balance,
    currency
    ) VALUES (
    $1,
    $2,
    $3
    ) RETURNING *;

-- name: GetAccount :one
SELECT * FROM account
WHERE id = $1
LIMIT 1;

-- name: GetAccounts :many
SELECT * FROM account
WHERE owner = $1
ORDER BY id DESC
OFFSET $2
LIMIT $3;

-- name: UpdateAccountBalance :one
UPDATE account
SET balance = $2
WHERE id = $1
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM account
WHERE id = $1;

