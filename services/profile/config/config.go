package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/isOdin-l/TinderArt/pkg/configs"
)

type Config struct {
	configs.ConfigRustFS
	configs.ConfigPostgres
	configs.ConfigGrpcClient
	configs.ConfigRedis
	InternalConfig
	ServerConfig
}

type ServerConfig struct {
	HttpServerPort string `env:"HTTP_SERVER_PORT"`
}

type InternalConfig struct {
	HashMinCost int `env:"HASH_MIN_COST"`
}

func NewConfig() (Config, error) {
	return env.ParseAs[Config]()
}
