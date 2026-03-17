package api

import (
	"github.com/google/uuid"
	"github.com/isOdin-l/TinderArt/services/profile/internal/entities"
	"github.com/isOdin-l/TinderArt/services/profile/pkg/api"
)

func FromAPIGetProfileToEntity(req *api.RequestGetProfile) *entities.Profile {
	return &entities.Profile{
		UserId: req.UserId,
	}
}

func FromAPIUpdateProfileToEntity(req api.RequestUpdateProfile, userId uuid.UUID) *entities.UpdateProfile {
	return &entities.UpdateProfile{
		UserId:      userId,
		Username:    req.Username,
		Name:        req.Name,
		Surname:     req.Surname,
		Email:       req.Email,
		Password:    req.Password,
		Description: req.Description,
		Latitude:    req.Latitude,
		Longitude:   req.Longitude,
	}
}

// all entity fields must be not nil
func FromEntityToAPIUpdateProfile(entity *entities.UpdateProfile) *api.ResponseProfile {
	return &api.ResponseProfile{
		Username:    *entity.Username,
		Name:        *entity.Name,
		Surname:     *entity.Surname,
		Email:       *entity.Email,
		Password:    *entity.Password,
		Description: *entity.Description,
		Latitude:    *entity.Latitude,
		Longitude:   *entity.Longitude,
	}
}

func FromEntityToApiGetProfile(entity *entities.Profile) *api.ResponseProfile {
	return &api.ResponseProfile{
		Username:    entity.Username,
		Name:        entity.Name,
		Surname:     entity.Surname,
		Email:       entity.Email,
		Password:    entity.Password,
		Description: entity.Description,
		Latitude:    entity.Latitude,
		Longitude:   entity.Longitude,
	}
}
