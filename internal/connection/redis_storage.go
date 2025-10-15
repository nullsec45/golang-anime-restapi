// internal/connection/redis_storage.go
package connection

import (
	"fmt"
	"strconv"
	"time"

	fiberredis "github.com/gofiber/storage/redis"
	"github.com/nullsec45/golang-anime-restapi/internal/config"
)

func GetRedisStorage(conf config.Redis) (*fiberredis.Storage, error) {
	if conf.Host == "" || conf.Port == "" {
		return nil, fmt.Errorf("redis host/port is empty")
	}

	port, err := strconv.Atoi(conf.Port)
	if err != nil {
		return nil, fmt.Errorf("invalid redis port: %w", err)
	}

	st := fiberredis.New(fiberredis.Config{
		Host:     conf.Host,
		Port:     port,
		Password: conf.Password, // kosongkan jika tidak pakai password
		// Database: 0,
		// TLSConfig: nil,
	})

	key := "healthcheck:" + strconv.FormatInt(time.Now().UnixNano(), 10)

	if err := st.Set(key, []byte("ok"), 5*time.Second); err != nil {
		_ = st.Close()
		return nil, fmt.Errorf("redis healthcheck set failed: %w", err)
	}
	if _, err := st.Get(key); err != nil {
		_ = st.Close()
		return nil, fmt.Errorf("redis healthcheck get failed: %w", err)
	}
	_ = st.Delete(key)

	return st, nil
}

func CloseRedisStorage(st *fiberredis.Storage) error {
	if st == nil {
		return nil
	}
	return st.Close()
}
