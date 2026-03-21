package redis

import (
	"context"
	"time"

	"github.com/isOdin-l/TinderArt/pkg/configs"
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Client *redis.Client
}

func NewRedis(cfg *configs.ConfigRedis) *Redis {
	return &Redis{
		Client: redis.NewClient(&redis.Options{
			Addr:     cfg.DSN(),
			Password: cfg.RedisPassword,
			// DB:       0,
		}),
	}
}

func (cache *Redis) Set(ctx context.Context, key string, value any, timeExpire time.Duration) error {
	return cache.Client.Set(ctx, key, value, timeExpire).Err()
}

func (cache *Redis) Get(ctx context.Context, key string) (string, error) {
	return cache.Client.Get(ctx, key).Result()
}

func (cache *Redis) LRange(ctx context.Context, key string, start, end int64) ([]any, error) {
	res := cache.Client.LRange(ctx, key, start, end)
	if res.Err() != nil {
		return nil, res.Err()
	}

	return res.Args(), nil
}

func (cache *Redis) RPush(ctx context.Context, key string, args ...any) error {
	return cache.Client.RPush(ctx, key, args...).Err()
}

func (cache *Redis) Eval(ctx context.Context, script string, keys []string, args ...any) error {
	return cache.Client.Eval(ctx, script, keys, args...).Err()
}

func (cache *Redis) Del(ctx context.Context, keys ...string) error {
	return cache.Client.Del(ctx, keys...).Err()
}

func (cache *Redis) LPopCount(ctx context.Context, key string, count int) ([]string, error) {
	return cache.Client.LPopCount(ctx, key, count).Result()
}
