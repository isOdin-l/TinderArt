-- name: InsertSwipe :one
INSERT INTO swipes (id, user_id_1, user_id_2, desicion_1, desicion_2)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (user_id_1, user_id_2) DO NOTHING
RETURNING id;

-- name: UpdateSwipe :one
UPDATE swipes
SET
    desicion_2 = COALESCE($3, desicion_2)
WHERE user_id_1 = $1 AND user_id_2 = $2
RETURNING desicion_1, desicion_2;
