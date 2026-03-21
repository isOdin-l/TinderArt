package mappers

import (
	"github.com/isOdin-l/TinderArt/pkg/db_models"
	"github.com/isOdin-l/TinderArt/services/profile/internal/entities"
)

// DB ---> ENTITY

func FromGetProfileToEntity(db *db_models.GetProfileRow) *entities.Profile {
	return &entities.Profile{
		UserId:      db_models.FromPgUUID(&db.ID),
		Username:    db_models.FromPgText(&db.Username),
		Surname:     db_models.FromPgText(&db.Surname),
		Name:        db_models.FromPgText(&db.Name),
		Email:       db_models.FromPgText(&db.Email),
		Description: db_models.FromPgText(&db.Description),
		Latitude:    db.Latitude.(float64),
		Longitude:   db.Longitude.(float64),
	}
}

func FromUpdateProfileToEntity(req *db_models.UpdateProfileRow) *entities.Profile {
	return &entities.Profile{
		UserId:      db_models.FromPgUUID(&req.ID),
		Username:    db_models.FromPgText(&req.Username),
		Surname:     db_models.FromPgText(&req.Surname),
		Name:        db_models.FromPgText(&req.Name),
		Email:       db_models.FromPgText(&req.Email),
		Description: db_models.FromPgText(&req.Description),
	}
}

// ENTITY ---> DB
func FromEntityToCreateProfile(entity *entities.Profile) *db_models.CreateProfileParams {
	return &db_models.CreateProfileParams{
		ID:            db_models.ToPgUUID(&entity.UserId),
		Username:      db_models.ToPgText(&entity.Username),
		Password:      db_models.ToPgText(&entity.Password),
		Surname:       db_models.ToPgText(&entity.Surname),
		Name:          db_models.ToPgText(&entity.Name),
		Email:         db_models.ToPgText(&entity.Email),
		Description:   db_models.ToPgText(&entity.Description),
		StMakepoint:   entity.Latitude,
		StMakepoint_2: entity.Longitude,
	}
}

func FromEntityToUpdateProfileParams(entity *entities.UpdateProfile) *db_models.UpdateProfileParams {
	return &db_models.UpdateProfileParams{
		ID:          db_models.ToPgUUID(&entity.UserId),
		Username:    db_models.ToPgText(entity.Username),
		Name:        db_models.ToPgText(entity.Name),
		Surname:     db_models.ToPgText(entity.Surname),
		Email:       db_models.ToPgText(entity.Email),
		Password:    db_models.ToPgText(entity.Password),
		Description: db_models.ToPgText(entity.Description),
		Longitude:   entity.Longitude,
		Latitude:    entity.Latitude,
	}
}
