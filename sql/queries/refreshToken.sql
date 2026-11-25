-- name: RefreshToken :one
INSERT INTO REFRESH_TOKENS(token, created_at, updated_at, user_id, expires_at)
VALUES ($1, Now(),  NOW(), $2, $3)
RETURNING *;