
-- name: CreateTransaction :one
INSERT INTO transactions (
    id,
    user_id,
    amount,
    currency,
    category,
    note,
    occurred_at,
    created_at,
    updated_at
)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: ListTransactionsByUser :many
SELECT
    id,
    user_id,
    amount,
    currency,
    category,
    note,
    occurred_at,
    created_at,
    updated_at
FROM transactions
WHERE user_id = ?
ORDER BY occurred_at DESC
LIMIT ? OFFSET ?;


