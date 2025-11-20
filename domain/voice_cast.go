package domain

import (
	"context"
	"database/sql"
	"github.com/nullsec45/golang-anime-restapi/dto"
)

type VoiceCast struct {
	Id            string `db:"id"`
	AnimeId       string `db:"anime_id"`
	CharacterId   string `db:"character_id"`
	PersonId      string `db:"person_id"`
	Language      string `db:"language"`
	RoleNote      string `db:"role_note"`

	CreatedAt     sql.NullTime  `db:"created_at"`
	UpdatedAt     sql.NullTime  `db:"updated_at"`
}


type VoiceCastWithRelation struct {
	// kolom voice_casts
	Id          string  `db:"id"`
	AnimeId     string  `db:"anime_id"`
	CharacterId string  `db:"character_id"`
	PersonId    string  `db:"person_id"`
	Language    *string `db:"language"`
	RoleNote    *string `db:"role_note"`

	// kolom characters
	CharacterSlug        string  `db:"character_slug"`
	CharacterName        string  `db:"character_name"`
	CharacterNameNative  string  `db:"character_name_native"`
	CharacterDescription string  `db:"description"`

	// tambahkan field lain kalau perlu
	// CharacterImageURL *string `db:"character_image_url"`

	// kolom people
	PeopleSlug        string         `db:"people_slug"`
	PeopleName        string         `db:"people_name"`
	PeopleNameNative  string         `db:"people_name_native"`
	PeopleBirthday    sql.NullTime   `db:"birthday"`
	PeopleGender      dto.GenderType `db:"gender"`
	PeopleCountry     string         `db:"country"`
	PeopleSiteURL     string         `db:"site_url"`
	PeopleBiography   string         `db:"biography"`
	// PeopleImageURL *string `db:"people_image_url"`
}

type VoiceCastRepository interface {
	FindById(ctx context.Context,  id string) (VoiceCast, error)
	FindByAnimeId(ctx context.Context, animeId string)([]VoiceCastWithRelation, error)
	FindUnique(ctx context.Context,  animeId string, characterId string, personId string) (VoiceCast,  bool, error)
	Save(ctx context.Context, data *VoiceCast) error
	Update(ctx context.Context, data *VoiceCast) error
	// DeleteByAnimeId(ctx context.Context, animeId string) error
	// DeleteByCharacterId(ctx context.Context, id string) error
	// DeleteByPersonId(ctx context.Context, id string) error
	DeleteById(ctx context.Context, id string) error
}


type VoiceCastService interface {
	// Index(ctx context.Context, opts CharacterListOptions) (dto.Paginated[dto.CharacterData], error)
	// Show(ctx context.Context, param string) (dto.CharacterData, error)
	Create(ctx context.Context, req dto.CreateVoiceCastRequest) error
	Update(ctx context.Context, req dto.UpdateVoiceCastRequest) error
	// DeleteByAnimeId(ctx context.Context, animeId string) error
	// DeleteByTagId(ctx context.Context, id string) error
	DeleteById(ctx context.Context, id string) error
}
