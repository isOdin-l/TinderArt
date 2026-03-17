package api

import (
	"github.com/isOdin-l/TinderArt/services/swipe/internal/entities"
	"github.com/isOdin-l/TinderArt/services/swipe/pkg/api"
)

func FromApiSwipeToEntity(req *api.CreateSwipeRequest) *entities.Swipe {
	return &entities.Swipe{
		UserId:    req.TargetId,
		TargetId:  req.UserId,
		Decision1: req.Decision,
	}
}
