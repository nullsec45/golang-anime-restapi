package repository

import (
	"database/sql"
	"github.com/nullsec45/golang-anime-restapi/domain"
	"github.com/doug-martin/goqu/v9"
	"context"
	"time"
	"fmt"
)

type AnimeRepository struct {
	db *goqu.Database
}

func NewAnime(conf *sql.DB) (domain.AnimeRepository){
	return &AnimeRepository{
		db:goqu.New("default", conf),
	}
}

func (ar *AnimeRepository) FindAll(ctx context.Context) (result []domain.Anime, err error) {
	dataset := ar.db.From("animes").Where(goqu.C("deleted_at").IsNull())
	err = dataset.ScanStructsContext(ctx, &result)
	return
}

func (ar *AnimeRepository) FindById(ctx context.Context, id string) (result domain.Anime, err error) {
	dataset := ar.db.From("animes").Where(
		goqu.I("animes.id").Eq(id),															
		goqu.I("animes.deleted_at").IsNull(),
	)

	found, scanErr := dataset.ScanStructContext(ctx, &result)

	sqlStr, args, _ := dataset.ToSQL()
    fmt.Println(sqlStr, args)
	
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