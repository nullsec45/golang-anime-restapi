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
)

type AnimeEpisodeService struct {
	animeRepository domain.AnimeRepository
	animeEpisodeRepository domain.AnimeEpisodeRepository
}

func NewAnimeEpisode(
	animeRepository domain.AnimeRepository,
	animeEpisodeRepository domain.AnimeEpisodeRepository) domain.AnimeEpisodeService {
	return &AnimeEpisodeService{
		animeRepository: animeRepository,
		animeEpisodeRepository:animeEpisodeRepository,
	}
}

func (as AnimeEpisodeService) Create(ctx context.Context, req dto.CreateAnimeEpisodeRequest) error {
	anime, err := as.animeRepository.FindById(ctx, req.AnimeId)

	if err != nil {
		return err
	}

	if anime.Id == "" {
		return domain.AnimeNotFound
	}

	// var airDate sql.NullTime
	// if s := strings.TrimSpace(req.AirDate); s != "" {
	// 	d, err := time.Parse("2006-01-02", s)
	// 	if err != nil {
	// 		return errors.New("air_date harus format YYYY-MM-DD")
	// 	}
	// 	d = time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, time.UTC)
	// 	airDate = sql.NullTime{Time: d, Valid: true}
	// }

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
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
	}

	// if as.episodeRepository.ExistsByAnimeAndNumber != nil {
	// 	if exists, err := as.episodeRepository.ExistsByAnimeAndNumber(ctx, req.AnimeId, req.Number); err != nil {
	// 		return err
	// 	} else if exists {
	// 		return domain.EpisodeNumberAlreadyExists
	// 	}
	// }

	return as.animeEpisodeRepository.Save(ctx, &ep)
}


func (as AnimeEpisodeService) DeleteByAnimeId (ctx context.Context, animeId string) error {
    exist, err := as.animeRepository.FindById(ctx, animeId)

    if err != nil && exist.Id == "" {
        return  domain.AnimeNotFound
    }
    
    if err != nil {
        return err
    }

    return as.animeEpisodeRepository.DeleteByAnimeId(ctx, animeId)
}

func (as AnimeEpisodeService) DeleteById (ctx context.Context, animeId string) error {
    exist, err := as.animeEpisodeRepository.FindById(ctx, animeId)

    if err != nil && exist.Id == "" {
        return  domain.AnimeNotFound
    }
    
    if err != nil {
        return err
    }

    return as.animeEpisodeRepository.DeleteById(ctx, animeId)
}