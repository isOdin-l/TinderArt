CREATE TABLE swipes (
    id UUID PRIMARY KEY,
    user_id_1 UUID NOT NULL,
    user_id_2 UUID NOT NULL,
    desicion_1 BOOLEAN,
    desicion_2 BOOLEAN,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Index for user's swipes (search by user in either position)
CREATE INDEX idx_swipes_user1 ON swipes(user_id_1);
CREATE INDEX idx_swipes_user2 ON swipes(user_id_2);
