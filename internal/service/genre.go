package service

import(
	"github.com/nullsec45/golang-anime-restapi/domain"
	"github.com/nullsec45/golang-anime-restapi/dto"
	"github.com/nullsec45/golang-anime-restapi/internal/utility"
	"context"
	"github.com/google/uuid"
	"database/sql"
	"time"
	"github.com/gosimple/slug"
)

type AnimeGenreService struct {
	animeGenreRepository domain.AnimeGenreRepository
}

func NewAnimeGenre(
	animeGenreRepository domain.AnimeGenreRepository,
) domain.AnimeGenreService {
	return &AnimeGenreService{
		animeGenreRepository: animeGenreRepository,
	}
}

func (ags AnimeGenreService) Index(ctx context.Context) ([]dto.AnimeGenreData, error) {
	genres, err := ags.animeGenreRepository.FindAll(ctx)

	if err != nil {
		return nil, err
	}

	var animeGenreData []dto.AnimeGenreData

	for _, v:= range genres {
		animeGenreData = append(animeGenreData, dto.AnimeGenreData{
			Id: v.Id,
			Slug: v.Slug,
			Name: v.Name,
		})
	}

	return animeGenreData, nil
}

func (ags AnimeGenreService) Show (ctx context.Context, param string) (dto.AnimeGenreData, error) {
	exist, err := func() (domain.AnimeGenre, error) {
		if utility.IsUUID(param) {
			return ags.animeGenreRepository.FindById(ctx, param)
		}
		return ags.animeGenreRepository.FindBySlug(ctx, param)
	}()


    if err != nil && exist.Id == "" {
        return dto.AnimeGenreData{}, domain.AnimeGenreNotFound
    }
    
    if err != nil {
        return dto.AnimeGenreData{}, err
    }

    return dto.AnimeGenreData{
		Id:exist.Id,
		Slug:exist.Slug,
		Name:exist.Name,
	}, nil
}

func (ags AnimeGenreService) Create(ctx context.Context, req dto.CreateAnimeGenreRequest) error {
	animeSlug := req.Slug

    if animeSlug == "" {
        animeSlug = slug.Make(req.Name) 
    }
	
 	anime := domain.AnimeGenre{
        Id: uuid.New().String(),
		Slug:animeSlug,
		Name:req.Name,
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
    }

	return ags.animeGenreRepository.Save(ctx, &anime)
}

func (ags AnimeGenreService) Update(ctx context.Context, req dto.UpdateAnimeGenreRequest)  error {
    // Cari data anime
    exist, err := ags.animeGenreRepository.FindById(ctx, req.Id)

    if err != nil && exist.Id == "" {
        return domain.AnimeGenreNotFound
    }
    
    if err != nil {
        return  err
    }

	animeSlug := req.Slug
    if animeSlug == "" {
        animeSlug = slug.Make(req.Name) 
    }

	exist.Slug = animeSlug	
	exist.Name=req.Name
	exist.UpdatedAt = sql.NullTime{Time: time.Now(), Valid: true}

	return ags.animeGenreRepository.Update(ctx, &exist)

}

func (ags AnimeGenreService) Delete (ctx context.Context, id string) error {
    exist, err := ags.animeGenreRepository.FindById(ctx, id)

    if err != nil && exist.Id == "" {
        return  domain.AnimeGenreNotFound
    }
    
    if err != nil {
        return err
    }

    return ags.animeGenreRepository.Delete(ctx, id)
} 