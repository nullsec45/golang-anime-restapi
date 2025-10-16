package domain

import (
	"context"
	"database/sql"
)

type User struct {
	Id         string        `db:"id"`
	Email      string        `db:"email"`
	Password   string        `db:"password"`
	CreatedAt  sql.NullTime  `db:"created_at"`
	UpdatedAt  sql.NullTime  `db:"updated_at"`
}

type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (User, error)
	Save(ctx context.Context, user *User)  error
	UpdatePassword(ctx context.Context, user *User) error
}