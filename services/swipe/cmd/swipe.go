package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	grpc_auth "github.com/isOdin-l/TinderArt/pkg/grpc"
	"github.com/isOdin-l/TinderArt/pkg/kafka"
	"github.com/isOdin-l/TinderArt/pkg/middleware"
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

	// Config
	cfg, errCfg := config.NewConfig()
	if errCfg != nil {
		router.Logger.Error(fmt.Sprintf("Error while initializing config: %s", errCfg.Error()))
		return
	}

	// Database
	DB, errDb := postgres.NewPostgresDB(&cfg.ConfigPostgres)
	if errDb != nil {
		router.Logger.Error(fmt.Sprintf("Error while initializing database: %s", errDb.Error()))
		return
	}
	defer DB.Close()

	// Message Broker
	kafka, errMsg := kafka.NewKafka(&cfg.ConfigKafka)
	if errMsg != nil {
		router.Logger.Error(fmt.Sprintf("Error while initializing kafka: %s", errMsg.Error()))
		return
	}

	// Grpc Client
	grpc_client, errClient := grpc_auth.NewGrpcAuthClient(&cfg.ConfigGrpcClient)
	if errClient != nil {
		router.Logger.Error(fmt.Sprintf("Error while initilizing connection with grpc server: %s", errClient.Error()))
		return
	}
	defer grpc_client.Conn.Close()

	// App Layers
	repository := repository.NewRepository(DB)           // Repository
	service := service.NewService(repository, kafka, DB) // Service
	handler := handler.NewHandler(service)               // Handler

	// Custom middleware with grpc call
	md := middleware.NewMiddleware(grpc_client)

	// Routing
	server.CreateRoutes(router, handler, md)

	// context for graceful shutdown
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	if err := server.RunServer(ctx, router, &cfg.ConfigServer); err != nil {
		router.Logger.Error(fmt.Sprintf("Error while running server %s", err.Error()))
	}
}
