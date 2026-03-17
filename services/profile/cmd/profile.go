package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/isOdin-l/TinderArt/pkg/postgres"
	"github.com/isOdin-l/TinderArt/pkg/s3"
	"github.com/isOdin-l/TinderArt/services/profile/config"
	"github.com/isOdin-l/TinderArt/services/profile/internal/handler"
	"github.com/isOdin-l/TinderArt/services/profile/internal/repository"
	"github.com/isOdin-l/TinderArt/services/profile/internal/server"
	"github.com/isOdin-l/TinderArt/services/profile/internal/service"
	"github.com/labstack/echo/v5"
)

func main() {
	router := echo.New()

	// Конфиг
	cfg, errCfg := config.NewConfig()
	if errCfg != nil {
		router.Logger.Error(fmt.Sprintf("Error while initilizing config: %s", errCfg.Error()))
		return
	}

	// Database
	DB, errDb := postgres.NewPostgresDB(&cfg.ConfigPostgres)
	if errDb != nil {
		router.Logger.Error(fmt.Sprintf("Error while initilizing database: %s", errDb.Error()))
		return
	}
	defer DB.Close()

	// S3 storage
	storage := s3.NewRustFS(&cfg.ConfigRustFS)

	repository := repository.NewRepository(DB)
	service := service.NewService(repository, storage)
	handler := handler.NewHandler(service)

	server.NewRoutes(router, handler)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	if errServer := server.RunServer(ctx, cfg.ServerConfig, router); errServer != nil {
		router.Logger.Error(fmt.Sprintf("Server error: %s", errServer.Error()))
		return
	}
}
