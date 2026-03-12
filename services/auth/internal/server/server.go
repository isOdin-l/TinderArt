package server

import (
	"context"
	"time"

	"github.com/labstack/echo/v5"
)

func RunServer(router *echo.Echo, ctx *context.Context, port string) error {
	server := echo.StartConfig{
		Address:         port,
		GracefulTimeout: 5 * time.Second,
	}

	return server.Start(*ctx, router)
}
