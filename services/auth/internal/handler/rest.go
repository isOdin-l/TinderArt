package handler

import (
	"context"
	"net/http"

	mappers "github.com/isOdin-l/TinderArt/services/auth/internal/api"
	"github.com/isOdin-l/TinderArt/services/auth/internal/entities"
	"github.com/isOdin-l/TinderArt/services/auth/pkg/api"
	"github.com/labstack/echo/v5"
)

type IServiceRest interface {
	Login(ctx context.Context, entity *entities.Login) error
	RefreshAccessToken(ctx context.Context, entity *entities.RefreshAccessToken) error
}

type HandlerRest struct {
	service IServiceRest
}

func NewHandlerRest(service IServiceRest) *HandlerRest {
	return &HandlerRest{service: service}
}

// Get new access_token and refresh_token, when both expired
func (h *HandlerRest) SignIn(c *echo.Context) error {
	userApi := new(api.Login)
	if errBind := c.Bind(&userApi); errBind != nil {
		return c.JSON(http.StatusBadRequest, api.ErrorResponse{Error: "invalid request"})
	}

	entity := mappers.FromAPILoginToLogin(userApi)
	if err := h.service.Login(c.Request().Context(), entity); err != nil {
		return c.JSON(http.StatusUnauthorized, api.ErrorResponse{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, mappers.FromLoginToTokenResponse(entity))
}

// Update access_token by refresh_token
func (h *HandlerRest) RefreshToken(c *echo.Context) error {
	tokenReq := new(api.RefreshAccessToken)
	if errBind := c.Bind(&tokenReq); errBind != nil {
		return c.JSON(http.StatusBadRequest, api.ErrorResponse{Error: "invalid request"})
	}

	entity := mappers.FromAPIRefreshTokenToRefreshToken(tokenReq)
	if err := h.service.RefreshAccessToken(c.Request().Context(), entity); err != nil {
		return c.JSON(http.StatusUnauthorized, api.ErrorResponse{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, mappers.FromRefreshAccessTokenToTokenResponse(entity))
}
