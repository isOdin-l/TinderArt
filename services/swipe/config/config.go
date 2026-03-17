package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/isOdin-l/TinderArt/pkg/configs"
)

type ConfigServer struct {
	HttpServerPort string `env:"HTTP_SERVER_PORT"`
}

type Config struct {
	ConfigServer
	configs.ConfigPostgres
	configs.ConfigKafka
	configs.ConfigGrpcClient
}

func NewConfig() (Config, error) {
	return env.ParseAs[Config]()
}
