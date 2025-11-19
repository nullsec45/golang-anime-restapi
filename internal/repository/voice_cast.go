package repository 

import (
	"database/sql"
	"github.com/nullsec45/golang-anime-restapi/domain"
	"github.com/doug-martin/goqu/v9"
	"context"
)

type VoiceCastRepository struct {
	db *goqu.Database
}

func NewVoiceCast(conf *sql.DB) (domain.VoiceCastRepository){
	return &VoiceCastRepository{
		db:goqu.New("default", conf),
	}
}

func (vcr *VoiceCastRepository) FindById(ctx context.Context, id string) (result domain.VoiceCast, err error) {
	dataset := vcr.db.From("voice_casts").Where(
		goqu.C("id").Eq(id),
	)
	found, err := dataset.ScanStructContext(ctx, &result)
	if 	!found {
		return result, sql.ErrNoRows
	}
	return result, err
}

func (vcr *VoiceCastRepository) FindUnique(ctx context.Context, animeId string, characterId string, personId string) (result domain.VoiceCast, found bool, err error) {
	dataset := vcr.db.From("voice_casts").Where(
		goqu.C("anime_id").Eq(animeId),
		goqu.C("character_id").Eq(characterId),
		goqu.C("person_id").Eq(personId),
	)
	found, err = dataset.ScanStructContext(ctx, &result)
	return
}

func (vcr *VoiceCastRepository) Save(ctx context.Context, vcs *domain.VoiceCast) error {
	executor := vcr.db.Insert("voice_casts").Rows(vcs).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (vcr *VoiceCastRepository) Update(ctx context.Context, vcs *domain.VoiceCast) error {
    executor := vcr.db.Update("voice_casts").Set(vcs).Where(goqu.C("id").Eq(vcs.Id)).Executor()
    _, err := executor.ExecContext(ctx)
    return err
}

// func (vcr *VoiceCastRepository) DeleteByAnimeId(ctx context.Context, animeId string) error {
// 	executor := vcr.db.Delete("voice_casts").
// 							Where(goqu.C("anime_id").Eq(animeId)).
// 							Executor()
// 	_, err := executor.ExecContext(ctx)
// 	return err
// }

// func (vcr *VoiceCastRepository) DeleteByTagId(ctx context.Context, tagId string) error {
// 	executor := vcr.db.Delete("voice_casts").
// 							Where(goqu.C("tag_id").Eq(tagId)).
// 							Executor()
// 	_, err := executor.ExecContext(ctx)
// 	return err
// }

func (vcr *VoiceCastRepository) DeleteById(ctx context.Context, id string) error {
	executor := vcr.db.Delete("voice_casts").
							Where(goqu.C("id").Eq(id)).
							Executor()
	_, err := executor.ExecContext(ctx)
	return err
}