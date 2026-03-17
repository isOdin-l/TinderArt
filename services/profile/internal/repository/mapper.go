package repository

import (
	"github.com/cridenour/go-postgis"
	"github.com/google/uuid"
	"github.com/isOdin-l/TinderArt/services/profile/internal/database/sqlc"
	"github.com/isOdin-l/TinderArt/services/profile/internal/entities"
	"github.com/jackc/pgx/v5/pgtype"
)

func ToPgUUID(v *uuid.UUID) pgtype.UUID {
	return pgtype.UUID{
		Bytes: *v,
		Valid: true,
	}
}

func FromPgUUID(v *pgtype.UUID) uuid.UUID {
	return uuid.UUID(v.Bytes)
}

func FromGetProfileRowToEntity(db *sqlc.GetProfileRow) *entities.Profile {
	return &entities.Profile{
		UserId:      FromPgUUID(&db.ID),
		Username:    db.Username,
		Surname:     db.Surname,
		Name:        db.Name,
		Email:       db.Email,
		Description: db.Description,
	}
}

func FromEntityToUpdateProfileParams(entity *entities.UpdateProfile) *sqlc.UpdateProfileParams {
	return &sqlc.UpdateProfileParams{}
}

func FromUpdateResultsToEntity(req *sqlc.UpdateProfileRow) *entities.Profile {
	return &entities.Profile{
		UserId:      FromPgUUID(&req.ID),
		Username:    req.Username,
		Surname:     req.Surname,
		Name:        req.Name,
		Email:       req.Email,
		Description: req.Description,
	}
}

func FromEntityToCreateProfile(entity *entities.Profile) *sqlc.CreateProfileParams {
	return &sqlc.CreateProfileParams{
		ID:          ToPgUUID(&entity.UserId),
		Username:    entity.Username,
		Password:    entity.Password,
		Surname:     entity.Surname,
		Name:        entity.Name,
		Email:       entity.Email,
		Description: entity.Description,
		Location:    postgis.Point{Y: entity.Latitude, X: entity.Longitude},
	}
}
