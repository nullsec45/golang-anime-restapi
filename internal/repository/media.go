package repository

import (
	"context"
	"database/sql"
	"github.com/doug-martin/goqu/v9"
	"github.com/nullsec45/golang-anime-restapi/domain"
)

type MediaRepository struct {
	db *goqu.Database
}

func NewMedia(con *sql.DB) domain.MediaRepository {
	return &MediaRepository{
		db:goqu.New("default", con),
	}
}

func (m MediaRepository) FindById(ctx context.Context, id string) (media domain.Media, err error) {
	dataset := m.db.From("media").Where(goqu.Ex{
		"id":id,
	})

	found, err := dataset.ScanStructContext(ctx, &media)
	if 	!found {
		return media, sql.ErrNoRows
	}
	
	return media, err

	// _, err = dataset.ScanStructContext(ctx, &media)
	// return
}

func (m MediaRepository) FindByIds(ctx context.Context, ids []string) (medias []domain.Media, err error){
	dataset := m.db.From("media").Where(goqu.C("id").In(ids))
	err = dataset.ScanStructsContext(ctx, &medias)
	return
}

func (m MediaRepository) Save(ctx context.Context, media *domain.Media) error {
	executor := m.db.Insert("media").Rows(media).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

// func (m *MediaRepository) Update(ctx context.Context, media *domain.Media) error {
//     executor := m.db.Update("media").Set(media).Where(goqu.C("id").Eq(media.Id)).Executor()
//     _, err := executor.ExecContext(ctx)
//     return err
// }

func (m MediaRepository) Delete(ctx context.Context, id string) error {
	executor := m.db.Delete("media").
							Where(goqu.C("id").Eq(id)).
							Executor()
	_, err := executor.ExecContext(ctx)
	return err
}