package domain

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/nullsec45/golang-anime-restapi/dto"
)

type Anime struct {
	Id                     string                 `db:"id"`
	Slug                   string                 `db:"slug"`
	TitleRomaji            string                 `db:"title_romaji"`
	TitleNative            sql.NullString         `db:"title_native"`
	TitleEnglish           sql.NullString         `db:"title_english"`
	Synopsis               sql.NullString         `db:"synopsis"`
	Type                   dto.AnimeType          `db:"type"`
	Season                 *dto.Season            `db:"season"`
	SeasonYear             int                    `db:"season_year"`
	Status                 dto.AnimeStatus        `db:"status"`
	AgeRating              *dto.AgeRating         `db:"age_rating"`
	TotalEpisodes          int                    `db:"total_episodes"`
	AverageDurationMinutes int                    `db:"average_duration_minutes"`
	Country                string                 `db:"country"` 
	PremieredAt            sql.NullTime           `db:"premiered_at"`
	EndedAt                sql.NullTime           `db:"ended_at"`
	Popularity             int                    `db:"popularity"`
	ScoreAvg               float32                `db:"score_avg"`
	AltTitles              json.RawMessage        `db:"alt_titles"`  
	ExternalIDs            json.RawMessage        `db:"external_ids"`
	CreatedAt              sql.NullTime           `db:"created_at"`
	UpdatedAt              sql.NullTime           `db:"updated_at"`
}

type AnimeRepository interface {
	FindAll(ctx context.Context) ([]Anime, error)
	FindById(ctx context.Context, id string) (Anime, error)
	Save(ctx context.Context, anime *Anime) error
	Update(ctx context.Context, anime *Anime) error
	Delete(ctx context.Context, id string) error
}

type AnimeService interface {
	Index(ctx context.Context) ([]dto.AnimeData, error)
	Show(ctx context.Context, id string) (dto.AnimeData, error)
	Create(ctx context.Context, req dto.CreateAnimeRequest) error
	Update(ctx context.Context, req dto.UpdateAnimeRequest) error
	Delete(ctx context.Context, id string) error
}
