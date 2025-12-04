
-- name: CreateUser :one
INSERT INTO users (id, email, username, password, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?)
RETURNING *;


-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = ? LIMIT 1;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = ? LIMIT 1;
