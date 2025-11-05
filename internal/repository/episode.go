package repository

import (
	"database/sql"
	"github.com/nullsec45/golang-anime-restapi/domain"
	"github.com/doug-martin/goqu/v9"
	"context"
	"strings"
	"fmt"
)

type AnimeEpisodeRepository struct {
	db *goqu.Database
}

func NewAnimeEpisode(conf *sql.DB) (domain.AnimeEpisodeRepository){
	return &AnimeEpisodeRepository{
		db:goqu.New("default", conf),
	}
}

type episodeWithTotal struct {
	domain.AnimeEpisode
	TotalCount int64 `db:"total_count"`
}

func (er *AnimeEpisodeRepository) FindAll(ctx context.Context, animeId string, opts domain.EpisodeListOptions) (items []domain.AnimeEpisode, total int64, err error) {
	q := opts.Pagination
    q.Normalize(1, 10, 100)
	limit, offset := q.LimitOffset()

	fmt.Println(opts.Pagination.Sort)
	uLimit := uint(limit)
	uOffset := uint(offset)

	dataset := er.db.From("episodes").
					Where(goqu.C("anime_id").Eq(animeId)).
					Select(
						goqu.I("episodes.id"),
						goqu.I("episodes.anime_id"),
						goqu.I("episodes.number"),
						goqu.I("episodes.season_number"),
						goqu.I("episodes.title"),
						goqu.I("episodes.synopsis"),
						goqu.I("episodes.air_date"),
						goqu.I("episodes.duration_minutes"),
						goqu.I("episodes.is_special"),
						goqu.L("COUNT(*) OVER()").As("total_count"),
					).Order(goqu.I("number").Asc())
	
    if s := opts.Filter.Search; s != "" {
        pat := "%" + s + "%"
        dataset = dataset.Where(goqu.Or(
            goqu.I("name").ILike(pat),
            goqu.I("name_native").ILike(pat),
        ))
    }

	switch opts.Pagination.Sort {
		case "number","name","created_at":
			if strings.ToLower(opts.Pagination.Order) == "asc" {
				dataset = dataset.Order(goqu.I(opts.Pagination.Sort).Asc())
			}else{
				dataset = dataset.Order(goqu.I(opts.Pagination.Sort).Desc())
			}
		default:
			dataset = dataset.Order(goqu.I("created_at").Desc())
	}

	dataset = dataset.Limit(uLimit).Offset(uOffset)

	var rows []episodeWithTotal
	if err = dataset.ScanStructsContext(ctx, &rows); err != nil {
		return nil, 0, err
	}

	if len(rows) > 0 {
		total=rows[0].TotalCount
	}

	items = make([]domain.AnimeEpisode, 0, len(rows))
	for _, r := range rows {
		items = append(items, r.AnimeEpisode)
	}

	return
}

func (er *AnimeEpisodeRepository) FindByAnimeId(ctx context.Context, animeId string) (result []domain.AnimeEpisode, err error) {
	dataset := er.db.From("episodes").Where(
		goqu.C("anime_id").Eq(animeId),
	)
	err = dataset.ScanStructsContext(ctx, &result)
	return
}

func (er *AnimeEpisodeRepository) FindByAnimeIds(ctx context.Context, animeId string) (result []domain.AnimeEpisode, err error) {
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

func (er *AnimeEpisodeRepository) Update(ctx context.Context, b *domain.AnimeEpisode) error {
    executor := er.db.Update("episodes").Set(b).Where(goqu.C("id").Eq(b.Id)).Executor()
    _, err := executor.ExecContext(ctx)
    return err
}

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