package server

import (
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

type IHandler interface {
	CreateUser(c *echo.Context) error    // Creates new access_token and refresh_token
	SignInHandler(c *echo.Context) error // Get new access_token and refresh_token, when both expired
	RefreshToken(c *echo.Context) error  // Update access_token by refresh_token
	ValidateToken(c *echo.Context) error // Validate access_token
}

func NewRouter(e *echo.Echo, h IHandler) {
	e.Use(middleware.RequestID())
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	auth := e.Group("/auth")

	auth.POST("/sign-in", h.SignInHandler)
	auth.POST("/refresh", h.RefreshToken)
	auth.POST("/register", h.CreateUser)
	auth.POST("/validate", h.ValidateToken)
}
