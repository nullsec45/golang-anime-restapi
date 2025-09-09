package domain

import (
	"context"
	"database/sql"
	"github.com/nullsec45/golang-anime-restapi/dto"
)

type AnimeEpisode struct {
	Id              string        `db:"id"`
	AnimeId         string        `db:"anime_id"`
	Number          int           `db:"number"`
	SeasonNumber    int           `db:"season_number"`     
	Title           string        `db:"title"`             
	Synopsis        string        `db:"synopsis"`          
	AirDate         sql.NullTime  `db:"air_date"` 
	DurationMinutes int  `db:"duration_minutes"`  
	IsSpecial       bool          `db:"is_special"`
	CreatedAt       sql.NullTime     `db:"created_at"`
	UpdatedAt       sql.NullTime     `db:"updated_at"`
}

type AnimeEpisodeRepository interface {
	FindByAnimeId(ctx context.Context, animeId string) ([]AnimeEpisode, error)
	FindById(ctx context.Context,  id string) (AnimeEpisode, error)
	Save(ctx context.Context, data *AnimeEpisode) error
	DeleteByAnimeId(ctx context.Context, animeId string) error
	DeleteById(ctx context.Context, id string) error
}

type AnimeEpisodeService interface {
	Create(ctx context.Context, req dto.CreateAnimeEpisodeRequest) error
	DeleteByAnimeId(ctx context.Context, animeId string) error
	DeleteById(ctx context.Context, id string) error
	// Delete(ctx context.Context, req dto.DeleteAnimeEpisodeRequest) error
}