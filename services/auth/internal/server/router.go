package server

import (
	config "github.com/isOdin-l/TinderArt/services/auth/configs"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

type IHandlerRest interface {
	SignIn(c *echo.Context) error       // Get new access_token and refresh_token, when both expired
	RefreshToken(c *echo.Context) error // Update access_token by refresh_token
}

func NewRouter(e *echo.Echo, cfg *config.Config, h IHandlerRest) {
	e.Use(middleware.RequestID())
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	auth := e.Group("/auth")

	auth.POST("/sign-in", h.SignIn)
	auth.POST("/refresh", h.RefreshToken)
}
