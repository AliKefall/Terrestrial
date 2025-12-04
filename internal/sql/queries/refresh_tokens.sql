-- name: CreateRefreshToken :one
INSERT INTO refresh_tokens (token, user_id, created_at, updated_at, expires_at)
VALUES (?, ?, ?, ?, ?)
RETURNING *;

-- name: RevokeRefreshToken :exec
UPDATE refresh_tokens SET revoked_at = ?, updated_at = ? WHERE token = ?;

-- name: GetValidRefreshToken :one
SELECT * FROM refresh_tokens
WHERE token = ? AND revoked_at IS NULL AND expires_at > ?
LIMIT 1;
