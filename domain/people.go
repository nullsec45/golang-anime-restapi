package domain

import (
	"context"
	"database/sql"
	"github.com/nullsec45/golang-anime-restapi/dto"
)

type PeopleFilter struct {
	Search string
}

type PeopleListOptions struct {
	Pagination dto.PaginationQuery
	Filter     PeopleFilter
}

type People struct {
	Id   		string         `db:"id"`
	Slug 		string         `db:"slug"`
	NameNative  string         `db:"name_native"`
	Name        string         `db:"name"`
	Birthday    sql.NullTime   `db:"birthday"`
	Gender      dto.GenderType `db:"gender"`
	Country     string         `db:"country"`
	SiteURL     string         `db:"site_url"`
	Biography   string         `db:"biography"`
	CreatedAt   sql.NullTime   `db:"created_at"`
	UpdatedAt   sql.NullTime   `db:"updated_at"`
}

type PeopleRepository interface {
	FindAll(ctx context.Context, opts PeopleListOptions) ([]People, int64, error)
	FindById(ctx context.Context, id string) (People, error)
	FindBySlug(ctx context.Context, slug string) (People, error)
	Save(ctx context.Context, anime *People) error
	Update(ctx context.Context, anime *People) error
	Delete(ctx context.Context, id string) error
}

type PeopleService interface {
	Index(ctx context.Context, opts PeopleListOptions) (dto.Paginated[dto.PeopleData], error)
	Show(ctx context.Context, param string) (dto.PeopleData, error)
	Create(ctx context.Context, req dto.CreatePeopleRequest) error
	Update(ctx context.Context, req dto.UpdatePeopleRequest) error
	Delete(ctx context.Context, id string) error
}
