// internal/session/session_manager.go
package session

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	fibersession "github.com/gofiber/fiber/v2/middleware/session"
)

type Manager struct {
	store *fibersession.Store
}

func New(store *fibersession.Store) *Manager {
	return &Manager{store: store}
}

const (
	KeyUserID    = "user_id"
	KeyUserEmail = "user_email"
)

func (m *Manager) getSess(c *fiber.Ctx) (*fibersession.Session, error) {
	return m.store.Get(c)
}

func (m *Manager) Set(c *fiber.Ctx, key string, val any) error {
	sess, err := m.getSess(c)
	if err != nil {
		return err
	}
	sess.Set(key, val)
	return sess.Save()
}

func (m *Manager) SetMany(c *fiber.Ctx, values map[string]any) error {
	sess, err := m.getSess(c)
	if err != nil {
		return err
	}
	for k, v := range values {
		sess.Set(k, v)
	}
	return sess.Save()
}

func (m *Manager) GetString(c *fiber.Ctx, key string) (string, bool) {
	sess, err := m.getSess(c)
	if err != nil {
		return "", false
	}
	v := sess.Get(key)
	s, ok := v.(string)
	return s, ok
}

func (m *Manager) SetUser(c *fiber.Ctx, userID string, email string) error {
	if userID == "" {
		return errors.New("empty user id")
	}
	return m.SetMany(c, map[string]any{
		KeyUserID:    userID,
		KeyUserEmail: email,
	})
}

func (m *Manager) MustAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		uid, ok := m.GetString(c, KeyUserID)
		if !ok || uid == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		c.Locals(KeyUserID, uid)

		if email, ok := m.GetString(c, KeyUserEmail); ok {
			c.Locals(KeyUserEmail, email)
		}
		return c.Next()
	}
}

func (m *Manager) DeleteKeys(c *fiber.Ctx, keys ...string) error {
	sess, err := m.getSess(c)
	if err != nil {
		return err
	}
	for _, k := range keys {
		sess.Delete(k)
	}
	return sess.Save()
}

func (m *Manager) Destroy(c *fiber.Ctx) error {
	sess, err := m.getSess(c)
	if err != nil {
		return err
	}
	return sess.Destroy()
}

func (m *Manager) Renew(c *fiber.Ctx) error {
	sess, err := m.getSess(c)
	if err != nil {
		return err
	}
	if err := sess.Destroy(); err != nil {
		return err
	}
	_, err = m.getSess(c)
	return err
}

func (m *Manager) GetUser(c *fiber.Ctx) (userID string, email string, err error) {
	sess, err := m.getSess(c)
	if err != nil { return "", "", err }

	if v := sess.Get(KeyUserID); v != nil {
		userID, _ = v.(string)
	}
	if v := sess.Get(KeyUserEmail); v != nil {
		email, _ = v.(string)
	}
	if userID == "" {
		return "", "", errors.New("no user session")
	}
	return userID, email, nil
}