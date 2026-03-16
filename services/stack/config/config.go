package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/isOdin-l/TinderArt/pkg/configs"
)

type Config struct {
	configs.ConfigRedis
	configs.ConfigPostgreWithPostGIS
}

func NewConfig() (Config, error) {
	return env.ParseAs[Config]()
}
