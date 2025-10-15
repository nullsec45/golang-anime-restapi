package cache

import (
	"time"
	"context"
	goredis "github.com/redis/go-redis/v9"
)

type Cache interface {
	Get(key string) (string, error)
	Set(key, value string, ttl time.Duration) error
	Del(key string) error
}

type RedisCache struct {
	rdb *goredis.Client
}

func NewRedisCache(rdb *goredis.Client) *RedisCache{
	return &RedisCache{rdb:rdb}
}

func (c *RedisCache) Get(key string) (string, error) {
	return c.rdb.Get(context.Background(), key).Result()
}

func (c *RedisCache) Set(key, val string, ttl time.Duration) error {
	return c.rdb.Set(context.Background(), key, val, ttl).Err()
}

func (c *RedisCache) Del(key string) error {
	return c.rdb.Del(context.Background(), key).Err()
}