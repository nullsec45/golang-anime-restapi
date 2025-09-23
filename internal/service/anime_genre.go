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

type AnimeGenresService struct {
	animeRepository domain.AnimeRepository
	genreRepository domain.AnimeGenreRepository
	animeGenresRepository domain.AnimeGenresRepository
}

func NewAnimeGenres(
	animeRepository domain.AnimeRepository,
	genreRepository domain.AnimeGenreRepository,
	animeGenresRepository domain.AnimeGenresRepository) domain.AnimeGenresService {
	return &AnimeGenresService{
		animeRepository: animeRepository,
		genreRepository: genreRepository,
		animeGenresRepository: animeGenresRepository,
	}
}

func (agrs AnimeGenresService) Create(ctx context.Context, req dto.CreateAnimeGenresRequest) error {
	anime, errAnime := agrs.animeRepository.FindById(ctx, req.AnimeId)

	if anime.Id == "" {
		return domain.AnimeNotFound
	}

	if errAnime != nil {
		return errAnime
	}

	genre, errGenre := agrs.genreRepository.FindById(ctx, req.GenreId)

	if genre.Id == "" {
		return domain.AnimeGenreNotFound
	}

	if errGenre != nil {
		return errGenre
	}

	_, found, errAnimeGenres := agrs.animeGenresRepository.FindByAnimeAndGenreId(ctx, req.AnimeId, req.GenreId)


	if errAnimeGenres != nil {
		return errAnimeGenres
	}

	if found {
		return domain.AnimeGenresAlready
	}	

 	ag := domain.AnimeGenres{
		Id:        uuid.NewString(),
		AnimeId:   req.AnimeId,
		GenreId:   req.GenreId,
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
	}

	return agrs.animeGenresRepository.Save(ctx, &ag)
}


func (agrs AnimeGenresService) Update(ctx context.Context, req dto.UpdateAnimeGenresRequest) error {
	exist, err := agrs.animeGenresRepository.FindById(ctx, req.Id)

    if err != nil && exist.Id == "" {
        return domain.AnimeGenresNotFound
    }
    
    if err != nil {
        return  err
    }

	anime, errAnime := agrs.animeRepository.FindById(ctx, req.AnimeId)

	if anime.Id == "" {
		return domain.AnimeNotFound
	}

	if errAnime != nil {
		return err
	}

	genre, errGenre := agrs.genreRepository.FindById(ctx, req.GenreId)

	if genre.Id == "" {
		return domain.AnimeGenreNotFound
	}

	if errGenre != nil {
		return err
	}

	_, found, errAnimeGenres := agrs.animeGenresRepository.FindByAnimeAndGenreId(ctx, req.AnimeId, req.GenreId)

	if errAnimeGenres != nil {
		return errAnimeGenres
	}

	if found {
		return domain.AnimeGenresAlready
	}	


	exist.AnimeId = req.AnimeId
	exist.GenreId = req.GenreId
	exist.UpdatedAt = sql.NullTime{Time: time.Now(), Valid: true}

	return agrs.animeGenresRepository.Update(ctx, &exist)
}


func (agrs AnimeGenresService) DeleteByAnimeId (ctx context.Context, animeId string) error {
    exist, err := agrs.animeRepository.FindById(ctx, animeId)

    if err != nil && exist.Id == "" {
        return  domain.AnimeNotFound
    }
    
    if err != nil {
        return err
    }

    return agrs.animeGenresRepository.DeleteByAnimeId(ctx, animeId)
}

func (agrs AnimeGenresService) DeleteByGenreId (ctx context.Context, genreId string) error {
    exist, err := agrs.genreRepository.FindById(ctx, genreId)

    if err != nil && exist.Id == "" {
        return  domain.AnimeGenreNotFound
    }
    
    if err != nil {
        return err
    }

    return agrs.animeGenresRepository.DeleteByGenreId(ctx, genreId)
}

func (agrs AnimeGenresService) DeleteById (ctx context.Context, Id string) error {
    exist, err := agrs.animeGenresRepository.FindById(ctx, Id)

    if err != nil && exist.Id == "" {
        return  domain.AnimeGenresNotFound
    }
    
    if err != nil {
        return err
    }

    return agrs.animeGenresRepository.DeleteById(ctx, Id)
}