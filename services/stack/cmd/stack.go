package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/isOdin-l/TinderArt/pkg/postgres"
	"github.com/isOdin-l/TinderArt/pkg/redis"
	"github.com/isOdin-l/TinderArt/services/stack/config"
	"github.com/isOdin-l/TinderArt/services/stack/internal/repository"
	"github.com/isOdin-l/TinderArt/services/stack/internal/service"
)

func main() {
	// Config initilizing
	cfg, errCfg := config.NewConfig()
	if errCfg != nil {
		slog.Error(fmt.Sprintf("Error while initilizing config: %s", errCfg.Error()))
		return
	}

	// Redis connections
	redis := redis.NewRedis(&cfg.ConfigRedis)
	defer redis.Client.Close()

	// Postgresql connection
	DB, errDb := postgres.NewPostgresDB(&cfg.ConfigPostgres)
	if errDb != nil {
		slog.Error(fmt.Sprintf("Error while initilizing database: %s", errDb.Error()))
		return
	}

	repository := repository.NewRepository(DB)
	service := service.NewService(redis, repository)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	// Run stack generation
	service.GenerateDailyStack(ctx)
}
