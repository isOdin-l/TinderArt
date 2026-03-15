package config

import (
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/isOdin-l/TinderArt/pkg/configs"
)

type Config struct {
	configs.ConfigPostgreWithPostGIS
	configs.ConfigGrpcAuth
	ServerConfig
	InternalConfig
}

type ServerConfig struct {
	HttpServerPort string `env:"SERVER_PORT"`
}

type InternalConfig struct {
	HashMinCost     int           `env:"HASH_MIN_COST"`
	AccessSignKey   string        `env:"ACCESS_SIGNING_KEY"`
	RefreshSignKey  string        `env:"REFRESH_SIGNING_KEY"`
	AccessTokenTTL  time.Duration `env:"ACCESS_TOKEN_TTL"`
	RefreshTokenTTL time.Duration `env:"REFRESH_TOKEN_TTL"`
}

func NewConfig() (Config, error) {
	return env.ParseAs[Config]()
}
