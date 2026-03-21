package mappers

import (
	"github.com/isOdin-l/TinderArt/pkg/db_models"
	"github.com/isOdin-l/TinderArt/services/profile/internal/entities"
)

// DB ---> ENTITY
func FromPreferenceToEntity(req *db_models.Preference) *entities.Profile {
	return &entities.Profile{
		UserId:                   db_models.FromPgUUID(&req.ProfileID),
		PreferencesMaxDistMeters: int(req.MaxDistanceMeters),
	}
}

// ENTITY ---> DB
func FromEntityToCreatePrefParams(entity *entities.Profile) *db_models.CreatePreferencesParams {
	return &db_models.CreatePreferencesParams{
		ProfileID:         db_models.ToPgUUID(&entity.UserId),
		MaxDistanceMeters: int32(entity.PreferencesMaxDistMeters),
	}
}

func FromEntityToUpdatePref(entity *entities.Profile) *db_models.UpdatePreferencesParams {
	return &db_models.UpdatePreferencesParams{}
}
