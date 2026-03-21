package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	ticker := time.NewTicker(time.Duration(cfg.TickPeriodHours) * time.Hour)

	go runDailyStackGeneration(ctx, ticker, service)

	<-ctx.Done()
}

func runDailyStackGeneration(ctx context.Context, ticker *time.Ticker, service *service.Service) error {
	// Infinity loop with block by ticker chanel, which opens every ticke
	// and we generate new stacks for users
	for {
		<-ticker.C
		if errServ := service.GenerateDailyStack(ctx); errServ != nil {
			slog.Error(fmt.Sprintf("Internal server error: %s", errServ.Error()))
		}
	}
}
