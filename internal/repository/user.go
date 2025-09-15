package repository 

import (
	"context"
	"database/sql"
	"github.com/doug-martin/goqu/v9"
	"github.com/nullsec45/golang-anime-restapi/domain"
)

type UserRepository struct {
	db *goqu.Database
}

func NewUser(con *sql.DB)domain.UserRepository{
	return &UserRepository{
		db: goqu.New("default", con),
	}
}

func (userRepo UserRepository) FindByEmail(ctx context.Context, email string) (user domain.User, err error) {
	dataset := userRepo.db.From("users").Where(goqu.C("email").Eq(email))
	_, err = dataset.ScanStructContext(ctx, &user)
	return
}

func (userRepo *UserRepository) Save (ctx context.Context, user *domain.User) error {
	executor := userRepo.db.Insert("users").Rows(user).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}