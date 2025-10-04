package domain

import (
	"context"
	"database/sql"
	"github.com/nullsec45/golang-anime-restapi/dto"
)

type AnimeStudio struct {
	Id         string        `db:"id"`
	Slug       string        `db:"slug"`
	Name       string        `db:"name"`
	Country    string 	     `db:"country"`
	SiteURL	   string 	     `db:"site_url"`
	CreatedAt  sql.NullTime  `db:"created_at"`
	UpdatedAt  sql.NullTime  `db:"updated_at"`
}

type AnimeStudioRepository interface {
	FindAll(ctx context.Context) ([]AnimeStudio, error)
	FindById(ctx context.Context, id string) (AnimeStudio, error)
	FindBySlug(ctx context.Context, slug string) (AnimeStudio, error)
	FindByAnimeId(ctx context.Context, animeId string)([]AnimeStudio, error)
	Save(ctx context.Context, studio *AnimeStudio) error
	Update(ctx context.Context, studio *AnimeStudio) error
	Delete(ctx context.Context, id string) error
}

type AnimeStudioService interface {	
	Index(ctx context.Context) ([]dto.AnimeStudioData, error)
	Show(ctx context.Context, param string) (dto.AnimeStudioData, error)
	Create(ctx context.Context, req dto.CreateAnimeStudioRequest) error
	Update(ctx context.Context, req dto.UpdateAnimeStudioRequest) error
	Delete(ctx context.Context, id string) error
}