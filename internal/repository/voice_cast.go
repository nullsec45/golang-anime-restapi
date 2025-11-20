package repository 

import (
	"database/sql"
	"github.com/nullsec45/golang-anime-restapi/domain"
	"github.com/doug-martin/goqu/v9"
	"context"
	"fmt"
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

func (vcr *VoiceCastRepository) FindByAnimeId(ctx context.Context, animeId string) ([]domain.VoiceCastWithRelation,  error) {
	ds := vcr.db.
		From(goqu.T("voice_casts").As("vc")).
		LeftJoin(
			goqu.T("characters").As("c"),
			goqu.On(goqu.I("vc.character_id").Eq(goqu.I("c.id"))),
		).LeftJoin(
			goqu.T("peoples").As("p"),	
			goqu.On(goqu.I("vc.person_id").Eq(goqu.I("p.id"))),
		).Select(
			// vc
			goqu.I("vc.character_id"),
			goqu.I("vc.person_id"),

			goqu.I("vc.id"),
			goqu.I("vc.language"),
			goqu.I("vc.role_note"),

			// characters
			// goqu.I("c.id").As("character_id"),
			goqu.I("c.slug").As("character_slug"),
			goqu.I("c.name").As("character_name"),
			goqu.I("c.name_native").As("character_name_native"),
			goqu.I("c.description"),
			// contoh tambahan:
			// goqu.I("c.image_url").As("character_image_url"),

			// people
			// goqu.I("p.id").As("people_id"),
			goqu.I("p.slug").As("people_slug"),
			goqu.I("p.name").As("people_name"),
			goqu.I("p.name_native").As("people_name_native"),
			goqu.I("p.birthday"),
			goqu.I("p.gender"),
			goqu.I("p.country"),
			goqu.I("p.site_url"),
			goqu.I("p.biography"),
			// contoh tambahan:
			// goqu.I("p.image_url").As("people_image_url"),
		).
		Where(goqu.I("vc.anime_id").Eq(animeId)).
		Order(goqu.I("c.name").Asc())


	var rows []domain.VoiceCastWithRelation
	err := ds.ScanStructsContext(ctx, &rows)

	sql, args, _ := ds.ToSQL()
	fmt.Println("VOICE_CAST SQL:", sql)
	fmt.Println("ARGS:", args)
	return rows, err
}

func (vcr *VoiceCastRepository) FindUnique(ctx context.Context, animeId string, characterId string, personId string) (result domain.VoiceCast, found bool, err error) {
	dataset := vcr.db.From("voice_casts").
	Where(
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