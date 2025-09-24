package domain

import (
	"context"
	"database/sql"
	"github.com/nullsec45/golang-anime-restapi/dto"
)

type AnimeTag struct {
	Id         string        `db:"id"`
	Slug       string        `db:"slug"`
	Name       string        `db:"name"`
	CreatedAt  sql.NullTime  `db:"created_at"`
	UpdatedAt  sql.NullTime  `db:"updated_at"`
}

type AnimeTagRepository interface {
	FindAll(ctx context.Context) ([]AnimeTag, error)
	FindById(ctx context.Context, id string) (AnimeTag, error)
	FindByAnimeId(ctx context.Context, animeId string)([]AnimeTag, error)
	Save(ctx context.Context, genre *AnimeTag) error
	Update(ctx context.Context, genre *AnimeTag) error
	Delete(ctx context.Context, id string) error
}

type AnimeTagService interface {	
	Index(ctx context.Context) ([]dto.AnimeTagData, error)
	Show(ctx context.Context, id string) (dto.AnimeTagData, error)
	Create(ctx context.Context, req dto.CreateAnimeTagRequest) error
	Update(ctx context.Context, req dto.UpdateAnimeTagRequest) error
	Delete(ctx context.Context, id string) error
}