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

func (atgs *AnimeTagsRepository) FindByAnimeId(ctx context.Context, animeId string) ([]domain.AnimeTag, error) {
	ds := atgs.db.
		From(goqu.T("anime_tags")).
		InnerJoin(
			goqu.T("tags"),
			goqu.On(goqu.I("anime_tags.tag_id").Eq(goqu.I("tags.id"))),
		).
		Select(goqu.I("tags.id"), goqu.I("tags.slug"), goqu.I("tags.name")). 
		Where(goqu.I("anime_tags.anime_id").Eq(animeId)).
		Order(goqu.I("tags.name").Asc())
	var rows []domain.AnimeTag
	err := ds.ScanStructsContext(ctx, &rows)
	return rows, err
}

func(atgs *AnimeTagsRepository) FindByAnimeIDs(ctx context.Context, animeIDs []string)(map[string][]domain.AnimeTag, error){
	if len(animeIDs) == 0 { 
		return nil, nil
	}

	subquery := goqu.From("anime_tags").
        Select(
            goqu.I("anime_tags.anime_id"),
            goqu.I("anime_tags.tag_id"),
            goqu.L("ROW_NUMBER() OVER (PARTITION BY anime_tags.anime_id ORDER BY anime_tags.created_at DESC)").As("rn"),
        ).Where(goqu.I("anime_tags.anime_id").In(animeIDs))

	dataset := atgs.db.From(subquery.As("sq")).
        Select(
            goqu.I("sq.anime_id"),
            goqu.I("g.id").As("id"),
            goqu.I("g.slug").As("slug"),
            goqu.I("g.name").As("name"),
        ).
        LeftJoin(
            goqu.T("genres").As("g"),
            goqu.On(goqu.I("sq.tag_id").Eq(goqu.I("g.id"))),
        ).
        Where(goqu.I("sq.rn").Lte(3)).
        Order(goqu.I("sq.anime_id").Asc(), goqu.I("sq.rn").Asc())

	type GenreRow struct {
		AnimeID string `db:"anime_id"`
		domain.AnimeTag
	}

	var rows []GenreRow
	if err := dataset.ScanStructsContext(ctx, &rows); err != nil {
		return nil, err
	}

	genresMap := make(map[string][]domain.AnimeTag)
	for _, row := range rows {
		genresMap[row.AnimeID] = append(genresMap[row.AnimeID], row.AnimeTag)
	}

	return genresMap, nil
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