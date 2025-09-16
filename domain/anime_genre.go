package domain

import (
	"context"
	"database/sql"
	"github.com/nullsec45/golang-anime-restapi/dto"
)

type AnimeGenres struct {
	Id         string       `db:"id"`
	AnimeId    string       `db:"anime_id"`
	GenreId    string       `db:"genre_id"`
    CreatedAt  sql.NullTime `db:"created_at"`
	UpdatedAt  sql.NullTime `db:"updated_at"`
}

type AnimeGenresRepository interface {
	Save(ctx context.Context, data *AnimeGenres) error
	Update(ctx context.Context, data *AnimeGenres) error
	DeleteByAnimeId(ctx context.Context, animeId string) error
	DeleteByGenreId(ctx context.Context, id string) error
	DeleteById(ctx context.Context, id string) error
}

type AnimeGenresService interface {
	Create(ctx context.Context, req dto.CreateAnimeGenresRequest) error
	Update(ctx context.Context, req dto.CreateAnimeGenresRequest) error
	DeleteByAnimeId(ctx context.Context, animeId string) error
	DeleteByGenreId(ctx context.Context, id string) error
	DeleteById(ctx context.Context, id string) error
}