package configs

import "fmt"

// POSTGRES
type ConfigPostgres struct {
	DbPassword string `env:"DB_PASSWORD"`
	DbUserName string `env:"DB_USERNAME"`
	DbHost     string `env:"DB_HOST"`
	DbPort     string `env:"DB_PORT"`
	DbName     string `env:"DB_NAME"`
}

func (c *ConfigPostgres) DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.DbUserName, c.DbPassword, c.DbHost, c.DbPort, c.DbName)
}

// GRPC
type ConfigGrpcAuth struct {
	GrpcServerPort string `env:"GRPC_SERVER_PORT"`
	GrpcServerHost string `env:"GRPC_SERVER_HOST"`
}

func (c *ConfigGrpcAuth) DSN() string {
	return fmt.Sprintf("%s:%s", c.GrpcServerHost, c.GrpcServerPort)
}

// REDIS
type ConfigRedis struct {
	RedisHost     string `env:"REDIS_HOST"`
	RedisPort     string `env:"REDIS_PORT"`
	RedisPassword string `env:"REDIS_PASSWORD"`
}

func (c *ConfigRedis) DSN() string {
	return fmt.Sprintf("%s:%s", c.RedisHost, c.RedisPort)
}

// KAFKA
type ConfigKafka struct {
	KafkaHost      string `env:"KAFKA_HOST"`
	KafkaPort      string `env:"KAFKA_PORT"`
	KafkaTopic     string `env:"KAFKA_TOPIC"`
	KafkaPartition int    `env:"KAFKA_PARTITION"`
}

func (c *ConfigKafka) DSN() string {
	return fmt.Sprintf("%s:%s", c.KafkaHost, c.KafkaPort)
}
