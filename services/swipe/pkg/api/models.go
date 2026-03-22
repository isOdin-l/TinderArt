package api

import "github.com/google/uuid"

type CreateSwipeRequest struct {
	TargetId uuid.UUID `json:"target_id"`
	Decision bool      `json:"decision"`
}
