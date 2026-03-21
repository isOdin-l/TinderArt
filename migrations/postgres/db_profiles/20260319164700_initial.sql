-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'art_style_enum') THEN
        CREATE TYPE art_style_enum AS ENUM (
            'realism',
            'minimalism',
            'futurism',
            'anarchism',
            'cubism',
            'surrealism',
            'impressionism',
            'expressionism',
            'constructivism',
            'dadaism',
            'photorealism',
            'romanticism',
            'cyberpunk'
        );
    END IF;
END$$;

CREATE TABLE IF NOT EXISTS profiles (
    id UUID PRIMARY KEY,
    username TEXT NOT NULL UNIQUE,
    name TEXT NOT NULL,
    surname TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    description TEXT NOT NULL,
    location GEOGRAPHY(POINT, 4326) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS preferences (
    profile_id UUID PRIMARY KEY REFERENCES profiles(id) ON DELETE CASCADE,
    max_distance_meters INT NOT NULL
);


CREATE TABLE IF NOT EXISTS fav_art_styles (
    id UUID PRIMARY KEY,
    profile_id UUID REFERENCES profiles(id) ON DELETE CASCADE,
    style art_style_enum NOT NULL,

    created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE IF NOT EXISTS jwt_tokens (
    id UUID PRIMARY KEY REFERENCES profiles(id) ON DELETE CASCADE,
    refresh_token TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS swipes (
    id UUID PRIMARY KEY,
    user_id_1 UUID NOT NULL,
    user_id_2 UUID NOT NULL,
    desicion_1 BOOLEAN,
    desicion_2 BOOLEAN,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS photos(
    id UUID PRIMARY KEY,
    profile_id UUID REFERENCES profiles(id) ON DELETE CASCADE,
    url TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);


-- Indexes
CREATE INDEX IF NOT EXISTS profiles_location_idx ON profiles USING GIST (location);
CREATE INDEX IF NOT EXISTS idx_profiles_username ON profiles(username);
CREATE INDEX IF NOT EXISTS idx_jwt_tokens_refresh ON jwt_tokens(refresh_token);
CREATE INDEX IF NOT EXISTS idx_swipes_user1 ON swipes(user_id_1);
CREATE INDEX IF NOT EXISTS idx_swipes_user2 ON swipes(user_id_2);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE profiles;
DROP TABLE photos;
DROP TABLE swipes;
DROP TABLE jwt_tokens;
DROP TABLE fav_art_styles;
DROP TABLE preferences;
-- +goose StatementEnd
