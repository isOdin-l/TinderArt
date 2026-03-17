package api

import (
	"github.com/google/uuid"
	"github.com/isOdin-l/TinderArt/services/profile/internal/entities"
	"github.com/isOdin-l/TinderArt/services/profile/pkg/api"
)

func FromAPIGetProfileToEntity(req *api.RequestGetProfile) uuid.UUID {
	return req.UserId
}

func FromAPIUpdateProfileToEntity(req *api.RequestUpdateProfile) *entities.UpdateProfile {
	return &entities.UpdateProfile{
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

func FromEntityToAPIGetProfile(entity *entities.Profile) *api.ResponseProfile {
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
