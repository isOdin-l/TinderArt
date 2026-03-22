#!/bin/sh
set -e

PROFILE_USER=$(cat /run/secrets/profile_user)
PROFILE_PASSWORD=$(cat /run/secrets/profile_password)
PROFILE_DB=$(cat /run/secrets/profile_db)

SWIPES_USER=$(cat /run/secrets/swipes_user)
SWIPES_PASSWORD=$(cat /run/secrets/swipes_password)
SWIPES_DB=$(cat /run/secrets/swipes_db)

goose -dir /migrations/postgres/db_profiles postgres "postgresql://${PROFILE_USER}:${PROFILE_PASSWORD}@db_profiles:5432/${PROFILE_DB}?sslmode=disable" up
goose -dir /migrations/postgres/db_swipes postgres "postgresql://${SWIPES_USER}:${SWIPES_PASSWORD}@db_swipes:5432/${SWIPES_DB}?sslmode=disable" up
