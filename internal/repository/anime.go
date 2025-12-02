package repository

import (
	"database/sql"
	"github.com/nullsec45/golang-anime-restapi/domain"
	"github.com/doug-martin/goqu/v9"
	"context"
	"time"
	// "fmt"
	"strings"
)

type animeWithTotal struct {
	domain.Anime
	TotalCount int64 `db:"total_count"`
}

type AnimeRepository struct {
	db *goqu.Database
}

func NewAnime(conf *sql.DB) (domain.AnimeRepository){
	return &AnimeRepository{
		db:goqu.New("default", conf),
	}
}

func (ar *AnimeRepository) FindAll(ctx context.Context, opts domain.AnimeListOptions) (items []domain.Anime, total int64, err error) {
	q := opts.Pagination
    q.Normalize(1, 10, 100)
	limit, offset := q.LimitOffset()

	uLimit := uint(limit)
	uOffset := uint(offset)

	dataset := ar.db.From("animes").
					LeftJoin(
						goqu.T("media").As("m"),
						goqu.On(goqu.I("animes.cover_id").Eq(goqu.I("m.id"))),
					).
					Where(goqu.C("deleted_at").IsNull()).
					Select(
						goqu.I("animes.id"),
						goqu.I("animes.slug"),
						goqu.I("animes.title_romaji"),
						goqu.I("animes.title_native"),
						goqu.I("animes.title_english"),
						goqu.I("animes.synopsis"),
						goqu.I("animes.type"),
						goqu.I("animes.season"),
						goqu.I("animes.season_year"),
						goqu.I("animes.status"),
						goqu.I("animes.age_rating"),
						goqu.I("animes.total_episodes"),
						goqu.I("animes.average_duration_minutes"),
						goqu.I("animes.country"),
						goqu.I("animes.premiered_at"),
						goqu.I("animes.ended_at"),
						goqu.I("animes.popularity"),
						goqu.I("animes.score_avg"),
						goqu.I("animes.alt_titles"),
						goqu.I("animes.external_ids"),
						goqu.I("animes.cover_id"),
						goqu.I("animes.created_at"),
						goqu.I("animes.updated_at"),
						goqu.L("COUNT(*) OVER()").As("total_count"),
						goqu.L("m.path").As("cover_id"),
					)
	
    if s := opts.Filter.Search; s != "" {
        pat := "%" + s + "%"
        dataset = dataset.Where(goqu.Or(
            goqu.I("title_romaji").ILike(pat),
            goqu.I("title_english").ILike(pat),
            goqu.I("title_native").ILike(pat),
        ))
    }

	switch opts.Pagination.Sort {
		case "title_romaji", "created_at", "popularity", "score_avg", "season_year":
			if strings.ToLower(opts.Pagination.Order) == "asc" {
				dataset = dataset.Order(goqu.I(opts.Pagination.Sort).Asc())
			}else{
				dataset = dataset.Order(goqu.I(opts.Pagination.Sort).Desc())
			}
		default:
			dataset = dataset.Order(goqu.I("created_at").Desc())
	}

	dataset = dataset.Limit(uLimit).Offset(uOffset)

	var rows []animeWithTotal
	if err = dataset.ScanStructsContext(ctx, &rows); err != nil {
		return nil, 0, err
	}

	if len(rows) > 0 {
		total=rows[0].TotalCount
	}

	items = make([]domain.Anime, 0, len(rows))
	for _, r := range rows {
		items = append(items, r.Anime)
	}

	// err = dataset.ScanStructsContext(ctx, &result)
	return
}

func (ar *AnimeRepository) FindById(ctx context.Context, id string) (result domain.Anime, err error) {
	dataset := ar.db.Select(
						goqu.I("animes.id"),
						goqu.I("animes.slug"),
						goqu.I("animes.title_romaji"),
						goqu.I("animes.title_native"),
						goqu.I("animes.title_english"),
						goqu.I("animes.synopsis"),
						goqu.I("animes.type"),
						goqu.I("animes.season"),
						goqu.I("animes.season_year"),
						goqu.I("animes.status"),
						goqu.I("animes.age_rating"),
						goqu.I("animes.total_episodes"),
						goqu.I("animes.average_duration_minutes"),
						goqu.I("animes.country"),
						goqu.I("animes.premiered_at"),
						goqu.I("animes.ended_at"),
						goqu.I("animes.popularity"),
						goqu.I("animes.score_avg"),
						goqu.I("animes.alt_titles"),
						goqu.I("animes.external_ids"),
						goqu.I("animes.cover_id"),
						goqu.I("animes.created_at"),
						goqu.I("animes.updated_at"),
						goqu.L("m.path").As("cover_id"),
					).From("animes").
					LeftJoin(
						goqu.T("media").As("m"),
						goqu.On(goqu.I("animes.cover_id").Eq(goqu.I("m.id"))),
					).Where(
						goqu.I("animes.id").Eq(id),															
						goqu.I("animes.deleted_at").IsNull(),
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

func (ar *AnimeRepository) FindBySlug(ctx context.Context, slug string) (result domain.Anime, err error) {
	dataset := ar.db.Select(
						goqu.I("animes.id"),
						goqu.I("animes.slug"),
						goqu.I("animes.title_romaji"),
						goqu.I("animes.title_native"),
						goqu.I("animes.title_english"),
						goqu.I("animes.synopsis"),
						goqu.I("animes.type"),
						goqu.I("animes.season"),
						goqu.I("animes.season_year"),
						goqu.I("animes.status"),
						goqu.I("animes.age_rating"),
						goqu.I("animes.total_episodes"),
						goqu.I("animes.average_duration_minutes"),
						goqu.I("animes.country"),
						goqu.I("animes.premiered_at"),
						goqu.I("animes.ended_at"),
						goqu.I("animes.popularity"),
						goqu.I("animes.score_avg"),
						goqu.I("animes.alt_titles"),
						goqu.I("animes.external_ids"),
						goqu.I("animes.cover_id"),
						goqu.I("animes.created_at"),
						goqu.I("animes.updated_at"),
						goqu.L("m.path").As("cover_id"),
					).
					From("animes").
					LeftJoin(
						goqu.T("media").As("m"),
						goqu.On(goqu.I("animes.cover_id").Eq(goqu.I("m.id"))),
					).
					Where(
						goqu.I("animes.slug").Eq(slug),															
						goqu.I("animes.deleted_at").IsNull(),
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

func (ar *AnimeRepository) Save(ctx context.Context, anm *domain.Anime) error {
	executor := ar.db.Insert("animes").Rows(anm).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (ar *AnimeRepository) Update(ctx context.Context, b *domain.Anime) error {
    executor := ar.db.Update("animes").Set(b).Where(goqu.C("id").Eq(b.Id)).Executor()
    _, err := executor.ExecContext(ctx)
    return err
}

func (ar *AnimeRepository) Delete(ctx context.Context, id string) error {
	executor := ar.db.Update("animes").
	                        Set(goqu.Record{"deleted_at":sql.NullTime{Valid:true, Time:time.Now()}}).
							Where(goqu.C("id").Eq(id)).
							Executor()
	_, err := executor.ExecContext(ctx)
	return err
}