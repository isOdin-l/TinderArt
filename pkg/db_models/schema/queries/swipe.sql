-- name: InsertUpdateSwipe :one
INSERT INTO swipes (id, user_id_1, user_id_2, desicion_1, desicion_2)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (user_id_1, user_id_2) DO UPDATE
SET
    desicion_1 = COALESCE($4, swipes.desicion_1),
    desicion_2 = COALESCE($5, swipes.desicion_2)
RETURNING desicion_1, desicion_2;
