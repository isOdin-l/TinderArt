package handler

import (
	"context"
	"fmt"
	"log/slog"

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

func (h *HandlerGrpc) SignTokens(ctx context.Context, req *grpc_auth.RequestSignTokens) (*grpc_auth.ResponseSignTokens, error) {
	entity, errMapper := mapper.FromAPIRegistrationToRegistration(req)
	if errMapper != nil {
		slog.Error(fmt.Sprintf("error: %s request data", errMapper.Error()))
		return nil, errMapper
	}
	err := h.service.Registrations(ctx, entity)

	if err != nil {
		slog.Error(fmt.Sprintf("error: %s request data", err.Error()))
		return nil, err
	}

	return mapper.FromRegistrationToTokenResponse(entity), nil
}

func (h *HandlerGrpc) Validate(ctx context.Context, req *grpc_auth.ValidateRequest) (*grpc_auth.ValidateResponse, error) {
	entity := mapper.FromAPIValidateTokenToValidateToken(req)
	userId, errVal := h.service.ValidateAccessToken(ctx, entity)

	if errVal != nil {
		slog.Error(fmt.Sprintf("error: %s data:%s", errVal.Error(), req))
		return nil, errVal
	}

	return &grpc_auth.ValidateResponse{
		UserId: userId.String(),
	}, nil
}
