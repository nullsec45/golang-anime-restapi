package dto

// import (
// 	"reflect"
// 	"time"
// 	"github.com/go-playground/validator/v10"
// )

type CreateAnimeEpisodeRequest struct {
	AnimeId         string `json:"anime_id" validate:"required,uuid4"`
	Number          int    `json:"number" validate:"required,min=1"`
	SeasonNumber    int    `json:"season_number,omitempty" validate:"omitempty,min=1"`
	Title           string `json:"title,omitempty" validate:"omitempty"`
	Synopsis        string `json:"synopsis,omitempty" validate:"omitempty"`
	AirDate         string `json:"air_date,omitempty" validate:"omitempty,datetime=2006-01-02"`
	DurationMinutes int    `json:"duration_minutes,omitempty" validate:"omitempty,min=1"`
	IsSpecial       bool   `json:"is_special,omitempty"`
}

// type DeleteAnimeEpisodeRequest struct {
	// BookId string
	// Id string
// }

// type UpdateEpisodeRequest struct {
// 	// Boleh partial update; semua opsional dengan validasi bila ada
// 	Number          int    `json:"number,omitempty" validate:"omitempty,min=1"`
// 	SeasonNumber    int    `json:"season_number,omitempty" validate:"omitempty,min=1"`
// 	Title           string `json:"title,omitempty" validate:"omitempty"`
// 	Synopsis        string `json:"synopsis,omitempty" validate:"omitempty"`
// 	AirDate         string `json:"air_date,omitempty" validate:"omitempty,datetime=2006-01-02"`
// 	DurationMinutes int    `json:"duration_minutes,omitempty" validate:"omitempty,min=1"`
// 	IsSpecial       bool   `json:"is_special,omitempty"`
// }