package server

import (
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

type IHandler interface {
	GetProfile(c *echo.Context) error
	UpdateProfile(c *echo.Context) error
	DeleteProfile(c *echo.Context) error
}

func NewRoutes(router *echo.Echo, h IHandler) {
	router.Use(middleware.RequestID())
	router.Use(middleware.RequestLogger())
	router.Use(middleware.Recover())

	router.GET("/profile", h.GetProfile)
	router.PATCH("/profile", h.UpdateProfile)
	router.DELETE("/profile", h.DeleteProfile)
}
