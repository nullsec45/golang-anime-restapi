package domain

import (
	"context"
	"database/sql"
	"github.com/nullsec45/golang-anime-restapi/dto"
)

type Character struct {
	Id                     string          `db:"id"`
	Slug                   string          `db:"slug"`
	Name                   string          `db:"name"`
	NameNative             string          `db:"name_native"`
	Description            string          `db:"description"`
	CharacterImage    	   sql.NullString  `db:"character_image"`
	CreatedAt              sql.NullTime    `db:"created_at"`
	UpdatedAt              sql.NullTime    `db:"updated_at"`
}

type CharacterFilter struct {
	Search string
}

type CharacterListOptions struct {
	Pagination dto.PaginationQuery
	Filter     CharacterFilter
}

type CharacterRepository interface {
	FindAll(ctx context.Context, opts CharacterListOptions) ([]Character, int64, error)
	FindById(ctx context.Context, id string) (Character, error)
	FindBySlug(ctx context.Context, slug string) (Character, error)
	Save(ctx context.Context, anime *Character) error
	Update(ctx context.Context, anime *Character) error
	Delete(ctx context.Context, id string) error
}

type CharacterService interface {
	Index(ctx context.Context, opts CharacterListOptions) (dto.Paginated[dto.CharacterData], error)
	Show(ctx context.Context, param string) (dto.CharacterData, error)
	Create(ctx context.Context, req dto.CreateCharacterRequest) error
	Update(ctx context.Context, req dto.UpdateCharacterRequest) error
	Delete(ctx context.Context, id string) error
}
