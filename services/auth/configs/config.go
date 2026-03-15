package config

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	DatabaseConfig
	ServerConfig
	InternalConfig
}

type DatabaseConfig struct {
	DbPassword string `env:"DB_PASSWORD"`
	DbUserName string `env:"DB_USERNAME"`
	DbHost     string `env:"DB_HOST"`
	DbPort     string `env:"DB_PORT"`
	DbName     string `env:"DB_NAME"`
}

type ServerConfig struct {
	HttpServerPort string `env:"SERVER_PORT"`
	GrpcServerPort string `env:"GRPC_SERVER_PORT"`
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

func (c *Config) DSNPsql() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.DbUserName, c.DbPassword, c.DbHost, c.DbPort, c.DbName)
}

func (c *ServerConfig) DSNgrpc() string {
	return fmt.Sprintf("localhost:%s", c.GrpcServerPort)
}
