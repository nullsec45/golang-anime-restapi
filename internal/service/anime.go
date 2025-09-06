package service

import(
	"github.com/nullsec45/golang-anime-restapi/domain"
	"github.com/nullsec45/golang-anime-restapi/dto"
	"github.com/nullsec45/golang-anime-restapi/internal/utility"
	"context"
	"github.com/google/uuid"
	"database/sql"
	"time"
	"errors"
	// "encoding/json"
)

type AnimeService struct {
	animeRepository domain.AnimeRepository
}

func NewAnime(animeRepository domain.AnimeRepository) domain.AnimeService {
	return &AnimeService{
		animeRepository: animeRepository,
	}
}

func (as AnimeService) Index(ctx context.Context) ([]dto.AnimeData, error) {
	animes, err := as.animeRepository.FindAll(ctx)

	if err != nil {
		return nil, err
	}

	var animeData []dto.AnimeData

	for _, v:= range animes {
		animeData = append(animeData, dto.AnimeData{
			Id:                     v.Id,
			Slug:                   v.Slug,
			TitleRomaji:            v.TitleRomaji,
			TitleNative:            utility.ToString(v.TitleNative),         
			TitleEnglish:           utility.ToString(v.TitleEnglish),        
			Synopsis:               utility.ToString(v.Synopsis),            	
			Type:                   dto.AnimeType(v.Type),
			Season:                 v.Season,
			SeasonYear:             v.SeasonYear,         
			Status:                 dto.AnimeStatus(v.Status), 
			AgeRating:              v.AgeRating,
			TotalEpisodes:          v.TotalEpisodes,            
			AverageDurationMinutes: v.AverageDurationMinutes,   
			Country:                v.Country,
			PremieredAt:            utility.ToTimePtr(v.PremieredAt),              
			EndedAt:                utility.ToTimePtr(v.EndedAt),                  
			Popularity:             v.Popularity,
			ScoreAvg:               v.ScoreAvg,                 
			AltTitles:              utility.ToRawMessage(v.AltTitles),     
			ExternalIDs:            utility.ToRawMessage(v.ExternalIDs),       
		})
	}

	return animeData, nil
}

func (as AnimeService) Show (ctx context.Context, id string) (dto.AnimeData, error) {
    exist, err := as.animeRepository.FindById(ctx,id)

    if err != nil && exist.Id == "" {
        return dto.AnimeData{}, errors.New("Data anime tidak ditemukan!.")
    }
    
    if err != nil {
        return dto.AnimeData{}, err
    }

    return dto.AnimeData{
      		Id:                     exist.Id,
			Slug:                   exist.Slug,
			TitleRomaji:            exist.TitleRomaji,
			TitleNative:            utility.ToString(exist.TitleNative),         
			TitleEnglish:           utility.ToString(exist.TitleEnglish),        
			Synopsis:               utility.ToString(exist.Synopsis),            	
			Type:                   dto.AnimeType(exist.Type),
			Season:                 exist.Season,
			SeasonYear:             exist.SeasonYear,         
			Status:                 dto.AnimeStatus(exist.Status), 
			AgeRating:              exist.AgeRating,
			TotalEpisodes:          exist.TotalEpisodes,            
			AverageDurationMinutes: exist.AverageDurationMinutes,   
			Country:                exist.Country,
			PremieredAt:            utility.ToTimePtr(exist.PremieredAt),              
			EndedAt:                utility.ToTimePtr(exist.EndedAt),                  
			Popularity:             exist.Popularity,
			ScoreAvg:               exist.ScoreAvg,                 
			AltTitles:              utility.ToRawMessage(exist.AltTitles),     
			ExternalIDs:            utility.ToRawMessage(exist.ExternalIDs),   
    }, nil
}

func (as AnimeService) Create(ctx context.Context, req dto.CreateAnimeRequest) error {
 	anime := domain.Anime{
        Id: uuid.New().String(),
		Slug:req.Slug,
		TitleRomaji: req.TitleRomaji,
		TitleNative : sql.NullString{String:req.TitleNative, Valid:true},
		TitleEnglish: sql.NullString{String:req.TitleEnglish, Valid:true},
		Synopsis: sql.NullString{String:req.Synopsis, Valid:true},
		Type: utility.ToAnimeType(req.Type),  
		Status: utility.ToAnimeStatus(req.Status),
		Season: utility.ToSeason(req.Season),
		SeasonYear: req.SeasonYear,
		AgeRating: utility.ToAgeRating(req.AgeRating),
		TotalEpisodes: req.TotalEpisodes,
		AverageDurationMinutes: req.AverageDurationMinutes,
		Country: req.Country,
		PremieredAt: utility.ToSqlNullTime(req.PremieredAt),
		EndedAt: utility.ToSqlNullTime(req.EndedAt),
		Popularity: req.Popularity,
		ScoreAvg: req.ScoreAvg,
		AltTitles: utility.ToRawMessage(req.AltTitles),
		ExternalIDs: utility.ToRawMessage(req.ExternalIDs),
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
    }

	return as.animeRepository.Save(ctx, &anime)
}

func (as AnimeService) Update(ctx context.Context, req dto.UpdateAnimeRequest)  error {
    // Cari data anime
    exist, err := as.animeRepository.FindById(ctx, req.Id)
    // fmt.Println(err)

    // Jika anime tidak ditemukan
    if err != nil && exist.Id == "" {
        return errors.New("Data anime tidak ditemukan!.")
    }
    
    if err != nil {
        return  err
    }

    // Update data sesuai request
	exist.Slug = req.Slug
	exist.TitleRomaji = req.TitleRomaji
	exist.TitleNative  = sql.NullString{String:req.TitleNative, Valid:true}
	exist.TitleEnglish = sql.NullString{String:req.TitleEnglish, Valid:true}
	exist.Synopsis = sql.NullString{String:req.Synopsis, Valid:true}
	exist.Type = utility.ToAnimeType(req.Type)  
	exist.Status = utility.ToAnimeStatus(req.Status)
	exist.Season = utility.ToSeason(req.Season)
	exist.SeasonYear = req.SeasonYear
	exist.AgeRating = utility.ToAgeRating(req.AgeRating)
	exist.TotalEpisodes = req.TotalEpisodes
	exist.AverageDurationMinutes = req.AverageDurationMinutes
	exist.Country =  req.Country
	exist.PremieredAt = utility.ToSqlNullTime(req.PremieredAt)
	exist.EndedAt =  utility.ToSqlNullTime(req.EndedAt)
	exist.Popularity = req.Popularity
	exist.ScoreAvg = req.ScoreAvg
	exist.AltTitles = utility.ToRawMessage(req.AltTitles)
	exist.ExternalIDs = utility.ToRawMessage(req.ExternalIDs)
	exist.CreatedAt = sql.NullTime{Time: time.Now(), Valid: true}
	exist.UpdatedAt = sql.NullTime{Time: time.Now(), Valid: true}

    // Simpan perubahan
    // err = as.animeRepository.Update(ctx, &exist)

    // if err != nil {
    //     return nil, err
    // }

    // Buat response DTO
    // updatedAnime := dto.AnimeData{
    //     Id: uuid.New().String(),
	// 	Slug:req.Slug,
	// 	TitleRomaji: req.TitleRomaji,
	// 	TitleNative : utility.ToString(req.TitleNative),
	// 	TitleEnglish: utility.ToString(req.TitleEnglish),
	// 	Synopsis: utility.ToString(req.Synopsis),
	// 	Type: string(req.Type),  
	// 	Status: string(req.Status),
	// 	Season: utility.ToString(seasonToString(req.Season)),
	// 	SeasonYear: nint16(req.SeasonYear),
	// 	AgeRating: utility.ToString(ageToString(req.AgeRating)),
	// 	TotalEpisodes: utility.nint(req.TotalEpisodes),
	// 	AverageDurationMinutes: utility.nint(req.AverageDurationMinutes),
	// 	Country: sql.NullString{String: firstNonEmpty(ptrToString(req.Country), "JP"), Valid: true},
	// 	PremieredAt: utility.ntime(req.PremieredAt),
	// 	EndedAt: utility.ntime(req.EndedAt),
	// 	Popularity: sql.NullInt32{Int32: int32(ptrToInt(req.Popularity)), Valid: req.Popularity != nil},
	// 	ScoreAvg: utility.nfloat32(req.ScoreAvg),
	// 	AltTitles: utility.njson(req.AltTitles),
	// 	ExternalIDs: utility.njson(req.ExternalIDs),
	// 	CreatedAt: sql.NullTime{Time: now, Valid: true},
	// 	UpdatedAt: sql.NullTime{Time: now, Valid: true},
    // }

    // return []dto.AnimeData{updatedAnime}, nil
	return as.animeRepository.Update(ctx, &exist)

}

func (as AnimeService) Delete (ctx context.Context, id string) error {
    exist, err := as.animeRepository.FindById(ctx, id)

    if err != nil && exist.Id == "" {
        return  errors.New("Data anime tidak ditemukan!.")
    }
    
    if err != nil {
        return err
    }

    return as.animeRepository.Delete(ctx, id)
}