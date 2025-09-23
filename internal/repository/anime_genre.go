package repository 

import (
	"database/sql"
	"github.com/nullsec45/golang-anime-restapi/domain"
	"github.com/doug-martin/goqu/v9"
	"context"
)

type AnimeGenresRepository struct {
	db *goqu.Database
}

func NewAnimeGenres(conf *sql.DB) (domain.AnimeGenresRepository){
	return &AnimeGenresRepository{
		db:goqu.New("default", conf),
	}
}

func (agrs *AnimeGenresRepository) FindById(ctx context.Context, id string) (result domain.AnimeGenres, err error) {
	dataset := agrs.db.From("anime_genres").Where(
		goqu.C("id").Eq(id),
	)
	found, err := dataset.ScanStructContext(ctx, &result)
	if 	!found {
		return result, sql.ErrNoRows
	}
	return result, err
}

func (agrs *AnimeGenresRepository) FindByAnimeAndGenreId(ctx context.Context, animeId string, genreId string) (result domain.AnimeGenres, found bool, err error) {
	dataset := agrs.db.From("anime_genres").Where(
		goqu.C("anime_id").Eq(animeId),
		goqu.C("genre_id").Eq(genreId),
	)
	found, err = dataset.ScanStructContext(ctx, &result)
	return
}



func (agrs *AnimeGenresRepository) Save(ctx context.Context, anmgrs *domain.AnimeGenres) error {
	executor := agrs.db.Insert("anime_genres").Rows(anmgrs).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (agrs *AnimeGenresRepository) Update(ctx context.Context, anmgrs *domain.AnimeGenres) error {
    executor := agrs.db.Update("anime_genres").Set(anmgrs).Where(goqu.C("id").Eq(anmgrs.Id)).Executor()
    _, err := executor.ExecContext(ctx)
    return err
}

func (agrs *AnimeGenresRepository) DeleteByAnimeId(ctx context.Context, animeId string) error {
	executor := agrs.db.Delete("anime_genres").
							Where(goqu.C("anime_id").Eq(animeId)).
							Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (agrs *AnimeGenresRepository) DeleteByGenreId(ctx context.Context, genreId string) error {
	executor := agrs.db.Delete("anime_genres").
							Where(goqu.C("genre_id").Eq(genreId)).
							Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (agrs *AnimeGenresRepository) DeleteById(ctx context.Context, id string) error {
	executor := agrs.db.Delete("anime_genres").
							Where(goqu.C("id").Eq(id)).
							Executor()
	_, err := executor.ExecContext(ctx)
	return err
}