package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/isOdin-l/TinderArt/pkg/configs"
)

type Config struct {
	configs.ConfigRedis
	configs.ConfigPostgres
	InternalConfig
}

type InternalConfig struct {
	// That variable used to set ticker, which rule after what period of hours
	// stack generation should be
	TickPeriodHours int64 `env:"TICK_PERIOD_HOURS"`
}

func NewConfig() (Config, error) {
	return env.ParseAs[Config]()
}
