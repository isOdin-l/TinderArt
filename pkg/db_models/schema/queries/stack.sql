-- name: FindMatches :many
SELECT DISTINCT f2.profile_id
FROM fav_art_styles f1
JOIN fav_art_styles f2 ON f1.style_id = f2.style_id
JOIN profiles p2 ON p2.id = f2.profile_id
JOIN profiles p1 ON p1.id = $1
JOIN preferences pref ON pref.profile_id = $1
WHERE f1.profile_id = $1 AND f2.profile_id != $1 AND ST_DWithin(p1.location, p2.location, pref.max_distance);

-- name: GetAllProfiles :many
SELECT id FROM profiles;
