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
func (agrs *AnimeGenresRepository) FindByAnimeId(ctx context.Context, animeId string) ([]domain.AnimeGenre, error) {
	ds := agrs.db.
		From(goqu.T("anime_genres")).
		InnerJoin(
			goqu.T("anime"),
			goqu.On(goqu.I("anime_genres.anime_id").Eq(goqu.I("animes.id"))),
		).
		Select(goqu.I("genres.id"), goqu.I("genres.slug"), goqu.I("genres.name")). 
		Where(goqu.I("anime_genres.anime_id").Eq(animeId)).
		Order(goqu.I("genres.name").Asc())
	var rows []domain.AnimeGenre
	err := ds.ScanStructsContext(ctx, &rows)
	return rows, err
}

func(agrs *AnimeGenresRepository) FindByAnimeIDs(ctx context.Context, animeIDs []string)(map[string][]domain.AnimeGenre, error){
	if len(animeIDs) == 0 { 
		return nil, nil
	}

	subquery := goqu.From("anime_genres").
        Select(
            goqu.I("anime_genres.anime_id"),
            goqu.I("anime_genres.genre_id"),
            goqu.L("ROW_NUMBER() OVER (PARTITION BY anime_genres.anime_id ORDER BY anime_genres.created_at DESC)").As("rn"),
        ).Where(goqu.I("anime_genres.anime_id").In(animeIDs))

	dataset := agrs.db.From(subquery.As("sq")).
        Select(
            goqu.I("sq.anime_id"),
            goqu.I("g.id").As("id"),
            goqu.I("g.slug").As("slug"),
            goqu.I("g.name").As("name"),
        ).
        LeftJoin(
            goqu.T("genres").As("g"),
            goqu.On(goqu.I("sq.genre_id").Eq(goqu.I("g.id"))),
        ).
        Where(goqu.I("sq.rn").Lte(3)).
        Order(goqu.I("sq.anime_id").Asc(), goqu.I("sq.rn").Asc())

	type GenreRow struct {
		AnimeID string `db:"anime_id"`
		domain.AnimeGenre
	}

	var rows []GenreRow
	if err := dataset.ScanStructsContext(ctx, &rows); err != nil {
		return nil, err
	}

	genresMap := make(map[string][]domain.AnimeGenre)
	for _, row := range rows {
		genresMap[row.AnimeID] = append(genresMap[row.AnimeID], row.AnimeGenre)
	}

	return genresMap, nil
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