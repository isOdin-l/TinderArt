package api

import (
	"github.com/isOdin-l/TinderArt/services/swipe/internal/entities"
	"github.com/isOdin-l/TinderArt/services/swipe/pkg/api"
)

func FromApiSwipeToEntity(req *api.CreateSwipeRequest) *entities.Swipe {
	if req.UserId.String() > req.TargetId.String() {
		return &entities.Swipe{
			UserId:    req.UserId,
			TargetId:  req.TargetId,
			Decision1: &req.Decision,
			Decision2: nil,
		}
	}
	return &entities.Swipe{
		UserId:    req.TargetId,
		TargetId:  req.UserId,
		Decision1: nil,
		Decision2: &req.Decision,
	}
}
