package configs

import "fmt"

type ConfigPostgreWithPostGIS struct {
	DbPassword string `env:"DB_PASSWORD"`
	DbUserName string `env:"DB_USERNAME"`
	DbHost     string `env:"DB_HOST"`
	DbPort     string `env:"DB_PORT"`
	DbName     string `env:"DB_NAME"`
}

func (c *ConfigPostgreWithPostGIS) DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.DbUserName, c.DbPassword, c.DbHost, c.DbPort, c.DbName)
}

type ConfigGrpcAuth struct {
	GrpcServerPort string `env:"GRPC_SERVER_PORT"`
	GrpcServerHost string `env:"GRPC_SERVER_HOST"`
}

func (c *ConfigGrpcAuth) DSN() string {
	return fmt.Sprintf("%s:%s", c.GrpcServerHost, c.GrpcServerPort)
}
