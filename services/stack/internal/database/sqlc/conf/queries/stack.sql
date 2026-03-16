-- name: FindMatches :many
SELECT p.id FROM profiles p
JOIN profiles me ON me.id = $1
JOIN preferences pref ON pref.profile_id = $1
WHERE p.id != $1
AND p.age BETWEEN pref.min_age AND pref.max_age
AND ST_DWithin(me.location, p.location, pref.max_distance_meters);

-- name: GetAllProfiles :many
SELECT id FROM profiles;
