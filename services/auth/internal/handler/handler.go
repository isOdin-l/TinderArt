package handler

import "github.com/labstack/echo/v5"

type IService interface {
}

type AuthHandler struct {
	s IService
}

func NewHandler(service IService) *AuthHandler {
	return &AuthHandler{s: service}
}

// Creates new access_token and refresh_token
func (h *AuthHandler) CreateUser(c *echo.Context) error {
	return nil
}

// Get new access_token and refresh_token, when both expired
func (h *AuthHandler) SignInHandler(c *echo.Context) error {
	return nil
}

// Update access_token by refresh_token
func (h *AuthHandler) RefreshToken(c *echo.Context) error {
	return nil
}

// Validate access_token
func (h *AuthHandler) ValidateToken(c *echo.Context) error {
	return nil
}
