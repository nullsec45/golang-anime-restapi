package repository

import (
	"database/sql"
	"github.com/nullsec45/golang-anime-restapi/domain"
	"github.com/doug-martin/goqu/v9"
	"context"
	"strings"
	"time"
)

type CharacterRepository struct {
	db *goqu.Database
}

type characterWithTotal struct {
	domain.Character
	TotalCount int64 `db:"total_count"`
}


func NewCharacter(conf *sql.DB) (domain.CharacterRepository){
	return &CharacterRepository{
		db:goqu.New("default", conf),
	}
}


func (cr *CharacterRepository) FindAll(ctx context.Context, opts domain.CharacterListOptions) (items []domain.Character, total int64, err error) {
	q := opts.Pagination
    q.Normalize(1, 10, 100)
	limit, offset := q.LimitOffset()

	uLimit := uint(limit)
	uOffset := uint(offset)

	dataset := cr.db.From("characters").
					Where(goqu.C("deleted_at").IsNull()).
					Select(
						goqu.I("characters.id"),
						goqu.I("characters.slug"),
						goqu.I("characters.name"),
						goqu.I("characters.name_native"),
						goqu.I("characters.description"),
						goqu.L("COUNT(*) OVER()").As("total_count"),
					)
	
    if s := opts.Filter.Search; s != "" {
        pat := "%" + s + "%"
        dataset = dataset.Where(goqu.Or(
            goqu.I("name").ILike(pat),
            goqu.I("name_native").ILike(pat),
        ))
    }

	switch opts.Pagination.Sort {
		case "name", "created_at", "name_native":
			if strings.ToLower(opts.Pagination.Order) == "asc" {
				dataset = dataset.Order(goqu.I(opts.Pagination.Sort).Asc())
			}else{
				dataset = dataset.Order(goqu.I(opts.Pagination.Sort).Desc())
			}
		default:
			dataset = dataset.Order(goqu.I("created_at").Desc())
	}

	dataset = dataset.Limit(uLimit).Offset(uOffset)

	var rows []characterWithTotal
	if err = dataset.ScanStructsContext(ctx, &rows); err != nil {
		return nil, 0, err
	}

	if len(rows) > 0 {
		total=rows[0].TotalCount
	}

	items = make([]domain.Character, 0, len(rows))
	for _, r := range rows {
		items = append(items, r.Character)
	}

	return
}

func (cr *CharacterRepository) FindById(ctx context.Context, id string) (result domain.Character, err error) {
	dataset := cr.db.From("characters").Where(
		goqu.C("id").Eq(id),
	)
	found, err := dataset.ScanStructContext(ctx, &result)
	if 	!found {
		return result, sql.ErrNoRows
	}
	return result, err
}


func (cr *CharacterRepository) FindBySlug(ctx context.Context, slug string) (result domain.Character, err error) {
	dataset := cr.db.From("characters").Where(
		goqu.I("characters.slug").Eq(slug),															
		goqu.I("characters.deleted_at").IsNull(),
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

func (cr *CharacterRepository) Save (ctx context.Context, anm *domain.Character) error {
	executor := cr.db.Insert("characters").Rows(anm).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (cr *CharacterRepository) Update(ctx context.Context, b *domain.Character) error {
    executor := cr.db.Update("characters").Set(b).Where(goqu.C("id").Eq(b.Id)).Executor()
    _, err := executor.ExecContext(ctx)
    return err
}

func (cr *CharacterRepository) Delete(ctx context.Context, id string) error {
	executor := cr.db.Update("characters").
	                        Set(goqu.Record{"deleted_at":sql.NullTime{Valid:true, Time:time.Now()}}).
							Where(goqu.C("id").Eq(id)).
							Executor()
	_, err := executor.ExecContext(ctx)
	return err
}