
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

-- name: SumByDay :many
SELECT
    DATE(occurred_at) AS day,
    SUM(amount) AS total
FROM transactions
WHERE user_id = ?
  AND occurred_at BETWEEN ? AND ?
GROUP BY day
ORDER BY day;


-- -------------------------------------
-- Sum By Month
-- -------------------------------------
-- name: SumByMonth :many
SELECT
    strftime('%Y-%m', occurred_at) AS month,
    SUM(amount) AS total
FROM transactions
WHERE user_id = :user_id
  AND occurred_at BETWEEN :start_date AND :end_date
GROUP BY month
ORDER BY month;

-- -------------------------------------
-- Sum By Year
-- -------------------------------------
-- name: SumByYear :many
SELECT
    strftime('%Y', occurred_at) AS year,
    SUM(amount) AS total
FROM transactions
WHERE user_id = :user_id
  AND occurred_at BETWEEN :start_date AND :end_date
GROUP BY year
ORDER BY year;

