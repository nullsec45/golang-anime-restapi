package repository

import (
	"database/sql"
	"github.com/nullsec45/golang-anime-restapi/domain"
	"github.com/doug-martin/goqu/v9"
	"context"
	"time"
)

type AnimeGenreRepository struct {
	db *goqu.Database
}

func NewAnimeGenre(conf *sql.DB) (domain.AnimeGenreRepository){
	return &AnimeGenreRepository{
		db:goqu.New("default", conf),
	}
}

func (agr *AnimeGenreRepository) FindAll(ctx context.Context) (result []domain.AnimeGenre, err error) {
	dataset := agr.db.From("genres").Where(goqu.C("deleted_at").IsNull())
	err = dataset.ScanStructsContext(ctx, &result)
	return
}

func (agr *AnimeGenreRepository) FindById(ctx context.Context, id string) (result domain.AnimeGenre, err error) {
	dataset := agr.db.From("genres").Where(
		goqu.C("id").Eq(id),
		goqu.C("deleted_at").IsNull(),
	)
	found, err := dataset.ScanStructContext(ctx, &result)
	if 	!found {
		return result, sql.ErrNoRows
	}
	return result, err
}

func (agr *AnimeGenreRepository) FindBySlug(ctx context.Context, slug string) (result domain.AnimeGenre, err error) {
	dataset := agr.db.From("genres").Where(
		goqu.I("genres.slug").Eq(slug),															
		goqu.I("genres.deleted_at").IsNull(),
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

func (agr *AnimeGenreRepository) FindByAnimeId(ctx context.Context, animeId string) ([]domain.AnimeGenre, error) {
	ds := agr.db.
		From(goqu.T("genres")).
		InnerJoin(
			goqu.T("anime_genres"),
			goqu.On(goqu.I("genres.id").Eq(goqu.I("anime_genres.genre_id"))),
		).
		Select(goqu.I("genres.id"), goqu.I("genres.slug"), goqu.I("genres.name")). 
		Where(goqu.I("anime_genres.anime_id").Eq(animeId)).
		Order(goqu.I("genres.name").Asc())
	var rows []domain.AnimeGenre
	err := ds.ScanStructsContext(ctx, &rows)
	return rows, err
}

func (agr *AnimeGenreRepository) Save(ctx context.Context, anm *domain.AnimeGenre) error {
	executor := agr.db.Insert("genres").Rows(anm).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (agr *AnimeGenreRepository) Update(ctx context.Context, b *domain.AnimeGenre) error {
    executor := agr.db.Update("genres").Set(b).Where(goqu.C("id").Eq(b.Id)).Executor()
    _, err := executor.ExecContext(ctx)
    return err
}

func (agr *AnimeGenreRepository) Delete(ctx context.Context, id string) error {
	executor := agr.db.Update("genres").
	                        Set(goqu.Record{"deleted_at":sql.NullTime{Valid:true, Time:time.Now()}}).
							Where(goqu.C("id").Eq(id)).
							Executor()
	_, err := executor.ExecContext(ctx)
	return err
}