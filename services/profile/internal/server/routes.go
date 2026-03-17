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

type IMiddleware interface {
	Validation() echo.MiddlewareFunc
}

func NewRoutes(router *echo.Echo, h IHandler, md IMiddleware) {
	router.Use(middleware.RequestID())
	router.Use(middleware.RequestLogger())
	router.Use(middleware.Recover())
	router.Use(md.Validation()) // Validate token in Authorization header and put it into context

	router.GET("/profile", h.GetProfile)
	router.PATCH("/profile", h.UpdateProfile)
	router.DELETE("/profile", h.DeleteProfile)
}
