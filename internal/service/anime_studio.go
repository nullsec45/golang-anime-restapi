package service

import(
	"github.com/nullsec45/golang-anime-restapi/domain"
	"github.com/nullsec45/golang-anime-restapi/dto"
	"context"
	"github.com/google/uuid"
	"database/sql"
	"time"
	// "fmt"
)

type AnimeStudiosService struct {
	animeRepository domain.AnimeRepository
	studioRepository domain.AnimeStudioRepository
	animeStudiosRepository domain.AnimeStudiosRepository
}

func NewAnimeStudios(
	animeRepository domain.AnimeRepository,
	studioRepository domain.AnimeStudioRepository,
	animeStudiosRepository domain.AnimeStudiosRepository) domain.AnimeStudiosService {
	return &AnimeStudiosService{
		animeRepository: animeRepository,
		studioRepository: studioRepository,
		animeStudiosRepository: animeStudiosRepository,
	}
}

func (astds AnimeStudiosService) Create(ctx context.Context, req dto.CreateAnimeStudiosRequest) error {
	anime, errAnime := astds.animeRepository.FindById(ctx, req.AnimeId)

	if anime.Id == "" {
		return domain.AnimeNotFound
	}

	if errAnime != nil {
		return errAnime
	}

	studio, errStudio := astds.studioRepository.FindById(ctx, req.StudioId)

	if studio.Id == "" {
		return domain.AnimeStudioNotFound
	}

	if errStudio != nil {
		return errStudio
	}

	_, found, errAnimeStudios := astds.animeStudiosRepository.FindByAnimeAndStudioId(ctx, req.AnimeId, req.StudioId)


	if errAnimeStudios != nil {
		return errAnimeStudios
	}

	if found {
		return domain.AnimeStudiosAlready
	}	

 	ag := domain.AnimeStudios{
		Id:        uuid.NewString(),
		AnimeId:   req.AnimeId,
		StudioId:   req.StudioId,
		Role:req.Role,
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
	}

	return astds.animeStudiosRepository.Save(ctx, &ag)
}


func (astds AnimeStudiosService) Update(ctx context.Context, req dto.UpdateAnimeStudiosRequest) error {
	exist, err := astds.animeStudiosRepository.FindById(ctx, req.Id)

    if err != nil && exist.Id == "" {
        return domain.AnimeStudiosNotFound
    }
    
    if err != nil {
        return  err
    }

	anime, errAnime := astds.animeRepository.FindById(ctx, req.AnimeId)

	if anime.Id == "" {
		return domain.AnimeNotFound
	}

	if errAnime != nil {
		return err
	}

	studio, errStudio := astds.studioRepository.FindById(ctx, req.StudioId)

	if studio.Id == "" {
		return domain.AnimeStudioNotFound
	}

	if errStudio != nil {
		return err
	}

	if exist.Role == req.Role {
		_, found, errAnimeStudios := astds.animeStudiosRepository.FindByAnimeAndStudioId(ctx, req.AnimeId, req.StudioId)

		if errAnimeStudios != nil {
			return errAnimeStudios
		}

		if found {
			return domain.AnimeStudiosAlready
		}	
	}
	
	exist.AnimeId = req.AnimeId
	exist.StudioId = req.StudioId
	exist.Role=req.Role
	exist.UpdatedAt = sql.NullTime{Time: time.Now(), Valid: true}

	return astds.animeStudiosRepository.Update(ctx, &exist)
}


func (astds AnimeStudiosService) DeleteByAnimeId (ctx context.Context, animeId string) error {
    exist, err := astds.animeRepository.FindById(ctx, animeId)

    if err != nil && exist.Id == "" {
        return  domain.AnimeNotFound
    }
    
    if err != nil {
        return err
    }

    return astds.animeStudiosRepository.DeleteByAnimeId(ctx, animeId)
}

func (astds AnimeStudiosService) DeleteByStudioId (ctx context.Context, studioId string) error {
    exist, err := astds.studioRepository.FindById(ctx, studioId)

    if err != nil && exist.Id == "" {
        return  domain.AnimeStudioNotFound
    }
    
    if err != nil {
        return err
    }

    return astds.animeStudiosRepository.DeleteByStudioId(ctx, studioId)
}

func (astds AnimeStudiosService) DeleteById (ctx context.Context, Id string) error {
    exist, err := astds.animeStudiosRepository.FindById(ctx, Id)

    if err != nil && exist.Id == "" {
        return  domain.AnimeStudiosNotFound
    }
    
    if err != nil {
        return err
    }

    return astds.animeStudiosRepository.DeleteById(ctx, Id)
}