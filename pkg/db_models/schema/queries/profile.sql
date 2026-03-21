-- name: CreateProfile :exec
INSERT INTO profiles (id, username, name, surname, email, password, description, location)
VALUES($1, $2, $3, $4, $5, $6, $7, ST_SetSRID(ST_MakePoint($8, $9), 4326));

-- name: GetProfile :one
SELECT id, username, name, surname, email, description,
    ST_X(location::GEOMETRY) AS longitude,
    ST_Y(location::GEOMETRY) AS latitude
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
RETURNING id, username, name, surname, email, description;

-- name: DeleteProfile :exec
DELETE FROM profiles WHERE id = $1;


-- name: CreatePreferences :exec
INSERT INTO preferences (profile_id, max_distance_meters)
VALUES ($1, $2);

-- name: UpdatePreferences :one
UPDATE preferences
SET max_distance_meters = COALESCE($2, max_distance_meters)
WHERE profile_id = $1 RETURNING profile_id, max_distance_meters;

-- name: CreateFavArtStyle :exec
INSERT INTO fav_art_styles (id, profile_id, style)
SELECT
    UNNEST(sqlc.arg(ids)::uuid[]),
    sqlc.arg(profile_id),
    UNNEST(sqlc.arg(styles)::text[])::art_style_enum;


-- name: CreatePhotos :exec
INSERT INTO photos (id, profile_id, url)
SELECT
    UNNEST(sqlc.arg(ids)::uuid[]),
    sqlc.arg(profile_id),
    UNNEST(sqlc.arg(urls)::text[]);

-- name: DeletePhotos :exec
DELETE FROM photos
WHERE profile_id = $1 AND id = ANY($2::uuid[]);
