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

type peopleWithTotal struct {
	domain.People
	TotalCount int64 `db:"total_count"`
}

type PeopleRepository struct {
	db *goqu.Database
}

func NewPeople(conf *sql.DB) (domain.PeopleRepository){
	return &PeopleRepository{
		db:goqu.New("default", conf),
	}
}

func (ar *PeopleRepository) FindAll(ctx context.Context, opts domain.PeopleListOptions) (items []domain.People, total int64, err error) {
	q := opts.Pagination
    q.Normalize(1, 10, 100)
	limit, offset := q.LimitOffset()

	uLimit := uint(limit)
	uOffset := uint(offset)

	dataset := ar.db.From("peoples").
					Where(goqu.C("deleted_at").IsNull()).
					Select(
						goqu.I("peoples.id"),
						goqu.I("peoples.slug"),
						goqu.I("peoples.name_native"),
						goqu.I("peoples.name"),
						goqu.I("peoples.birthday"),
						goqu.I("peoples.gender"),
						goqu.I("peoples.country"),
						goqu.I("peoples.site_url"),
						goqu.I("peoples.biography"),
						goqu.I("peoples.created_at"),
						goqu.I("peoples.updated_at"),
						goqu.L("COUNT(*) OVER()").As("total_count"),
					)
	
    if s := opts.Filter.Search; s != "" {
        pat := "%" + s + "%"
        dataset = dataset.Where(goqu.Or(
            goqu.I("name").ILike(pat),
            goqu.I("name_native").ILike(pat),
            goqu.I("country").ILike(pat),
        ))
    }

	switch opts.Pagination.Sort {
		case "birthdary", "created_at":
			if strings.ToLower(opts.Pagination.Order) == "asc" {
				dataset = dataset.Order(goqu.I(opts.Pagination.Sort).Asc())
			}else{
				dataset = dataset.Order(goqu.I(opts.Pagination.Sort).Desc())
			}
		default:
			dataset = dataset.Order(goqu.I("created_at").Desc())
	}

	dataset = dataset.Limit(uLimit).Offset(uOffset)

	var rows []peopleWithTotal
	if err = dataset.ScanStructsContext(ctx, &rows); err != nil {
		return nil, 0, err
	}

	if len(rows) > 0 {
		total=rows[0].TotalCount
	}

	items = make([]domain.People, 0, len(rows))
	for _, r := range rows {
		items = append(items, r.People)
	}

	return
}

func (ar *PeopleRepository) FindById(ctx context.Context, id string) (result domain.People, err error) {
	dataset := ar.db.From("peoples").Where(
		goqu.I("peoples.id").Eq(id),															
		goqu.I("peoples.deleted_at").IsNull(),
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

func (ar *PeopleRepository) FindBySlug(ctx context.Context, slug string) (result domain.People, err error) {
	dataset := ar.db.From("peoples").Where(
		goqu.I("peoples.slug").Eq(slug),															
		goqu.I("peoples.deleted_at").IsNull(),
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

func (ar *PeopleRepository) Save(ctx context.Context, anm *domain.People) error {
	executor := ar.db.Insert("peoples").Rows(anm).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (ar *PeopleRepository) Update(ctx context.Context, b *domain.People) error {
    executor := ar.db.Update("peoples").Set(b).Where(goqu.C("id").Eq(b.Id)).Executor()
    _, err := executor.ExecContext(ctx)
    return err
}

func (ar *PeopleRepository) Delete(ctx context.Context, id string) error {
	executor := ar.db.Update("peoples").
	                        Set(goqu.Record{"deleted_at":sql.NullTime{Valid:true, Time:time.Now()}}).
							Where(goqu.C("id").Eq(id)).
							Executor()
	_, err := executor.ExecContext(ctx)
	return err
}