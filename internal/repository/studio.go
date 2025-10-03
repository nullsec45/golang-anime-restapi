package repository

import (
	"database/sql"
	"github.com/nullsec45/golang-anime-restapi/domain"
	"github.com/doug-martin/goqu/v9"
	"context"
	"time"
)

type AnimeStudioRepository struct {
	db *goqu.Database
}

func NewAnimeStudio(conf *sql.DB) (domain.AnimeStudioRepository){
	return &AnimeStudioRepository{
		db:goqu.New("default", conf),
	}
}

func (asr *AnimeStudioRepository) FindAll(ctx context.Context) (result []domain.AnimeStudio, err error) {
	dataset := asr.db.From("studios").Where(goqu.C("deleted_at").IsNull())
	err = dataset.ScanStructsContext(ctx, &result)
	return
}

func (asr *AnimeStudioRepository) FindById(ctx context.Context, id string) (result domain.AnimeStudio, err error) {
	dataset := asr.db.From("studios").Where(
		goqu.C("id").Eq(id),
		goqu.C("deleted_at").IsNull(),
	)
	found, err := dataset.ScanStructContext(ctx, &result)
	if 	!found {
		return result, sql.ErrNoRows
	}
	return result, err
}

func (asr *AnimeStudioRepository) FindByAnimeId(ctx context.Context, animeId string) ([]domain.AnimeStudio, error) {
	ds := asr.db.
		From(goqu.T("studios")).
		InnerJoin(
			goqu.T("anime_studios"),
			goqu.On(goqu.I("studios.id").Eq(goqu.I("anime_studios.studio_id"))),
		).
		Select(goqu.I("studios.id"), goqu.I("studios.slug"), goqu.I("studios.name")). 
		Where(goqu.I("anime_studios.anime_id").Eq(animeId)).
		Order(goqu.I("studios.name").Asc())
	var rows []domain.AnimeStudio
	err := ds.ScanStructsContext(ctx, &rows)
	return rows, err
}

func (asr *AnimeStudioRepository) Save(ctx context.Context, anms *domain.AnimeStudio) error {
	executor := asr.db.Insert("studios").Rows(anms).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (asr *AnimeStudioRepository) Update(ctx context.Context, anms *domain.AnimeStudio) error {
    executor := asr.db.Update("studios").Set(anms).Where(goqu.C("id").Eq(anms.Id)).Executor()
    _, err := executor.ExecContext(ctx)
    return err
}

func (asr *AnimeStudioRepository) Delete(ctx context.Context, id string) error {
	executor := asr.db.Update("studios").
	                        Set(goqu.Record{"deleted_at":sql.NullTime{Valid:true, Time:time.Now()}}).
							Where(goqu.C("id").Eq(id)).
							Executor()
	_, err := executor.ExecContext(ctx)
	return err
}