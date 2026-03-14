package config

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v11"
	echojwt "github.com/labstack/echo-jwt/v5"
)

type Config struct {
	ServerPort string `env:"SERVER_PORT"`
	DbPassword string `env:"DB_PASSWORD"`
	DbUserName string `env:"DB_USERNAME"`
	DbHost     string `env:"DB_HOST"`
	DbPort     string `env:"DB_PORT"`
	DbName     string `env:"DB_NAME"`
	InternalConfig
}

type InternalConfig struct {
	HashMinCost int           `env:"COST"`
	JwtSignKey  string        `env:"JWT_SIGNING_KEY"`
	TokenTTL    time.Duration `env:"TOKEN_TTL"`
	JwtConfig   echojwt.Config
}

func NewConfig() (Config, error) {
	return env.ParseAs[Config]()
}

func (c *Config) DSNPsql() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.DbUserName, c.DbPassword, c.DbHost, c.DbPort, c.DbName)
}
