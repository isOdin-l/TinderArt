-- name: GetUserByUsername :one
SELECT id, password FROM profiles WHERE username = $1;

-- name: SaveRefreshToken :exec
INSERT INTO jwt_tokens (id, refresh_token)
VALUES ($1, $2)
ON CONFLICT (id) DO UPDATE SET refresh_token = $2;
