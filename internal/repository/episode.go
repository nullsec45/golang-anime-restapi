package repository

import (
	"database/sql"
	"github.com/nullsec45/golang-anime-restapi/domain"
	"github.com/doug-martin/goqu/v9"
	"context"
)

type AnimeEpisodeRepository struct {
	db *goqu.Database
}

func NewAnimeEpisode(conf *sql.DB) (domain.AnimeEpisodeRepository){
	return &AnimeEpisodeRepository{
		db:goqu.New("default", conf),
	}
}

func (er *AnimeEpisodeRepository) FindByAnimeId(ctx context.Context, animeId string) (result []domain.AnimeEpisode, err error) {
	dataset := er.db.From("episodes").Where(
		goqu.C("anime_id").Eq(animeId),
	)
	err = dataset.ScanStructsContext(ctx, &result)
	return
}

func (er *AnimeEpisodeRepository) FindById(ctx context.Context, id string) (result domain.AnimeEpisode, err error) {
	dataset := er.db.From("episodes").Where(
		goqu.C("id").Eq(id),
	)
	found, err := dataset.ScanStructContext(ctx, &result)
	if 	!found {
		return result, sql.ErrNoRows
	}
	return result, err
}

func (er *AnimeEpisodeRepository) Save (ctx context.Context, anm *domain.AnimeEpisode) error {
	executor := er.db.Insert("episodes").Rows(anm).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

// func (er *AnimeEpisodeRepository) Update(ctx context.Context, b *domain.AnimeEpisode) error {
//     executor := er.db.Update("episodes").Set(b).Where(goqu.C("id").Eq(b.Id)).Executor()
//     _, err := executor.ExecContext(ctx)
//     return err
// }

func (er *AnimeEpisodeRepository) DeleteByAnimeId(ctx context.Context, animeId string) error {
	executor := er.db.Delete("episodes").
							Where(goqu.C("anime_id").Eq(animeId)).
							Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (er *AnimeEpisodeRepository) DeleteById(ctx context.Context, id string) error {
	executor := er.db.Delete("episodes").
							Where(goqu.C("id").Eq(id)).
							Executor()
	_, err := executor.ExecContext(ctx)
	return err
}