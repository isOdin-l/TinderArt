-- name: CreateProfile :exec
INSERT INTO profiles (id, username, name, surname, email, password, description, location)
VALUES($1, $2, $3, $4, $5, $6, $7, $8);

-- name: GetProfile :one
SELECT
    id, username, name, surname, email, description,
    ST_X(location::GEOGRAPHY) AS longitude,
    ST_Y(location::GEOGRAPHY) AS latitude
FROM profiles WHERE id = $1;

-- name: UpdateProfile :one
UPDATE profiles
SET
    username = COALESCE($2, username),
    name = COALESCE($3, name),
    surname = COALESCE($4, surname),
    email = COALESCE($5, email),
    password = COALESCE($6, password),
    description = COALESCE($7, description),
    location = COALESCE(ST_SetSRID(ST_MakePoint(sqlc.arg(longitude), sqlc.arg(latitude)), 4326)::GEOGRAPHY, location)
WHERE id = $1
RETURNING id, username, name, surname, email, description;--, location;

-- name: DeleteProfile :exec
DELETE FROM profiles WHERE id = $1;
