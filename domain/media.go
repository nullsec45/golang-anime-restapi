package domain

import (
	"context"
	"database/sql"
	"github.com/nullsec45/golang-anime-restapi/dto"
)

type Media struct {
	Id        string       `db:"id"`
	Path      string       `db:"path"`
	CreatedAt sql.NullTime `db:"created_at"` 
}

type MediaRepository interface {
	FindById(ctx context.Context, id string) (Media, error)
	// FindByIds(ctx context.Context, ids []string)([]Media, error)
	Save(ctx context.Context, media *Media) error
	Delete(ctx context.Context, id string) error
}

type MediaService interface {
	Create(ctx context.Context, req dto.CreateMediaRequest) (dto.MediaData, error)
	Delete(ctx context.Context, id string) error
}