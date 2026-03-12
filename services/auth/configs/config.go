package config

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v11"
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
	Salt       string        `env:"SALT"`
	JwtSignKey string        `env:"JWT_SIGNING_KEY"`
	TokenTTL   time.Duration `env:"TOKEN_TTL"`
}

func NewConfig() (Config, error) {
	return env.ParseAs[Config]()
}

func (c *Config) DSNPsql() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.DbUserName, c.DbPassword, c.DbHost, c.DbPort, c.DbName)
}
