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

// GRPC Client
type ConfigGrpcClient struct {
	GrpcServerPort string `env:"GRPC_SERVER_PORT"`
	GrpcServerHost string `env:"GRPC_SERVER_HOST"`
}

func (c *ConfigGrpcClient) DSN() string {
	return fmt.Sprintf("%s:%s", c.GrpcServerHost, c.GrpcServerPort)
}

// GRPC Server
type ConfigGrpcServer struct {
	GrpcServerPort string `env:"GRPC_SERVER_PORT"`
}

func (c *ConfigGrpcServer) DSN() string {
	return fmt.Sprintf(":%s", c.GrpcServerPort)
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

// RUSTFS
type ConfigRustFS struct {
	RustFSRegion            string `env:"RUSTFS_REGION"`
	RustFSAccess_key        string `env:"RUSTFS_ACCESS_KEY_ID"`
	RustFSSecret_access_key string `env:"RUSTFS_SECRET_ACCESS_KEY"`
	RustFSEndpoint          string `env:"RUSTFS_ENDPOINT_URL"`
}
