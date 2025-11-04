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

type AnimeTagService struct {
	animeTagRepository domain.AnimeTagRepository
}

func NewAnimeTag(
	animeTagRepository domain.AnimeTagRepository,
) domain.AnimeTagService {
	return &AnimeTagService{
		animeTagRepository: animeTagRepository,
	}
}

func (ats AnimeTagService) Index(ctx context.Context) ([]dto.AnimeTagData, error) {
	genres, err := ats.animeTagRepository.FindAll(ctx)

	if err != nil {
		return nil, err
	}

	var animeTagData []dto.AnimeTagData

	for _, v:= range genres {
		animeTagData = append(animeTagData, dto.AnimeTagData{
			Id: v.Id,
			Slug: v.Slug,
			Name: v.Name,
		})
	}

	return animeTagData, nil
}

func (ats AnimeTagService) Show (ctx context.Context, param string) (dto.AnimeTagData, error) {
   exist, err := func() (domain.AnimeTag, error) {
		if utility.IsUUID(param) {
			return ats.animeTagRepository.FindById(ctx, param)
		}
		return ats.animeTagRepository.FindBySlug(ctx, param)
	}()

    if err != nil && exist.Id == "" {
        return dto.AnimeTagData{}, utility.NewNotFound("Anime Tag")
    }
    
    if err != nil {
        return dto.AnimeTagData{}, err
    }

    return dto.AnimeTagData{
		Id:exist.Id,
		Slug:exist.Slug,
		Name:exist.Name,
	}, nil
}

func (ats AnimeTagService) Create(ctx context.Context, req dto.CreateAnimeTagRequest) error {
	animeSlug := req.Slug

    if animeSlug == "" {
        animeSlug = slug.Make(req.Name) 
    }
	
 	anime := domain.AnimeTag{
        Id: uuid.New().String(),
		Slug:animeSlug,
		Name:req.Name,
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
    }

	return ats.animeTagRepository.Save(ctx, &anime)
}

func (ats AnimeTagService) Update(ctx context.Context, req dto.UpdateAnimeTagRequest)  error {
    // Cari data anime
    exist, err := ats.animeTagRepository.FindById(ctx, req.Id)

    if err != nil && exist.Id == "" {
        return utility.NewNotFound("Anime Tag")
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

	return ats.animeTagRepository.Update(ctx, &exist)

}

func (ats AnimeTagService) Delete (ctx context.Context, id string) error {
    exist, err := ats.animeTagRepository.FindById(ctx, id)

    if err != nil && exist.Id == "" {
        return  utility.NewNotFound("Anime Tag")
    }
    
    if err != nil {
        return err
    }

    return ats.animeTagRepository.Delete(ctx, id)
} 