package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	grpcClient "github.com/isOdin-l/TinderArt/pkg/grpc/auth"
	"github.com/isOdin-l/TinderArt/pkg/postgres"
	config "github.com/isOdin-l/TinderArt/services/auth/configs"
	"github.com/isOdin-l/TinderArt/services/auth/internal/handler"
	"github.com/isOdin-l/TinderArt/services/auth/internal/repository"
	"github.com/isOdin-l/TinderArt/services/auth/internal/server"
	"github.com/isOdin-l/TinderArt/services/auth/internal/service"
	"github.com/labstack/echo/v5"
)

func main() {
	router := echo.New()

	// Config
	cfg, errCfg := config.NewConfig()
	if errCfg != nil {
		router.Logger.Error(fmt.Sprintf("Error while initialize config: %s", errCfg.Error()))
		return
	}

	// Database
	DB, errDb := postgres.NewPostgresDB(&cfg.ConfigPostgreWithPostGIS)
	if errDb != nil {
		router.Logger.Error(fmt.Sprintf("failed to initialize db: %s", errDb.Error()))
		return
	}
	defer DB.Close()

	// grpc Client
	grpcClient, errGrpc := grpcClient.NewGrpcClient(&cfg.ConfigGrpcAuth)
	if errGrpc != nil {
		router.Logger.Error(fmt.Sprintf("failed to connect to grpc server: %s", errGrpc.Error()))
		return
	}
	defer grpcClient.Conn.Close()

	repository := repository.NewRepository(DB)                                     //  Repository
	service := service.NewService(&cfg.InternalConfig, repository, DB, grpcClient) //  Service
	handler := handler.NewHandler(service)                                         //  Handler

	// Definition context for server's graceful shutdown
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	server.NewRouter(router, &cfg, handler) //  Routing
	if err := server.RunServer(router, &ctx, &cfg.ServerConfig); err != nil {
		router.Logger.Error(fmt.Sprintf("Error while running server %s", err.Error()))
	}
}
