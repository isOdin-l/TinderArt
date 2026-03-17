package main

import (
	"fmt"

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
	DB, errDb := postgres.NewPostgresDB(&cfg.ConfigPostgres)
	if errDb != nil {
		router.Logger.Error(fmt.Sprintf("failed to initialize db: %s", errDb.Error()))
		return
	}
	defer DB.Close()

	repository := repository.NewRepository(DB)                         //  Repository
	service := service.NewService(&cfg.InternalConfig, repository, DB) //  Service
	handler := handler.NewHandler(service)                             //  Handler

	server.NewRouter(router, &cfg, handler)                                                              //  Routing
	serverRunner, errServ := server.NewServer(&cfg.ServerConfig, &cfg.ConfigGrpcServer, router, handler) // Server initialization
	if errServ != nil {
		router.Logger.Error(fmt.Sprintf("Error whili initilizing server runner: %s", errServ.Error()))
		return
	}

	serverRunner.RunServer()
}
