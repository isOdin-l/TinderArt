package server

import (
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

type IHandler interface {
	CreateProfile(c *echo.Context) error
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

	router.POST("/create", h.CreateProfile)

	protected := router.Group("/protected")
	protected.Use(md.Validation()) // Validate token in Authorization header and put it into context

	protected.GET("/get", h.GetProfile)
	protected.PATCH("/update", h.UpdateProfile)
	protected.DELETE("/delete", h.DeleteProfile)

	// router.PATCH("/update", h.UpdatePreferences)
	// router.DELETE("/delete", h.DeletePhotos)
	// etc
}
