package repository

import (
	"github.com/cridenour/go-postgis"
	"github.com/google/uuid"
	"github.com/isOdin-l/TinderArt/pkg/db_models"
	"github.com/isOdin-l/TinderArt/services/profile/internal/entities"
	"github.com/jackc/pgx/v5/pgtype"
)

func ToPgUUID(v *uuid.UUID) pgtype.UUID {
	if v == nil {
		return pgtype.UUID{
			Valid: false,
		}
	}

	return pgtype.UUID{
		Bytes: *v,
		Valid: true,
	}
}

func FromPgUUID(v *pgtype.UUID) uuid.UUID {
	return uuid.UUID(v.Bytes)
}

func ToPgText(v *string) pgtype.Text {
	if v == nil {
		return pgtype.Text{
			Valid: false,
		}
	}

	return pgtype.Text{
		String: *v,
		Valid:  true,
	}
}

func FromPgText(v *pgtype.Text) string {
	return v.String
}

// DB ---> ENTITY

func FromGetProfileRowToEntity(db *db_models.GetProfileRow) *entities.Profile {
	return &entities.Profile{
		UserId:      FromPgUUID(&db.ID),
		Username:    FromPgText(&db.Username),
		Surname:     FromPgText(&db.Surname),
		Name:        FromPgText(&db.Name),
		Email:       FromPgText(&db.Email),
		Description: FromPgText(&db.Description),
		Latitude:    db.Latitude.(float64),
		Longitude:   db.Longitude.(float64),
	}
}

func FromUpdateResultsToEntity(req *db_models.UpdateProfileRow) *entities.Profile {
	return &entities.Profile{
		UserId:      FromPgUUID(&req.ID),
		Username:    FromPgText(&req.Username),
		Surname:     FromPgText(&req.Surname),
		Name:        FromPgText(&req.Name),
		Email:       FromPgText(&req.Email),
		Description: FromPgText(&req.Description),
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
		Username:    ToPgText(&entity.Username),
		Password:    ToPgText(&entity.Password),
		Surname:     ToPgText(&entity.Surname),
		Name:        ToPgText(&entity.Name),
		Email:       ToPgText(&entity.Email),
		Description: ToPgText(&entity.Description),
		Location:    postgis.Point{Y: entity.Latitude, X: entity.Longitude},
	}
}

func FromEntityToUpdateProfileParams(entity *entities.UpdateProfile) *db_models.UpdateProfileParams {
	return &db_models.UpdateProfileParams{
		ID:          ToPgUUID(&entity.UserId),
		Username:    ToPgText(entity.Username),
		Name:        ToPgText(entity.Name),
		Surname:     ToPgText(entity.Surname),
		Email:       ToPgText(entity.Email),
		Password:    ToPgText(entity.Password),
		Description: ToPgText(entity.Description),
		Longitude:   entity.Longitude,
		Latitude:    entity.Latitude,
	}
}

func FromEntityToCreatePrefParams(entity *entities.Profile) *db_models.CreatePreferencesParams {
	return &db_models.CreatePreferencesParams{}
}

func FromEntityToUpdatePref(entity *entities.Profile) *db_models.UpdatePreferencesParams {
	return &db_models.UpdatePreferencesParams{}
}

//	func FromEntityToCreatePhotos(entity *entities.Profile) *db_models.CreatePhotosParams {
//		return &db_models.CreatePhotosParams{
//			ProfileID: ToPgUUID(&entity.UserId),
//			Urls:      entity.PhotoUrls,
//		}
//	}
func FromEntityToDeletePhotos(entity *entities.Profile) *db_models.DeletePhotosParams {
	return &db_models.DeletePhotosParams{}
}
