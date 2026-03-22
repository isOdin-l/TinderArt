package api

import (
	"github.com/isOdin-l/TinderArt/services/swipe/internal/entities"
	"github.com/isOdin-l/TinderArt/services/swipe/pkg/api"
)

func FromApiSwipeToEntity(req *api.CreateSwipeRequest) *entities.Swipe {
	return &entities.Swipe{
		TargetId:  req.TargetId,
		Decision1: &req.Decision,
		Decision2: nil,
	}
}

func ValidateSwipeStruct(ent *entities.Swipe) {
	if ent.UserId.String() <= ent.TargetId.String() {
		ent.UserId, ent.TargetId = ent.TargetId, ent.UserId
		ent.Decision1, ent.Decision2 = ent.Decision2, ent.Decision1
	}
}
