package server

import (
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

type IHandler interface {
	CreateSwipe(c *echo.Context) error
}

func CreateRoutes(router *echo.Echo, h IHandler) *echo.Echo {
	router.Use(middleware.RequestID())
	router.Use(middleware.RequestLogger())
	router.Use(middleware.Recover())

	router.POST("/swipe", h.CreateSwipe)

	return router
}
