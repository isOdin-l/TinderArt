package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	grpc_auth "github.com/isOdin-l/TinderArt/pkg/grpc"
	"github.com/isOdin-l/TinderArt/pkg/middleware"
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

	// Config
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

	// Grpc Client
	grpc_client, errClient := grpc_auth.NewGrpcAuthClient(&cfg.ConfigGrpcClient)
	if errClient != nil {
		router.Logger.Error(fmt.Sprintf("Error while initilizing connection with grpc server: %s", errClient.Error()))
		return
	}
	defer grpc_client.Conn.Close()

	// Layers
	repository := repository.NewRepository(DB)                                           // Repository
	service := service.NewService(repository, storage, grpc_client, &cfg.InternalConfig) // Service
	handler := handler.NewHandler(service)                                               // Handler

	// Custom middleware with grpc call
	md := middleware.NewMiddleware(grpc_client)

	server.NewRoutes(router, handler, md) // Routing

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	if errServer := server.RunServer(ctx, cfg.ServerConfig, router); errServer != nil {
		router.Logger.Error(fmt.Sprintf("Server error: %s", errServer.Error()))
		return
	}
}
