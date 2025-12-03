package domain

import (
	"context"
	"database/sql"
	"github.com/nullsec45/golang-anime-restapi/dto"
)

type AnimeTags struct {
	Id         string       `db:"id"`
	AnimeId    string       `db:"anime_id"`
	TagId      string       `db:"tag_id"`
    CreatedAt  sql.NullTime `db:"created_at"`
	UpdatedAt  sql.NullTime `db:"updated_at"`
}

type AnimeTagsRepository interface {
	FindById(ctx context.Context,  id string) (AnimeTags, error)
	FindByAnimeId(ctx context.Context, animeId string)([]AnimeTag, error)
	FindByAnimeIDs(ctx context.Context, animeIDs []string)(map[string][]AnimeTag, error)
	FindByAnimeAndTagId(ctx context.Context,  animeId string, tagId string) (AnimeTags,  bool, error)
	Save(ctx context.Context, data *AnimeTags) error
	Update(ctx context.Context, data *AnimeTags) error
	DeleteByAnimeId(ctx context.Context, animeId string) error
	DeleteByTagId(ctx context.Context, id string) error
	DeleteById(ctx context.Context, id string) error
}

type AnimeTagsService interface {
	Create(ctx context.Context, req dto.CreateAnimeTagsRequest) error
	Update(ctx context.Context, req dto.UpdateAnimeTagsRequest) error
	DeleteByAnimeId(ctx context.Context, animeId string) error
	DeleteByTagId(ctx context.Context, id string) error
	DeleteById(ctx context.Context, id string) error
}