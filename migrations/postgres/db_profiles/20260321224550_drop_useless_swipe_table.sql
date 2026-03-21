-- +goose Up
-- +goose StatementBegin
DROP TABLE swipes;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
CREATE TABLE swipes (
    id UUID PRIMARY KEY,
    user_id_1 UUID NOT NULL,
    user_id_2 UUID NOT NULL,
    desicion_1 BOOLEAN,
    desicion_2 BOOLEAN,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd
