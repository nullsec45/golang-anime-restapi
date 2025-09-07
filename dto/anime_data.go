package dto

import (
	// "database/sql"
	// "time"
	// "encoding/json"
	"reflect"
	"time"
	"github.com/go-playground/validator/v10"
)

type AnimeType string
const (
	TypeTv AnimeType="TV"
	TypeMovie AnimeType="Movie"
	TypeOVA    AnimeType = "OVA"
	TypeONA    AnimeType = "ONA"
	TypeSpecial AnimeType = "Special"
)

type Season string
const (
	SeasonWinter Season = "Winter"
	SeasonSpring Season = "Spring"
	SeasonSummer Season = "Summer"
	SeasonFall   Season = "Fall"
)

type AnimeStatus string
const (
	StatusUpcoming AnimeStatus = "Upcoming"
	StatusAiring   AnimeStatus = "Airing"
	StatusFinished AnimeStatus = "Finished"
	StatusHiatus   AnimeStatus = "Hiatus"
)

type AgeRating string
const (
	AgeG   AgeRating = "G"
	AgePG  AgeRating = "PG"
	AgePG13 AgeRating = "PG-13"
	AgeR   AgeRating = "R"
	AgeRPlus AgeRating = "R+"
	AgeRx  AgeRating = "Rx"
)

type AnimeData struct {
	Id                     string                  `json:"id"`
	Slug                   string                 `json:"slug"`
	TitleRomaji            string                 `json:"title_romaji"`
	TitleNative            string                `json:"title_native,omitempty"`
	TitleEnglish           string                `json:"title_english,omitempty"`
	Synopsis               string                `json:"synopsis,omitempty"`
	Type                   AnimeType              `json:"type"`
	Season                 *Season                `json:"season,omitempty"`
	SeasonYear             int                  `json:"season_year,omitempty"`
	Status                 AnimeStatus            `json:"status"`
	AgeRating              *AgeRating             `json:"age_rating,omitempty"`
	TotalEpisodes          int                   `json:"total_episodes,omitempty"`
	AverageDurationMinutes int                   `json:"average_duration_minutes,omitempty"`
	Country                string                 `json:"country"` 
	PremieredAt            *FlexibleTime             `json:"premiered_at,omitempty"`
	EndedAt                *FlexibleTime             `json:"ended_at,omitempty"`
	Popularity             int                   `json:"popularity"`
	ScoreAvg               float32               `json:"score_avg,omitempty"`
	AltTitles              AltTitles       `json:"alt_titles"`  
	ExternalIDs            ExternalIDs       `json:"external_ids"`
}

func NewValidator() *validator.Validate {
	v := validator.New(validator.WithRequiredStructEnabled())

	// Ajari validator cara “melihat” FlexibleTime sebagai time.Time
	v.RegisterCustomTypeFunc(func(field reflect.Value) interface{} {
		switch ft := field.Interface().(type) {
		case FlexibleTime:
			return ft.Time
		case *FlexibleTime:
			if ft == nil {
				// biar omitempty jalan
				return time.Time{}
			}
			return ft.Time
		default:
			return nil
		}
	}, FlexibleTime{}, (*FlexibleTime)(nil))

	return v
}


type CreateAnimeRequest struct {
	Id                     string                  `json:"id"`
	Slug                   string                  `json:"slug,omitempty" validate:"omitempty"`
	TitleRomaji            string                  `json:"title_romaji" validate:"required,min=1"`
	TitleNative            string                  `json:"title_native,omitempty" validate:"omitempty"`
	TitleEnglish           string                  `json:"title_english,omitempty" validate:"omitempty"`
	Synopsis               string                  `json:"synopsis,omitempty" validate:"omitempty"`
	Type                   AnimeType               `json:"type" validate:"required,oneof=TV Movie OVA ONA Special"`
	Season                 *Season                 `json:"season,omitempty" validate:"omitempty,oneof=Winter Spring Summer Fall"`
	SeasonYear             int                     `json:"season_year,omitempty" validate:"omitempty,gte=1917,lte=2100"`
	Status                 AnimeStatus             `json:"status" validate:"required,oneof=Upcoming Airing Finished Hiatus"`
	AgeRating              *AgeRating              `json:"age_rating,omitempty" validate:"omitempty,oneof=G PG PG-13 R R+ Rx"`
	TotalEpisodes          int                     `json:"total_episodes,omitempty" validate:"omitempty,gte=0"`
	AverageDurationMinutes int                     `json:"average_duration_minutes,omitempty" validate:"omitempty,gte=0"`
	Country                string                  `json:"country,omitempty" validate:"omitempty,len=2"` 
	PremieredAt            *FlexibleTime               `json:"premiered_at,omitempty" validate:"omitempty"`
	EndedAt                *FlexibleTime               `json:"ended_at,omitempty" validate:"omitempty"`
	Popularity             int                    `json:"popularity,omitempty" validate:"omitempty,gte=0"`
	ScoreAvg               float32                `json:"score_avg,omitempty" validate:"omitempty,gte=0,lte=9.99"` 
	AltTitles              AltTitles  `json:"alt_titles,omitempty"`
	ExternalIDs            ExternalIDs  `json:"external_ids,omitempty"`
}

type UpdateAnimeRequest struct {
	Id                     string                  `json:"id"`
	Slug                   string                  `json:"slug,omitempty" validate:"omitempty"`
	TitleRomaji            string                  `json:"title_romaji" validate:"required,min=1"`
	TitleNative            string                  `json:"title_native,omitempty" validate:"omitempty"`
	TitleEnglish           string                  `json:"title_english,omitempty" validate:"omitempty"`
	Synopsis               string                  `json:"synopsis,omitempty" validate:"omitempty"`
	Type                   AnimeType               `json:"type" validate:"required,oneof=TV Movie OVA ONA Special"`
	Season                 *Season                 `json:"season,omitempty" validate:"omitempty,oneof=Winter Spring Summer Fall"`
	SeasonYear             int                     `json:"season_year,omitempty" validate:"omitempty,gte=1917,lte=2100"`
	Status                 AnimeStatus             `json:"status" validate:"required,oneof=Upcoming Airing Finished Hiatus"`
	AgeRating              *AgeRating              `json:"age_rating,omitempty" validate:"omitempty,oneof=G PG PG-13 R R+ Rx"`
	TotalEpisodes          int                     `json:"total_episodes,omitempty" validate:"omitempty,gte=0"`
	AverageDurationMinutes int                     `json:"average_duration_minutes,omitempty" validate:"omitempty,gte=0"`
	Country                string                  `json:"country,omitempty" validate:"omitempty,len=2"` 
	PremieredAt            *FlexibleTime               `json:"premiered_at,omitempty" validate:"omitempty"`
	EndedAt                *FlexibleTime               `json:"ended_at,omitempty" validate:"omitempty"`
	Popularity             int                    `json:"popularity,omitempty" validate:"omitempty,gte=0"`
	ScoreAvg               float32                `json:"score_avg,omitempty" validate:"omitempty,gte=0,lte=9.99"` 
	AltTitles              AltTitles  `json:"alt_titles,omitempty"`
	ExternalIDs            ExternalIDs  `json:"external_ids,omitempty"`
}