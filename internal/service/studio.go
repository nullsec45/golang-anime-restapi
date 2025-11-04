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

type AnimeStudioService struct {
	animeStudioRepository domain.AnimeStudioRepository
}

func NewAnimeStudio(
	animeStudioRepository domain.AnimeStudioRepository,
) domain.AnimeStudioService {
	return &AnimeStudioService{
		animeStudioRepository: animeStudioRepository,
	}
}	

func (ass AnimeStudioService) Index(ctx context.Context) ([]dto.AnimeStudioData, error) {
	studios, err := ass.animeStudioRepository.FindAll(ctx)

	if err != nil {
		return nil, err
	}

	var animeStudioData []dto.AnimeStudioData

	for _, v:= range studios {
		animeStudioData = append(animeStudioData, dto.AnimeStudioData{
			Id: v.Id,
			Slug: v.Slug,
			Name: v.Name,
			Country: v.Country,
			SiteURL: v.SiteURL,
		})
	}

	return animeStudioData, nil
}

func (ass AnimeStudioService) Show (ctx context.Context, param string) (dto.AnimeStudioData, error) {
	exist, err := func() (domain.AnimeStudio, error) {
		if utility.IsUUID(param) {
			return ass.animeStudioRepository.FindById(ctx, param)
		}
		return ass.animeStudioRepository.FindBySlug(ctx, param)
	}()

    if err != nil && exist.Id == "" {
        return dto.AnimeStudioData{}, utility.NewNotFound("Anime Studio")
    }
    
    if err != nil {
        return dto.AnimeStudioData{}, err
    }

    return dto.AnimeStudioData{
		Id:exist.Id,
		Slug:exist.Slug,
		Name:exist.Name,
		Country: exist.Country,
		SiteURL: exist.SiteURL,
	}, nil
}

func (ass AnimeStudioService) Create(ctx context.Context, req dto.CreateAnimeStudioRequest) error {
	studioSlug := req.Slug

    if studioSlug == "" {
        studioSlug = slug.Make(req.Name) 
    }
	
 	anime := domain.AnimeStudio{
        Id: uuid.New().String(),
		Slug:studioSlug,
		Name:req.Name,
		Country: req.Country,
		SiteURL: req.SiteURL,
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
    }

	return ass.animeStudioRepository.Save(ctx, &anime)
}

func (ass AnimeStudioService) Update(ctx context.Context, req dto.UpdateAnimeStudioRequest)  error {
    // Cari data anime
    exist, err := ass.animeStudioRepository.FindById(ctx, req.Id)

    if err != nil && exist.Id == "" {
        return utility.NewNotFound("Anime Studio")
    }
    
    if err != nil {
        return  err
    }

	studioSlug := req.Slug
    if studioSlug == "" {
        studioSlug = slug.Make(req.Name) 
    }

	exist.Slug = studioSlug	
	exist.Name=req.Name
	exist.Country=req.Country
	exist.SiteURL=req.SiteURL
	exist.UpdatedAt = sql.NullTime{Time: time.Now(), Valid: true}

	return ass.animeStudioRepository.Update(ctx, &exist)

}

func (ass AnimeStudioService) Delete (ctx context.Context, id string) error {
    exist, err := ass.animeStudioRepository.FindById(ctx, id)

    if err != nil && exist.Id == "" {
        return  utility.NewNotFound("Anime Studio")
    }
    
    if err != nil {
        return err
    }

    return ass.animeStudioRepository.Delete(ctx, id)
} 