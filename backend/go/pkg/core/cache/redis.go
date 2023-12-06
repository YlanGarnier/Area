package cache

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Cache interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value string) error
}

type RedisCache struct {
	redisClient *redis.Client
}

func NewRedisCache() Cache {
	// TODO: env addr
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0, //default
	})
	return &RedisCache{
		redisClient: redisClient,
	}
}

func (c *RedisCache) Get(ctx context.Context, key string) (string, error) {
	return c.redisClient.Get(ctx, key).Result()
}

func (c *RedisCache) Set(ctx context.Context, key string, value string) error {
	return c.redisClient.Set(ctx, key, value, 0).Err()
}
