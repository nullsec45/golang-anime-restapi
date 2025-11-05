package dto

import (
	"reflect"
	// "time"
	"sync"
	"github.com/go-playground/validator/v10"
)

var (
	episodeOnce sync.Once
	episodeV    *validator.Validate
)

func episodeValidator() *validator.Validate {
	episodeOnce.Do(func() {
		v := validator.New(validator.WithRequiredStructEnabled())

		v.RegisterCustomTypeFunc(func(field reflect.Value) interface{} {
			if !field.IsValid() {
				return nil
			}
			if ft, ok := field.Interface().(FlexibleTime); ok {
				return ft.Time
			}
			if p, ok := field.Interface().(*FlexibleTime); ok && p != nil {
				return p.Time
			}
			return nil
		}, FlexibleTime{}, &FlexibleTime{})

		v.RegisterValidation("dateymd", func(fl validator.FieldLevel) bool {
			switch x := fl.Field().Interface().(type) {
			case FlexibleTime:
				return x.IsZero() || x.Layout == "2006-01-02"
			case *FlexibleTime:
				return x == nil || x.IsZero() || x.Layout == "2006-01-02"
			default:
				return false
			}
		})

		episodeV = v
	})
	return episodeV
}

type AnimeEpisodeData struct {
	Id              string         `json:"id"`
	AnimeId         string         `json:"anime_id" validate:"required,uuid4"`
	Number          int            `json:"number" validate:"required,min=1"`
	SeasonNumber    int            `json:"season_number,omitempty" validate:"omitempty,min=1"`
	Title           string         `json:"title,omitempty" validate:"omitempty"`
	Synopsis        string         `json:"synopsis,omitempty" validate:"omitempty"`
	AirDate         *FlexibleTime  `json:"air_date,omitempty" validate:"omitempty"`
	DurationMinutes int            `json:"duration_minutes,omitempty" validate:"omitempty,min=1"`
	IsSpecial       bool           `json:"is_special"`
	Video           string         `json:"video"`
}

type CreateAnimeEpisodeRequest struct {
	AnimeId         string         `json:"anime_id" validate:"required,uuid4"`
	Number          int            `json:"number" validate:"required,min=1"`
	SeasonNumber    int            `json:"season_number,omitempty" validate:"omitempty,min=1"`
	Title           string         `json:"title,omitempty" validate:"omitempty"`
	Synopsis        string         `json:"synopsis,omitempty" validate:"omitempty"`
	AirDate         *FlexibleTime  `json:"air_date,omitempty" validate:"omitempty"`
	DurationMinutes int            `json:"duration_minutes,omitempty" validate:"omitempty,min=1"`
	IsSpecial       bool           `json:"is_special,omitempty"`
	Video           string         `json:"video,omitempty"`
}

func (r *CreateAnimeEpisodeRequest) Validate() error {
	return episodeValidator().Struct(r)
}

type UpdateAnimeEpisodeRequest struct {
	Id              string         `json:"id"`
	AnimeId         string         `json:"anime_id" validate:"required,uuid4"`
	Number          int            `json:"number" validate:"required,min=1"`
	SeasonNumber    int            `json:"season_number,omitempty" validate:"omitempty,min=1"`
	Title           string         `json:"title,omitempty" validate:"omitempty"`
	Synopsis        string         `json:"synopsis,omitempty" validate:"omitempty"`
	AirDate         *FlexibleTime  `json:"air_date,omitempty" validate:"omitempty"`
	DurationMinutes int            `json:"duration_minutes,omitempty" validate:"omitempty,min=1"`
	IsSpecial       bool           `json:"is_special"`
	Video           string         `json:"video,omitempty"`
}