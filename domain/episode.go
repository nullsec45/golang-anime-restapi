package domain

import (
	"context"
	"database/sql"
	"github.com/nullsec45/golang-anime-restapi/dto"
)

type AnimeEpisode struct {
	Id              string          `db:"id"`
	AnimeId         string          `db:"anime_id"`
	Number          int             `db:"number"`
	SeasonNumber    int             `db:"season_number"`     
	Title           string          `db:"title"`             
	Synopsis        string          `db:"synopsis"`          
	AirDate         sql.NullTime    `db:"air_date"` 
	DurationMinutes int             `db:"duration_minutes"`  
	IsSpecial       bool            `db:"is_special"`
	Video    	    sql.NullString  `db:"video"`
	CreatedAt       sql.NullTime    `db:"created_at"`
	UpdatedAt       sql.NullTime    `db:"updated_at"`
}

type EpisodeFilter struct {
	Search string
}

type EpisodeListOptions struct {
	Pagination dto.PaginationQuery
	Filter     EpisodeFilter
}

type AnimeEpisodeRepository interface {
	FindAll(ctx context.Context, animeId string, epts EpisodeListOptions) ([]AnimeEpisode, int64, error)
	FindByAnimeId(ctx context.Context, animeId string) ([]AnimeEpisode, error)
	FindById(ctx context.Context,  id string) (AnimeEpisode, error)
	Save(ctx context.Context, data *AnimeEpisode) error

	Update(ctx context.Context, data *AnimeEpisode) error
	DeleteByAnimeId(ctx context.Context, animeId string) error
	DeleteById(ctx context.Context, id string) error
}

type AnimeEpisodeService interface {
	Index(ctx context.Context, animeId string, epts EpisodeListOptions	) (dto.Paginated[dto.AnimeEpisodeData], error)
	Show(ctx context.Context, id string) (dto.AnimeEpisodeData, error)
	Create(ctx context.Context, req dto.CreateAnimeEpisodeRequest) error
	Update(ctx context.Context, req dto.UpdateAnimeEpisodeRequest) error
	DeleteByAnimeId(ctx context.Context, animeId string) error
	DeleteById(ctx context.Context, id string) error
}