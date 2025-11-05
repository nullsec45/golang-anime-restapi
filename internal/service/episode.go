package service

import(
	"github.com/nullsec45/golang-anime-restapi/domain"
	"github.com/nullsec45/golang-anime-restapi/dto"
	"github.com/nullsec45/golang-anime-restapi/internal/utility"
	"context"
	"github.com/google/uuid"
	"database/sql"
	"time"
	// "encoding/json"
	// "fmt"
	"github.com/nullsec45/golang-anime-restapi/internal/config"
)

type AnimeEpisodeService struct {
	animeRepository domain.AnimeRepository
	animeEpisodeRepository domain.AnimeEpisodeRepository
	mediaRepository domain.MediaRepository
	config *config.Config
}

func NewAnimeEpisode(
	animeRepository domain.AnimeRepository,
	animeEpisodeRepository domain.AnimeEpisodeRepository,
	mediaRepository domain.MediaRepository,
	config *config.Config) domain.AnimeEpisodeService{
	return &AnimeEpisodeService{
		animeRepository: animeRepository,
		animeEpisodeRepository:animeEpisodeRepository,
		mediaRepository:mediaRepository,
		config:config,
	}
}

func (as AnimeEpisodeService) Index(ctx context.Context, animeId string, epts domain.EpisodeListOptions) (dto.Paginated[dto.AnimeEpisodeData], error) {
	items, total, err := as.animeEpisodeRepository.FindAll(ctx, animeId, epts)

	if err != nil {
		return dto.Paginated[dto.AnimeEpisodeData]{}, err
	}

	var episodeData []dto.AnimeEpisodeData

	for _, v:= range items {

		episodeData = append(episodeData, dto.AnimeEpisodeData{
			Id:               v.Id,
			AnimeId:          v.AnimeId, 
			Number:           v.Number,
			SeasonNumber:     v.SeasonNumber,       
			Title:            v.Title,
			Synopsis:         v.Synopsis,
			AirDate:          utility.ToTimePtr(v.AirDate),
			DurationMinutes:  v.DurationMinutes,
			IsSpecial:        v.IsSpecial,
		})
	}

	return dto.Paginated[dto.AnimeEpisodeData]{
		Data: episodeData,
		Meta:epts.Pagination.BuildMeta(total),
	}, nil
}

func (as  AnimeEpisodeService) Show (ctx context.Context, id string) (dto.AnimeEpisodeData, error) {
	exist, err := as.animeEpisodeRepository.FindById(ctx, id)

    if err != nil && exist.Id == "" {
        return dto.AnimeEpisodeData{}, utility.NewNotFound("Anime Episode")
    }

	var video string

	if exist.Video.Valid {
		media, _ := as.mediaRepository.FindById(ctx, exist.Video.String)

		if media.Path != "" {
			video = as.config.Server.AssetPrivate+"/"+media.Id
		}
	}

    return dto.AnimeEpisodeData{
		Id:               exist.Id,
		AnimeId:          exist.AnimeId,
		Number:           exist.Number,       
		SeasonNumber:     exist.SeasonNumber,
		Title:            exist.Title, 
		Synopsis:         exist.Synopsis,
		AirDate:          utility.ToTimePtr(exist.AirDate),
		DurationMinutes:  exist.DurationMinutes,
		IsSpecial:        exist.IsSpecial,
		Video:            video,
    }, nil
}

func (as AnimeEpisodeService) Create(ctx context.Context, req dto.CreateAnimeEpisodeRequest) error {
	anime, err := as.animeRepository.FindById(ctx, req.AnimeId)

	if anime.Id == "" {
		return utility.NewNotFound("Anime")
	}

	if err != nil {
		return err
	}

	video := sql.NullString{String:req.Video, Valid:false}

	if req.Video != "" {
		video.Valid = true 

		media, err := as.mediaRepository.FindById(ctx, req.Video)

		if err != nil && media.Id == "" {
       		return utility.NewNotFound("Anime Media")
		} 
	}
	
 	ep := domain.AnimeEpisode{
		Id:              uuid.NewString(),
		AnimeId:         req.AnimeId,
		Number:          req.Number,
		SeasonNumber:    req.SeasonNumber,
		Title:           req.Title,
		Synopsis:        req.Synopsis,
		AirDate:         utility.ToSqlNullTime(req.AirDate),
		DurationMinutes: req.DurationMinutes,
		IsSpecial:       req.IsSpecial,
		Video:           video,
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
	}

	return as.animeEpisodeRepository.Save(ctx, &ep)
}

func (as AnimeEpisodeService) Update(ctx context.Context, req dto.UpdateAnimeEpisodeRequest)  error {
	anime, err := as.animeRepository.FindById(ctx, req.AnimeId)

	if anime.Id == "" {
		return utility.NewNotFound("Anime")
	}

	if err != nil {
		return err
	}

    exist, err := as.animeEpisodeRepository.FindById(ctx, req.Id)

    if err != nil && exist.Id == "" {
        return utility.NewNotFound("Anime Episode")
    }
    
    if err != nil {
        return  err
    }

	video := sql.NullString{String:req.Video, Valid:false}

	if req.Video != "" {
		video.Valid = true 

		media, err := as.mediaRepository.FindById(ctx, req.Video)

		if err != nil && media.Id == "" {
       		return utility.NewNotFound("Anime Media")
		} 
	}			
	
	exist.Number = req.Number
	exist.SeasonNumber=req.SeasonNumber
	exist.Title=req.Title
	exist.Synopsis=req.Synopsis
	exist.AirDate=utility.ToSqlNullTime(req.AirDate)
	exist.DurationMinutes=req.DurationMinutes
	exist.IsSpecial=req.IsSpecial
	exist.Video=video
	exist.UpdatedAt = sql.NullTime{Time: time.Now(), Valid: true}

	return as.animeEpisodeRepository.Update(ctx, &exist)
}

func (as AnimeEpisodeService) DeleteByAnimeId (ctx context.Context, animeId string) error {
    exist, err := as.animeRepository.FindById(ctx, animeId)

    if err != nil && exist.Id == "" {
        return  utility.NewNotFound("Anime")
    }
    
    if err != nil {
        return err
    }

    return as.animeEpisodeRepository.DeleteByAnimeId(ctx, animeId)
}

func (as AnimeEpisodeService) DeleteById (ctx context.Context, animeId string) error {
    exist, err := as.animeEpisodeRepository.FindById(ctx, animeId)

    if err != nil && exist.Id == "" {
        return  utility.NewNotFound("Anime")
    }
    
    if err != nil {
        return err
    }

    return as.animeEpisodeRepository.DeleteById(ctx, animeId)
}