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
			TitleNative:            v.TitleNative,         
			TitleEnglish:           v.TitleEnglish,        
			Synopsis:               v.Synopsis,            	
			Type:                   dto.AnimeType(v.Type),
			Season:                 utility.toSeasonPtr(v.Season),
			SeasonYear:             v.SeasonYear,         
			Status:                 dto.AnimeStatus(v.Status), 
			AgeRating:              utility.toAgeRatingPtr(v.AgeRating),
			TotalEpisodes:          v.TotalEpisodes,            
			AverageDurationMinutes: v.AverageDurationMinutes,   
			Country:                v.Country,
			PremieredAt:            v.PremieredAt,              
			EndedAt:                v.EndedAt,                  
			Popularity:             v.Popularity,
			ScoreAvg:               v.ScoreAvg,                 
			AltTitles:              utility.toJSONMap(v.AltTitles),     
			ExternalIDs:            utility.toJSONMap(v.ExternalIDs),       
		})
	}
}

func (as AnimeService) Show (ctx context.Context, id string) (dto.AnimeData, error) {
    exist, err := c.customerRepository.FindById(ctx,id)

    if err != nil && exist.ID == "" {
        return dto.AnimeData{}, errors.New("Data anime tidak ditemukan!.")
    }
    
    if err != nil {
        return dto.AnimeData{}, err
    }

    return dto.AnimeData{
      		Id:                     v.Id,
			Slug:                   v.Slug,
			TitleRomaji:            v.TitleRomaji,
			TitleNative:            v.TitleNative,         
			TitleEnglish:           v.TitleEnglish,        
			Synopsis:               v.Synopsis,            	
			Type:                   dto.AnimeType(v.Type),
			Season:                 utility.toSeasonPtr(v.Season),
			SeasonYear:             v.SeasonYear,         
			Status:                 dto.AnimeStatus(v.Status), 
			AgeRating:              utility.toAgeRatingPtr(v.AgeRating),
			TotalEpisodes:          v.TotalEpisodes,            
			AverageDurationMinutes: v.AverageDurationMinutes,   
			Country:                v.Country,
			PremieredAt:            v.PremieredAt,              
			EndedAt:                v.EndedAt,                  
			Popularity:             v.Popularity,
			ScoreAvg:               v.ScoreAvg,                 
			AltTitles:              utility.toJSONMap(v.AltTitles),     
			ExternalIDs:            utility.toJSONMap(v.ExternalIDs),   
    }, nil
}

func (as AnimeService) Create(ctx context.Context, req dto.CreateAnimeRequest) error {
 	anime := domain.Anime{
        Id: uuid.New().String(),
		Slug:req.Slug,
		TitleRomaji: req.TitleRomaji,
		TitleNative : utility.nstr(req.TitleNative),
		TitleEnglish: utility.nstr(req.TitleEnglish),
		Synopsis: utility.nstr(req.Synopsis),
		Type: string(req.Type),  
		Status: string(req.Status),
		Season: utility.nstr(seasonToString(req.Season)),
		SeasonYear: nint16(req.SeasonYear),
		AgeRating: utility.nstr(ageToString(req.AgeRating)),
		TotalEpisodes: utility.nint(req.TotalEpisodes),
		AverageDurationMinutes: utility.nint(req.AverageDurationMinutes),
		Country: sql.NullString{String: firstNonEmpty(ptrToString(req.Country), "JP"), Valid: true},
		PremieredAt: utility.ntime(req.PremieredAt),
		EndedAt: utility.ntime(req.EndedAt),
		Popularity: sql.NullInt32{Int32: int32(ptrToInt(req.Popularity)), Valid: req.Popularity != nil},
		ScoreAvg: utility.nfloat32(req.ScoreAvg),
		AltTitles: utility.njson(req.AltTitles),
		ExternalIDs: utility.njson(req.ExternalIDs),
		CreatedAt: sql.NullTime{Time: now, Valid: true},
		UpdatedAt: sql.NullTime{Time: now, Valid: true},
    }

	return as.animeRepository.Save(ctx, &anime)
}

func (as AnimeService) Update(ctx context.Context, req dto.UpdateAnimeRequest)  error {
    // Cari data anime
    exist, err := c.animeRepository.FindById(ctx, req.ID)
    // fmt.Println(err)

    // Jika anime tidak ditemukan
    if err != nil && exist.Id == "" {
        return nil, errors.New("Data anime tidak ditemukan!.")
    }
    
    if err != nil {
        return nil, err
    }

    // Update data sesuai request
	exist.Slug = req.Slug
	exist.TitleRomaji = req.TitleRomaji
	exist.TitleNative  = utility.nstr(req.TitleNative)
	exist.TitleEnglish = utility.nstr(req.TitleEnglish)
	exist.Synopsis = utility.nstr(req.Synopsis)
	exist.Type = string(req.Type)  
	exist.Status = string(req.Status)
	exist.Season = utility.nstr(seasonToString(req.Season))
	exist.SeasonYear = nint16(req.SeasonYear)
	exist.AgeRating = utility.nstr(ageToString(req.AgeRating))
	exist.TotalEpisodes = utility.nint(req.TotalEpisodes)
	exist.AverageDurationMinutes = utility.nint(req.AverageDurationMinutes)
	exist.Country =  sql.NullString{String: firstNonEmpty(ptrToString(req.Country), "JP"), Valid: true}
	exist.PremieredAt = utility.ntime(req.PremieredAt)
	exist.EndedAt =  utility.ntime(req.EndedAt)
	exist.Popularity = sql.NullInt32{Int32: int32(ptrToInt(req.Popularity)), Valid: req.Popularity != nil}
	exist.ScoreAvg = utility.nfloat32(req.ScoreAvg)
	exist.AltTitles = utility.njson(req.AltTitles)
	exist.ExternalIDs = utility.njson(req.ExternalIDs)
	exist.CreatedAt = sql.NullTime{Time: now, Valid: true}
	exist.UpdatedAt = sql.NullTime{Time: now, Valid: true}

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
	// 	TitleNative : utility.nstr(req.TitleNative),
	// 	TitleEnglish: utility.nstr(req.TitleEnglish),
	// 	Synopsis: utility.nstr(req.Synopsis),
	// 	Type: string(req.Type),  
	// 	Status: string(req.Status),
	// 	Season: utility.nstr(seasonToString(req.Season)),
	// 	SeasonYear: nint16(req.SeasonYear),
	// 	AgeRating: utility.nstr(ageToString(req.AgeRating)),
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