package session

import (
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis"
)

func NewWithRedisStorage(storage *redis.Storage, secureCookie bool) *Manager {
	store := session.New(session.Config{
		Storage:        storage,
		CookieHTTPOnly: true,
		CookieSecure:   secureCookie, // true di prod (HTTPS)
		CookieSameSite: "Lax",        // atau "Strict"/"None"
		// CookieDomain: "example.com",
		// Expiration: 24 * time.Hour, // opsional
	})
	return New(store)
}
