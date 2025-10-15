package connection

import (
	"context"
	"fmt"
	"net"
	"time"
	goredis "github.com/redis/go-redis/v9"
	"github.com/nullsec45/golang-anime-restapi/internal/config"
)

func GetRedisClient(conf config.Redis) (*goredis.Client, error) {
	addr := net.JoinHostPort(conf.Host, conf.Port)

	rdb := goredis.NewClient(&goredis.Options{
		Addr:addr,
		Password:conf.Password,
		DB:0,
		PoolSize:20,
		MinIdleConns:4,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("redis ping failed:%w", err)
	}
	return rdb, nil
}