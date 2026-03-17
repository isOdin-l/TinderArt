package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/isOdin-l/TinderArt/pkg/kafka"
	"github.com/isOdin-l/TinderArt/pkg/postgres"
	"github.com/isOdin-l/TinderArt/services/swipe/config"
	"github.com/isOdin-l/TinderArt/services/swipe/internal/handler"
	"github.com/isOdin-l/TinderArt/services/swipe/internal/repository"
	"github.com/isOdin-l/TinderArt/services/swipe/internal/server"
	"github.com/isOdin-l/TinderArt/services/swipe/internal/service"
	"github.com/labstack/echo/v5"
)

func main() {
	router := echo.New()

	cfg, errCfg := config.NewConfig()
	if errCfg != nil {
		router.Logger.Error(fmt.Sprintf("Error while initializing config: %s", errCfg.Error()))
		return
	}

	DB, errDb := postgres.NewPostgresDB(&cfg.ConfigPostgres)
	if errDb != nil {
		router.Logger.Error(fmt.Sprintf("Error while initializing database: %s", errDb.Error()))
		return
	}
	defer DB.Close()

	kafka, errMsg := kafka.NewKafka(&cfg.ConfigKafka)
	if errMsg != nil {
		router.Logger.Error(fmt.Sprintf("Error while initializing kafka: %s", errMsg.Error()))
		return
	}

	repository := repository.NewRepository(DB)
	service := service.NewService(repository, kafka, DB)
	handler := handler.NewHandler(service)

	server.CreateRoutes(router, handler)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	if err := server.RunServer(ctx, router, &cfg.ConfigServer); err != nil {
		router.Logger.Error(fmt.Sprintf("Error while running server %s", err.Error()))
	}
}
