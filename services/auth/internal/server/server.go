package server

import (
	"context"
	"time"

	config "github.com/isOdin-l/TinderArt/services/auth/configs"
	"github.com/labstack/echo/v5"
)

// HTTP server to handler REST requests
func RunServer(router *echo.Echo, ctx *context.Context, cfg *config.ServerConfig) error {
	server := echo.StartConfig{
		Address:         cfg.HttpServerPort,
		GracefulTimeout: 5 * time.Second,
	}

	return server.Start(*ctx, router)
}
