package dto

import "time"

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
	TitleNative            *string                `json:"title_native,omitempty"`
	TitleEnglish           *string                `json:"title_english,omitempty"`
	Synopsis               *string                `json:"synopsis,omitempty"`
	Type                   AnimeType              `json:"type"`
	Season                 *Season                `json:"season,omitempty"`
	SeasonYear             *int16                 `json:"season_year,omitempty"`
	Status                 AnimeStatus            `json:"status"`
	AgeRating              *AgeRating             `json:"age_rating,omitempty"`
	TotalEpisodes          *int                   `json:"total_episodes,omitempty"`
	AverageDurationMinutes *int                   `json:"average_duration_minutes,omitempty"`
	Country                string                 `json:"country"` 
	PremieredAt            *time.Time             `json:"premiered_at,omitempty"`
	EndedAt                *time.Time             `json:"ended_at,omitempty"`
	Popularity             int                    `json:"popularity"`
	ScoreAvg               *float32               `json:"score_avg,omitempty"`
	AltTitles              map[string]interface{} `json:"alt_titles"`  
	ExternalIDs            map[string]interface{} `json:"external_ids"`
}


type CreateAnimeRequest struct {
	Id                     string                  `json:"id"`
	Slug                   *string                 `json:"slug,omitempty" validate:"omitempty"`
	TitleRomaji            string                  `json:"title_romaji" validate:"required,min=1"`
	TitleNative            *string                 `json:"title_native,omitempty" validate:"omitempty"`
	TitleEnglish           *string                 `json:"title_english,omitempty" validate:"omitempty"`
	Synopsis               *string                 `json:"synopsis,omitempty" validate:"omitempty"`
	Type                   AnimeType        `json:"type" validate:"required,oneof=TV Movie OVA ONA Special"`
	Season                 *Season          `json:"season,omitempty" validate:"omitempty,oneof=Winter Spring Summer Fall"`
	SeasonYear             *int16                  `json:"season_year,omitempty" validate:"omitempty,gte=1917,lte=2100"`
	Status                 AnimeStatus      `json:"status" validate:"required,oneof=Upcoming Airing Finished Hiatus"`
	AgeRating              *AgeRating       `json:"age_rating,omitempty" validate:"omitempty,oneof=G PG PG-13 R R+ Rx"`
	TotalEpisodes          *int                    `json:"total_episodes,omitempty" validate:"omitempty,gte=0"`
	AverageDurationMinutes *int                    `json:"average_duration_minutes,omitempty" validate:"omitempty,gte=0"`
	Country                *string                 `json:"country,omitempty" validate:"omitempty,len=2"` 
	PremieredAt            *string                 `json:"premiered_at,omitempty" validate:"omitempty,datetime=2006-01-02"`
	EndedAt                *string                 `json:"ended_at,omitempty" validate:"omitempty,datetime=2006-01-02"`
	Popularity             *int                    `json:"popularity,omitempty" validate:"omitempty,gte=0"`
	ScoreAvg               *float32                `json:"score_avg,omitempty" validate:"omitempty,gte=0,lte=9.99"` 
	AltTitles              map[string]interface{}  `json:"alt_titles,omitempty"`
	ExternalIDs            map[string]interface{}  `json:"external_ids,omitempty"`
}

type UpdateAnimeRequest struct {
	Id                     string                  `json:"id"`
	Slug                   *string                 `json:"slug,omitempty" validate:"omitempty"`
	TitleRomaji            string                  `json:"title_romaji" validate:"required,min=1"`
	TitleNative            *string                 `json:"title_native,omitempty" validate:"omitempty"`
	TitleEnglish           *string                 `json:"title_english,omitempty" validate:"omitempty"`
	Synopsis               *string                 `json:"synopsis,omitempty" validate:"omitempty"`
	Type                   AnimeType        `json:"type" validate:"required,oneof=TV Movie OVA ONA Special"`
	Season                 *Season          `json:"season,omitempty" validate:"omitempty,oneof=Winter Spring Summer Fall"`
	SeasonYear             *int16                  `json:"season_year,omitempty" validate:"omitempty,gte=1917,lte=2100"`
	Status                 AnimeStatus      `json:"status" validate:"required,oneof=Upcoming Airing Finished Hiatus"`
	AgeRating              *AgeRating       `json:"age_rating,omitempty" validate:"omitempty,oneof=G PG PG-13 R R+ Rx"`
	TotalEpisodes          *int                    `json:"total_episodes,omitempty" validate:"omitempty,gte=0"`
	AverageDurationMinutes *int                    `json:"average_duration_minutes,omitempty" validate:"omitempty,gte=0"`
	Country                *string                 `json:"country,omitempty" validate:"omitempty,len=2"` 
	PremieredAt            *string                 `json:"premiered_at,omitempty" validate:"omitempty,datetime=2006-01-02"`
	EndedAt                *string                 `json:"ended_at,omitempty" validate:"omitempty,datetime=2006-01-02"`
	Popularity             *int                    `json:"popularity,omitempty" validate:"omitempty,gte=0"`
	ScoreAvg               *float32                `json:"score_avg,omitempty" validate:"omitempty,gte=0,lte=9.99"` 
	AltTitles              map[string]interface{}  `json:"alt_titles,omitempty"`
	ExternalIDs            map[string]interface{}  `json:"external_ids,omitempty"`
}