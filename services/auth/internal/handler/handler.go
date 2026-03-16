package handler

import (
	"context"
	"net/http"

	mappers "github.com/isOdin-l/TinderArt/services/auth/internal/api"
	"github.com/isOdin-l/TinderArt/services/auth/internal/entities"
	"github.com/isOdin-l/TinderArt/services/auth/pkg/api"
	"github.com/labstack/echo/v5"
)

type IService interface {
	Registrations(ctx context.Context, entity *entities.Registration) error
	Login(ctx context.Context, entity *entities.Login) error
	RefreshAccessToken(ctx context.Context, entity *entities.RefreshAccessToken) error
	ValidateAccessToken(ctx context.Context, entity *entities.ValidateToken) error
}

type AuthHandler struct {
	service IService
}

func NewHandler(service IService) *AuthHandler {
	return &AuthHandler{service: service}
}

// Creates new access_token and refresh_token
func (h *AuthHandler) Registrations(c *echo.Context) error {
	userApi := new(api.Registration)
	if errBind := c.Bind(&userApi); errBind != nil {
		return c.JSON(http.StatusBadRequest, api.ErrorResponse{Error: "invalid request"})
	}

	entity := mappers.FromAPIRegistrationToRegistration(userApi)
	if err := h.service.Registrations(c.Request().Context(), entity); err != nil {
		return c.JSON(http.StatusBadRequest, api.ErrorResponse{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, mappers.FromRegistrationToTokenResponse(entity))
}

// Get new access_token and refresh_token, when both expired
func (h *AuthHandler) SignIn(c *echo.Context) error {
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
func (h *AuthHandler) RefreshToken(c *echo.Context) error {
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

// Validate access_token
func (h *AuthHandler) ValidateToken(c *echo.Context) error {
	accessToken := new(api.ValidateToken)
	if errBind := c.Bind(&accessToken); errBind != nil {
		return c.JSON(http.StatusUnauthorized, api.ErrorResponse{Error: "missing token"})
	}

	entity := mappers.FromAPIValidateTokenToValidateToken(accessToken)
	if err := h.service.ValidateAccessToken(c.Request().Context(), entity); err != nil {
		return c.JSON(http.StatusUnauthorized, api.ErrorResponse{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, api.ValidateResponse{Valid: true})
}
