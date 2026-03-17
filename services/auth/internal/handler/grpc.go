package handler

import (
	"context"

	"github.com/google/uuid"
	grpc_auth "github.com/isOdin-l/TinderArt/pkg/grpc/auth"
	mapper "github.com/isOdin-l/TinderArt/services/auth/internal/api"
	"github.com/isOdin-l/TinderArt/services/auth/internal/entities"
)

type IServiceGrpc interface {
	Registrations(ctx context.Context, entity *entities.Registration) error
	ValidateAccessToken(ctx context.Context, entity *entities.ValidateToken) (uuid.UUID, error)
}

type HandlerGrpc struct {
	grpc_auth.UnimplementedAuthServiceServer
	service IServiceGrpc
}

func NewHandlerGrpc(s IServiceGrpc) *HandlerGrpc {
	return &HandlerGrpc{service: s}
}

func (h *HandlerGrpc) CreateUser(ctx context.Context, req *grpc_auth.CreateUserRequest) (*grpc_auth.CreateUserResponse, error) {
	entity := mapper.FromAPIRegistrationToRegistration(req)
	err := h.service.Registrations(ctx, entity)
	if err != nil {
		return nil, err
	}

	return mapper.FromRegistrationToTokenResponse(entity), nil
}

func (h *HandlerGrpc) Validate(ctx context.Context, req *grpc_auth.ValidateRequest) (*grpc_auth.ValidateResponse, error) {
	entity := mapper.FromAPIValidateTokenToValidateToken(req)
	userId, errVal := h.service.ValidateAccessToken(ctx, entity)
	if errVal != nil {
		return nil, errVal
	}

	return &grpc_auth.ValidateResponse{
		UserId: userId.String(),
	}, nil
}
