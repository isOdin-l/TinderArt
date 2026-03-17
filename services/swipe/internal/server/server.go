package server

import (
	"context"
	"time"

	"github.com/isOdin-l/TinderArt/services/swipe/config"
	"github.com/labstack/echo/v5"
)

func RunServer(ctx context.Context, router *echo.Echo, cfg *config.ConfigServer) error {
	server := echo.StartConfig{
		Address:         cfg.HttpServerPort,
		GracefulTimeout: 5 * time.Second,
	}

	return server.Start(ctx, router)
}
