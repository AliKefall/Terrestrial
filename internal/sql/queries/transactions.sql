-- name: CreateTransaction :one
INSERT INTO transactions (id, user_id, amount, currency, category, note, occurred_at, created_at, updated_at)
VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: ListTransactionsByUser :many
SELECT * FROM transactions 
WHERE user_id = ?
ORDER BY occurred_at DESC
LIMIT ? OFFSET ?;

-- name: SumByDay :many
SELECT date(occurred_at) AS day, SUM(amount) AS total
FROM TRANSACTIONS 
WHERE user_id = ? AND occurred_at BETWEEN ? AND ?
GROUP BY day
ORDER BY day;

-- name: SumByMonth :many
SELECT strftime('%Y-%m', occurred_at) AS month, SUM(amount) AS total
from transactions
where user_id = ? AND occurred_at BETWEEN ? AND ?
GROUP BY month
ORDER BY month;

-- name: SumByYear :many
SELECT strftime('%Y', occurred_at) AS year, SUM(amount) AS total
FROM transactions
WHERE user_id = ? AND occurred_at BETWEEN ? AND ?
GROUP BY year
ORDER BY year;
