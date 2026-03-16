package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/isOdin-l/TinderArt/pkg/configs"
)

type Config struct {
	configs.ConfigPostgres
	configs.ConfigKafka
}

func NewConfig() (Config, error) {
	return env.ParseAs[Config]()
}
