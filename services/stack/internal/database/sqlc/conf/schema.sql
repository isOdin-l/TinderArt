CREATE TABLE profiles (
    id UUID PRIMARY KEY,
    username TEXT NOT NULL UNIQUE,
    name TEXT NOT NULL,
    surname TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    description TEXT ,
    location GEOGRAPHY(POINT, 4326) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE preferences (
    profile_id UUID PRIMARY KEY REFERENCES profiles(id) ON DELETE CASCADE,
    min_age INT NOT NULL,
    max_age INT NOT NULL,
    max_distance_meters INT NOT NULL
);

CREATE INDEX profiles_location_idx
ON profiles
USING GIST (location);
