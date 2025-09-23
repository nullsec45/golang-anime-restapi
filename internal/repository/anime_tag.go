package repository 

import (
	"database/sql"
	"github.com/nullsec45/golang-anime-restapi/domain"
	"github.com/doug-martin/goqu/v9"
	"context"
)

type AnimeTagsRepository struct {
	db *goqu.Database
}

func NewAnimeTags(conf *sql.DB) (domain.AnimeTagsRepository){
	return &AnimeTagsRepository{
		db:goqu.New("default", conf),
	}
}

func (atgs *AnimeTagsRepository) FindById(ctx context.Context, id string) (result domain.AnimeTags, err error) {
	dataset := atgs.db.From("anime_tags").Where(
		goqu.C("id").Eq(id),
	)
	found, err := dataset.ScanStructContext(ctx, &result)
	if 	!found {
		return result, sql.ErrNoRows
	}
	return result, err
}

func (atgs *AnimeTagsRepository) FindByAnimeAndTagId(ctx context.Context, animeId string, tagId string) (result domain.AnimeTags, found bool, err error) {
	dataset := atgs.db.From("anime_tags").Where(
		goqu.C("anime_id").Eq(animeId),
		goqu.C("tag_id").Eq(tagId),
	)
	found, err = dataset.ScanStructContext(ctx, &result)
	return
}



func (atgs *AnimeTagsRepository) Save(ctx context.Context, anmtgs *domain.AnimeTags) error {
	executor := atgs.db.Insert("anime_tags").Rows(anmtgs).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (atgs *AnimeTagsRepository) Update(ctx context.Context, anmtgs *domain.AnimeTags) error {
    executor := atgs.db.Update("anime_tags").Set(anmtgs).Where(goqu.C("id").Eq(anmtgs.Id)).Executor()
    _, err := executor.ExecContext(ctx)
    return err
}

func (atgs *AnimeTagsRepository) DeleteByAnimeId(ctx context.Context, animeId string) error {
	executor := atgs.db.Delete("anime_tags").
							Where(goqu.C("anime_id").Eq(animeId)).
							Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (atgs *AnimeTagsRepository) DeleteByTagId(ctx context.Context, tagId string) error {
	executor := atgs.db.Delete("anime_tags").
							Where(goqu.C("tag_id").Eq(tagId)).
							Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (atgs *AnimeTagsRepository) DeleteById(ctx context.Context, id string) error {
	executor := atgs.db.Delete("anime_tags").
							Where(goqu.C("id").Eq(id)).
							Executor()
	_, err := executor.ExecContext(ctx)
	return err
}