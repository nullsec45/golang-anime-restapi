package domain

import (
	"context"
	"database/sql"
	"github.com/nullsec45/golang-anime-restapi/dto"
)

type AnimeStudios struct {
	Id         string       `db:"id"`
	AnimeId    string       `db:"anime_id"`
	StudioId   string       `db:"studio_id"`
	Role       string       `db:"role"`
    CreatedAt  sql.NullTime `db:"created_at"`
	UpdatedAt  sql.NullTime `db:"updated_at"`
}

type AnimeStudiosRepository interface {
	FindById(ctx context.Context,  id string) (AnimeStudios, error)
	FindByAnimeAndStudioId(ctx context.Context,  animeId string, studioId string) (AnimeStudios, bool, error)
	Save(ctx context.Context, data *AnimeStudios) error
	Update(ctx context.Context, data *AnimeStudios) error
	DeleteByAnimeId(ctx context.Context, animeId string) error
	DeleteByStudioId(ctx context.Context, studioId string) error
	DeleteById(ctx context.Context, id string) error
}

type AnimeStudiosService interface {
	Create(ctx context.Context, req dto.CreateAnimeStudiosRequest) error
	Update(ctx context.Context, req dto.UpdateAnimeStudiosRequest) error
	DeleteByAnimeId(ctx context.Context, animeId string) error
	DeleteByStudioId(ctx context.Context, studioId string) error
	DeleteById(ctx context.Context, id string) error
}