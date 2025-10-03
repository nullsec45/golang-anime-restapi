package repository 

import (
	"database/sql"
	"github.com/nullsec45/golang-anime-restapi/domain"
	"github.com/doug-martin/goqu/v9"
	"context"
)

type AnimeStudiosRepository struct {
	db *goqu.Database
}

func NewAnimeStudios(conf *sql.DB) (domain.AnimeStudiosRepository){
	return &AnimeStudiosRepository{
		db:goqu.New("default", conf),
	}
}

func (astd *AnimeStudiosRepository) FindById(ctx context.Context, id string) (result domain.AnimeStudios, err error) {
	dataset := astd.db.From("anime_studios").Where(
		goqu.C("id").Eq(id),
	)
	found, err := dataset.ScanStructContext(ctx, &result)
	if 	!found {
		return result, sql.ErrNoRows
	}
	return result, err
}

func (astd *AnimeStudiosRepository) FindByAnimeAndStudioId(ctx context.Context, animeId string, studioId string) (result domain.AnimeStudios, found bool, err error) {
	dataset := astd.db.From("anime_studios").Where(
		goqu.C("anime_id").Eq(animeId),
		goqu.C("studio_id").Eq(studioId),
	)
	found, err = dataset.ScanStructContext(ctx, &result)
	return
}



func (astd *AnimeStudiosRepository) Save(ctx context.Context, anmstd *domain.AnimeStudios) error {
	executor := astd.db.Insert("anime_studios").Rows(anmstd).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (astd *AnimeStudiosRepository) Update(ctx context.Context, anmstd *domain.AnimeStudios) error {
    executor := astd.db.Update("anime_studios").Set(anmstd).Where(goqu.C("id").Eq(anmstd.Id)).Executor()
    _, err := executor.ExecContext(ctx)
    return err
}

func (astd *AnimeStudiosRepository) DeleteByAnimeId(ctx context.Context, animeId string) error {
	executor := astd.db.Delete("anime_studios").
							Where(goqu.C("anime_id").Eq(animeId)).
							Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (astd *AnimeStudiosRepository) DeleteByStudioId(ctx context.Context, studioId string) error {
	executor := astd.db.Delete("anime_studios").
							Where(goqu.C("studio_id").Eq(studioId)).
							Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (astd *AnimeStudiosRepository) DeleteById(ctx context.Context, id string) error {
	executor := astd.db.Delete("anime_studios").
							Where(goqu.C("id").Eq(id)).
							Executor()
	_, err := executor.ExecContext(ctx)
	return err
}