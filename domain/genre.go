package domain

import (
	"context"
	"database/sql"
	"github.com/nullsec45/golang-anime-restapi/dto"
)

type AnimeGenre struct {
	Id         string        `db:"id"`
	Slug       string        `db:"slug"`
	Name       string        `db:"name"`
	CreatedAt  sql.NullTime  `db:"created_at"`
	UpdatedAt  sql.NullTime  `db:"updated_at"`
}

type AnimeGenreRepository interface {
	FindAll(ctx context.Context) ([]AnimeGenre, error)
	FindById(ctx context.Context, id string) (AnimeGenre, error)
	FindByAnimeId(ctx context.Context, animeId string)([]AnimeGenre, error)
	Save(ctx context.Context, genre *AnimeGenre) error
	Update(ctx context.Context, genre *AnimeGenre) error
	Delete(ctx context.Context, id string) error
}

type AnimeGenreService interface {	
	Index(ctx context.Context) ([]dto.AnimeGenreData, error)
	Show(ctx context.Context, id string) (dto.AnimeGenreData, error)
	Create(ctx context.Context, req dto.CreateAnimeGenreRequest) error
	Update(ctx context.Context, req dto.UpdateAnimeGenreRequest) error
	Delete(ctx context.Context, id string) error
}