package server

import (
	"context"
	"time"

	"github.com/isOdin-l/TinderArt/services/profile/config"
	"github.com/labstack/echo/v5"
)

func RunServer(ctx context.Context, cfg config.ServerConfig, router *echo.Echo) error {
	server := echo.StartConfig{
		Address:         cfg.HttpServerPort,
		GracefulTimeout: 5 * time.Second,
	}

	return server.Start(ctx, router)
}
