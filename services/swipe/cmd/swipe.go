package main

import (
	"fmt"
	"log/slog"

	"github.com/isOdin-l/TinderArt/pkg/postgres"
	"github.com/isOdin-l/TinderArt/services/swipe/config"
)

func main() {
	cfg, errCfg := config.NewConfig()
	if errCfg != nil {
		slog.Error(fmt.Sprintf("Error while initilizing config: %s", errCfg.Error()))
		return
	}

	DB, errDb := postgres.NewPostgresDB(&cfg.ConfigPostgres)
	if errDb != nil {
		slog.Error(fmt.Sprintf("Error while initilizing database: %s", errDb.Error()))
		return
	}
	defer DB.Close()
}
