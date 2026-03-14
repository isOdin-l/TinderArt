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
		return c.JSON(http.StatusBadRequest, "")
	}

	h.service.Registrations(c.Request().Context(), mappers.FromAPIRegistrationToRegistration(userApi))

	return nil
}

// Get new access_token and refresh_token, when both expired
func (h *AuthHandler) SignIn(c *echo.Context) error {
	userApi := new(api.Login)
	if errBind := c.Bind(&userApi); errBind != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	h.service.Login(c.Request().Context(), mappers.FromAPILoginToLogin(userApi))

	return nil
}

// Update access_token by refresh_token
func (h *AuthHandler) RefreshToken(c *echo.Context) error {
	tokenReq := new(api.RefreshAccessToken)
	if errBind := c.Bind(&tokenReq); errBind != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	h.service.RefreshAccessToken(c.Request().Context(), mappers.FromAPIRefreshTokenToRefreshToken(tokenReq))

	return nil
}

// Validate access_token
func (h *AuthHandler) ValidateToken(c *echo.Context) error {
	accessToken := new(api.ValidateToken)
	if errBind := c.Bind(&accessToken); errBind != nil {
		return c.JSON(http.StatusUnauthorized, "")
	}

	h.service.ValidateAccessToken(c.Request().Context(), mappers.FromAPIValidateTokenToValidateToken(accessToken))

	return nil
}
