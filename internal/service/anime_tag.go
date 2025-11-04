package service

import(
	"github.com/nullsec45/golang-anime-restapi/domain"
	"github.com/nullsec45/golang-anime-restapi/dto"
	"context"
	"github.com/google/uuid"
	"database/sql"
	"time"
	"github.com/nullsec45/golang-anime-restapi/internal/utility"
)

type AnimeTagsService struct {
	animeRepository domain.AnimeRepository
	tagRepository domain.AnimeTagRepository
	animeTagsRepository domain.AnimeTagsRepository
}

func NewAnimeTags(
	animeRepository domain.AnimeRepository,
	tagRepository domain.AnimeTagRepository,
	animeTagsRepository domain.AnimeTagsRepository) domain.AnimeTagsService {
	return &AnimeTagsService{
		animeRepository: animeRepository,
		tagRepository: tagRepository,
		animeTagsRepository: animeTagsRepository,
	}
}

func (ats AnimeTagsService) Create(ctx context.Context, req dto.CreateAnimeTagsRequest) error {
	anime, errAnime := ats.animeRepository.FindById(ctx, req.AnimeId)

	if anime.Id == "" {
		return utility.NewNotFound("Anime")
	}

	if errAnime != nil {
		return errAnime
	}

	tag, errTag := ats.tagRepository.FindById(ctx, req.TagId)

	if tag.Id == "" {
		return utility.NewNotFound("Anime Tag")
	}

	if errTag != nil {
		return errTag
	}

	_,  found, errAnimeTags := ats.animeTagsRepository.FindByAnimeAndTagId(ctx, req.AnimeId, req.TagId)

	if errAnimeTags != nil {
		return errAnimeTags
	}

	if found {
		return utility.NewAlreadyExist("Anime Tags")
	}	
	
 	ag := domain.AnimeTags{
		Id:        uuid.NewString(),
		AnimeId:   req.AnimeId,
		TagId:     req.TagId,
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
	}

	return ats.animeTagsRepository.Save(ctx, &ag)
}

func (ats AnimeTagsService) Update(ctx context.Context, req dto.UpdateAnimeTagsRequest) error {
	exist, err := ats.animeTagsRepository.FindById(ctx, req.Id)

    if err != nil && exist.Id == "" {
        return utility.NewNotFound("Anime Tags")
    }
    
    if err != nil {
        return  err
    }

	anime, errAnime := ats.animeRepository.FindById(ctx, req.AnimeId)

	if anime.Id == "" {
		return utility.NewNotFound("Anime")
	}

	if errAnime != nil {
		return err
	}

	tag, errTag := ats.tagRepository.FindById(ctx, req.TagId)

	if tag.Id == "" {
		return utility.NewNotFound("Anime Tag")
	}

	if errTag != nil {
		return err
	}

	_, found, errAnimeTags := ats.animeTagsRepository.FindByAnimeAndTagId(ctx, req.AnimeId, req.TagId)

	if errAnimeTags != nil {
		return errAnimeTags
	}

	if found {
		return utility.NewAlreadyExist("Anime Tags")
	}

	exist.AnimeId = req.AnimeId
	exist.TagId = req.TagId
	exist.UpdatedAt = sql.NullTime{Time: time.Now(), Valid: true}

	return ats.animeTagsRepository.Update(ctx, &exist)
}


func (ats AnimeTagsService) DeleteByAnimeId (ctx context.Context, animeId string) error {
    exist, err := ats.animeRepository.FindById(ctx, animeId)

    if err != nil && exist.Id == "" {
        return  utility.NewNotFound("Anime")
    }
    
    if err != nil {
        return err
    }

    return ats.animeTagsRepository.DeleteByAnimeId(ctx, animeId)
}

func (ats AnimeTagsService) DeleteByTagId (ctx context.Context, tagId string) error {
    exist, err := ats.tagRepository.FindById(ctx, tagId)

    if err != nil && exist.Id == "" {
        return  utility.NewNotFound("Anime Tag")
    }
    
    if err != nil {
        return err
    }

    return ats.animeTagsRepository.DeleteByTagId(ctx, tagId)
}

func (ats AnimeTagsService) DeleteById (ctx context.Context, Id string) error {
    exist, err := ats.animeTagsRepository.FindById(ctx, Id)

    if err != nil && exist.Id == "" {
        return  utility.NewNotFound("Anime Tags")
    }
    
    if err != nil {
        return err
    }

    return ats.animeTagsRepository.DeleteById(ctx, Id)
}