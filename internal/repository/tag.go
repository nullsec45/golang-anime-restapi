package repository

import (
	"database/sql"
	"github.com/nullsec45/golang-anime-restapi/domain"
	"github.com/doug-martin/goqu/v9"
	"context"
	"time"
)

type AnimeTagRepository struct {
	db *goqu.Database
}

func NewAnimeTag(conf *sql.DB) (domain.AnimeTagRepository){
	return &AnimeTagRepository{
		db:goqu.New("default", conf),
	}
}

func (atr *AnimeTagRepository) FindAll(ctx context.Context) (result []domain.AnimeTag, err error) {
	dataset := atr.db.From("tags").Where(goqu.C("deleted_at").IsNull())
	err = dataset.ScanStructsContext(ctx, &result)
	return
}

func (atr *AnimeTagRepository) FindById(ctx context.Context, id string) (result domain.AnimeTag, err error) {
	dataset := atr.db.From("tags").Where(
		goqu.C("id").Eq(id),
		goqu.C("deleted_at").IsNull(),
	)
	found, err := dataset.ScanStructContext(ctx, &result)
	if 	!found {
		return result, sql.ErrNoRows
	}
	return result, err
}

func (atr *AnimeTagRepository) FindBySlug(ctx context.Context, slug string) (result domain.AnimeTag, err error) {
	dataset := atr.db.From("tags").Where(
		goqu.I("tags.slug").Eq(slug),															
		goqu.I("tags.deleted_at").IsNull(),
	)

	found, scanErr := dataset.ScanStructContext(ctx, &result)
	
	if scanErr != nil {
		return result, scanErr
	}

	if !found {
		return result, sql.ErrNoRows
	}

	return result, err
}

func (atr *AnimeTagRepository) Save(ctx context.Context, anm *domain.AnimeTag) error {
	executor := atr.db.Insert("tags").Rows(anm).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (atr *AnimeTagRepository) Update(ctx context.Context, b *domain.AnimeTag) error {
    executor := atr.db.Update("tags").Set(b).Where(goqu.C("id").Eq(b.Id)).Executor()
    _, err := executor.ExecContext(ctx)
    return err
}

func (atr *AnimeTagRepository) Delete(ctx context.Context, id string) error {
	executor := atr.db.Update("tags").
	                        Set(goqu.Record{"deleted_at":sql.NullTime{Valid:true, Time:time.Now()}}).
							Where(goqu.C("id").Eq(id)).
							Executor()
	_, err := executor.ExecContext(ctx)
	return err
}