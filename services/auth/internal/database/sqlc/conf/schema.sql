CREATE TABLE profiles (
    id UUID PRIMARY KEY,
    username TEXT NOT NULL UNIQUE,
    name TEXT NOT NULL,
    surname TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    description TEXT,
    latitude FLOAT,
    longitude FLOAT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE jwt_tokens (
    id UUID PRIMARY KEY REFERENCES profiles(id) ON DELETE CASCADE,
    refresh_token TEXT NOT NULL
);

-- Indexes for faster lookups
CREATE INDEX idx_profiles_username ON profiles(username);
CREATE INDEX idx_profiles_email ON profiles(email);
CREATE INDEX idx_jwt_tokens_refresh ON jwt_tokens(refresh_token);