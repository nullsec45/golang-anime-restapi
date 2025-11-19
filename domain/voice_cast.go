package domain

import (
	"context"
	"database/sql"
	"github.com/nullsec45/golang-anime-restapi/dto"
)

type VoiceCast struct {
	Id                     string        `db:"id"`
	AnimeId                string        `db:"anime_id"`
	CharacterId            string        `db:"character_id"`
	PersonId               string        `db:"person_id"`
	Language               string        `db:"language"`
	RoleNote               string        `db:"role_note"`
	CreatedAt              sql.NullTime  `db:"created_at"`
	UpdatedAt              sql.NullTime  `db:"updated_at"`
}


type VoiceCastRepository interface {
	FindById(ctx context.Context,  id string) (VoiceCast, error)
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
