package domain

import (
	"context"
	"database/sql"
	"github.com/nullsec45/golang-anime-restapi/dto"
	"time"
)

type Media struct {
	Id        string       `db:"id"`
	Path      string       `db:"path"`
	// Url       string       `db:"url"`
	CreatedAt sql.NullTime `db:"created_at"` 
	UpdatedAt sql.NullTime `db:"updated_at"`
}

type MediaRepository interface {
	FindById(ctx context.Context, id string) (Media, error)
	FindByIds(ctx context.Context, ids []string)([]Media, error)
	Save(ctx context.Context, media *Media) error
	Update(ctx context.Context, media *Media) (*Media, error)
	Delete(ctx context.Context, id string) error
}

type MediaService interface {
	Create(ctx context.Context, req dto.CreateMediaRequest) (dto.MediaData, error)

	Show(ctx context.Context, id string) (dto.MediaData, error)
	
	View(ctx context.Context, id string) (absPath string, filename string, modTime time.Time, err error)
	Update(ctx context.Context, req dto.UpdateMediaRequest) (dto.MediaData, error)
	Delete(ctx context.Context, id string) (path string, err error)
}