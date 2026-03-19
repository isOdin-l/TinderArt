package repository

import (
	"github.com/cridenour/go-postgis"
	"github.com/google/uuid"
	"github.com/isOdin-l/TinderArt/pkg/db_models"
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

// DB ---> ENTITY

func FromGetProfileRowToEntity(db *db_models.GetProfileRow) *entities.Profile {
	return &entities.Profile{
		UserId:      FromPgUUID(&db.ID),
		Username:    db.Username,
		Surname:     db.Surname,
		Name:        db.Name,
		Email:       db.Email,
		Description: db.Description,
		Latitude:    db.Latitude.(float64),
		Longitude:   db.Longitude.(float64),
	}
}

func FromUpdateResultsToEntity(req *db_models.UpdateProfileRow) *entities.Profile {
	return &entities.Profile{
		UserId:      FromPgUUID(&req.ID),
		Username:    req.Username,
		Surname:     req.Surname,
		Name:        req.Name,
		Email:       req.Email,
		Description: req.Description,
	}
}

func FromPreferenceToEntity(req *db_models.Preference) *entities.Profile {
	return &entities.Profile{
		UserId:                   FromPgUUID(&req.ProfileID),
		PreferencesMaxDistMeters: int(req.MaxDistanceMeters),
	}
}

// ENTITY ---> DB

func FromEntityToCreateProfile(entity *entities.Profile) *db_models.CreateProfileParams {
	return &db_models.CreateProfileParams{
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

func FromEntityToUpdateProfileParams(entity *entities.UpdateProfile) *db_models.UpdateProfileParams {
	return &db_models.UpdateProfileParams{}
}

func FromEntityToCreatePrefParams(entity *entities.Profile) *db_models.CreatePreferencesParams {
	return &db_models.CreatePreferencesParams{}
}

func FromEntityToUpdatePref(entity *entities.Profile) *db_models.UpdatePreferencesParams {
	return &db_models.UpdatePreferencesParams{}
}

func FromEntityToCreatePhotos(entity *entities.Profile) *db_models.CreatePhotosParams {
	return &db_models.CreatePhotosParams{
		ProfileID: ToPgUUID(&entity.UserId),
		Urls:      entity.PhotoUrls,
	}
}
func FromEntityToDeletePhotos(entity *entities.Profile) *db_models.DeletePhotosParams {
	return &db_models.DeletePhotosParams{}
}
