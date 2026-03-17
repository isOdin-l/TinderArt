package api

import "github.com/google/uuid"

type CreateSwipeRequest struct {
	UserId   uuid.UUID `json:"user_id"`
	TargetId uuid.UUID `json:"target_id"`
	Decision bool      `json:"decision"`
}
