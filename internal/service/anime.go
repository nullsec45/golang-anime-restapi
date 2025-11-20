package service

import(
	"github.com/nullsec45/golang-anime-restapi/domain"
	"github.com/nullsec45/golang-anime-restapi/dto"
	"github.com/nullsec45/golang-anime-restapi/internal/utility"
	"github.com/nullsec45/golang-anime-restapi/internal/config"
	"context"
	"github.com/google/uuid"
	"database/sql"
	"time"
	"github.com/gosimple/slug"
	"fmt"
)

type AnimeService struct {
	config *config.Config
	animeRepository domain.AnimeRepository
	animeEpisodeRepository domain.AnimeEpisodeRepository
	animeGenreRepository domain.AnimeGenreRepository
	animeTagRepository domain.AnimeTagRepository
	mediaRepository domain.MediaRepository
	animeStudioRepository domain.AnimeStudioRepository
	voiceCastRepository domain.VoiceCastRepository
}

func NewAnime(
	config *config.Config,
	animeRepository domain.AnimeRepository,
	animeEpisodeRepository domain.AnimeEpisodeRepository,
	animeGenreRepository domain.AnimeGenreRepository,
	animeTagRepository domain.AnimeTagRepository,
	mediaRepository domain.MediaRepository,
	animeStudioRepository domain.AnimeStudioRepository,
	voiceCastRepository domain.VoiceCastRepository,
) domain.AnimeService {
	return &AnimeService{
		config:config,
		animeRepository: animeRepository,
		animeEpisodeRepository:animeEpisodeRepository,
		animeGenreRepository:animeGenreRepository,
		animeTagRepository:animeTagRepository,
		mediaRepository:mediaRepository,
		animeStudioRepository:animeStudioRepository,
		voiceCastRepository:voiceCastRepository,
	}
}

func (as AnimeService) Index(ctx context.Context, opts domain.AnimeListOptions) (dto.Paginated[dto.AnimeData], error) {
	items, total, err := as.animeRepository.FindAll(ctx, opts)

	if err != nil {
		return dto.Paginated[dto.AnimeData]{}, err
	}

	coverId := make([]string, 0)
	for _, v := range items {
		if v.CoverId.Valid {
			coverId = append(coverId, v.CoverId.String)
		}
	}

	covers := make(map[string]string)

	if len(coverId) > 0 {	
		media, _ := as.mediaRepository.FindByIds(ctx, coverId)

		for _, v := range media {
			covers[v.Id] = as.config.Server.AssetPrivate+"/"+v.Id
		}
	}

	var animeData []dto.AnimeData

	for _, v:= range items {
		var coverUrl string

		if v2, e := covers[v.CoverId.String]; e {
			coverUrl = v2
		}

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
			AltTitles:              v.AltTitles,     
			ExternalIDs:            v.ExternalIDs,     
			CoverUrl:			    coverUrl,  
		})
	}

	return dto.Paginated[dto.AnimeData]{
		Data: animeData,
		Meta:opts.Pagination.BuildMeta(total),
	}, nil
}

func (as AnimeService) Show (ctx context.Context, param string) (dto.AnimeShowData, error) {
	exist, err := func() (domain.Anime, error) {
		if utility.IsUUID(param) {
			return as.animeRepository.FindById(ctx, param)
		}
		return as.animeRepository.FindBySlug(ctx, param)
	}()


    if err != nil && exist.Id == "" {
        return dto.AnimeShowData{}, utility.NewNotFound("Anime")
    }
    
    if err != nil {
        return dto.AnimeShowData{}, err
    }

	episodes, err := as.animeEpisodeRepository.FindByAnimeId(ctx, exist.Id)
	
	if err != nil {
		return dto.AnimeShowData{}, err
	}

	episodesData := make([]dto.AnimeEpisodeData, 0, len(episodes))
	for _, v := range episodes {
		episodesData = append(episodesData, dto.AnimeEpisodeData{
			Id:              v.Id,
			// AnimeId:         v.AnimeId,
			Number:          v.Number,
			SeasonNumber:    v.SeasonNumber,
			Title:           v.Title,
			Synopsis:        v.Synopsis,
			AirDate:         utility.ToTimePtr(v.AirDate),
			DurationMinutes: v.DurationMinutes,
			IsSpecial:       v.IsSpecial,
		})
	}

	genres, err := as.animeGenreRepository.FindByAnimeId(ctx, exist.Id)
	if err != nil { 
		return dto.AnimeShowData{}, err 
	}

	genresData := make([]dto.AnimeGenreData,0, len(genres))
	for _, g := range genres {
		genresData = append(genresData, dto.AnimeGenreData{
			Id:g.Id,
			Slug:g.Slug,			
			Name:g.Name,
		})
	} 

	tags, err := as.animeTagRepository.FindByAnimeId(ctx, exist.Id)
	if err != nil { 
		return dto.AnimeShowData{}, err 
	}
	
	tagsData := make([]dto.AnimeTagData,0,len(tags))
	for _, t := range tags {
		tagsData = append(tagsData, dto.AnimeTagData{
			Id:t.Id,
			Slug:t.Slug,			
			Name:t.Name,
		})
	} 

	studios, err := as.animeStudioRepository.FindByAnimeId(ctx, exist.Id)
	if err != nil { 
		return dto.AnimeShowData{}, err 
	}
	
	studiosData := make([]dto.AnimeStudioData,0,len(studios))
	for _, s := range studios {
		studiosData = append(studiosData, dto.AnimeStudioData{
			Id:s.Id,
			Slug:s.Slug,			
			Name:s.Name,
			Country:s.Country,
			SiteURL:s.SiteURL,
		})
	}

	voiceCasts, err := as.voiceCastRepository.FindByAnimeId(ctx,exist.Id)

	if err != nil {
		return dto.AnimeShowData{}, nil
	}

	voiceCastData := make([]dto.VoiceCastShowData, 0, len(voiceCasts))
	for _, vc := range voiceCasts {
		voiceCastData = append(voiceCastData, dto.VoiceCastShowData{
			VoiceCastData: dto.VoiceCastData{
				Id:          vc.Id,
				Language:    utility.ToStringPtr(vc.Language),  // atau utility.ToString(*vc.Language)
				RoleNote:    utility.ToStringPtr(vc.Language),
			},
			Character: dto.CharacterData{
				Id:   vc.CharacterId,
				Slug: vc.CharacterSlug,
				Name: vc.CharacterName,
				NameNative:vc.CharacterNameNative,
				Description:vc.CharacterDescription,
				// isi field lain kalau CharacterData kamu punya field tambahan
			},
			People: dto.PeopleData{
				Id:   vc.PersonId,
				Slug: vc.PeopleSlug,
				Name: vc.PeopleName,
				NameNative:vc.PeopleNameNative,
				Birthday:utility.ToTimePtr(vc.PeopleBirthday),
				Gender:vc.PeopleGender,
				Country:vc.PeopleCountry,
				SiteURL:vc.PeopleSiteURL,
				Biography:vc.PeopleBiography,
				// isi field lain kalau PeopleData kamu punya field tambahan
			},
		})
	}


	var coverUrl string

	if exist.CoverId.Valid {
		cover, _ := as.mediaRepository.FindById(ctx, exist.CoverId.String)

		if cover.Path != "" {
			coverUrl = as.config.Server.AssetPrivate+"/"+cover.Id
		}
	}

    return dto.AnimeShowData{
		AnimeData:dto.AnimeData{
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
			AltTitles:              exist.AltTitles,     
			ExternalIDs:            exist.ExternalIDs, 
			CoverUrl:               coverUrl,  
		},
		Episodes:episodesData,
		Genres:genresData,
		Tags:tagsData,
		Studios:studiosData,
		VoiceCast: voiceCastData,
    }, nil
}

func (as AnimeService) Create(ctx context.Context, req dto.CreateAnimeRequest) error {
	animeSlug := req.Slug

    if animeSlug == "" {
        animeSlug = slug.Make(req.TitleRomaji) 
		fmt.Println(animeSlug)
    }

	coverId := sql.NullString{String:req.CoverId, Valid:false}

	if req.CoverId != "" {
		coverId.Valid = true 
		
		cover, err := as.mediaRepository.FindById(ctx, req.CoverId)

		if err != nil && cover.Id == "" {
       		return utility.NewNotFound("Anime Media")
		} 
	}											
	
 	anime := domain.Anime{
        Id: uuid.New().String(),
		Slug:animeSlug,
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
		AltTitles: req.AltTitles,
		ExternalIDs: req.ExternalIDs,
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
    }

	return as.animeRepository.Save(ctx, &anime)
}

func (as AnimeService) Update(ctx context.Context, req dto.UpdateAnimeRequest)  error {
    exist, err := as.animeRepository.FindById(ctx, req.Id)

    if err != nil && exist.Id == "" {
        return utility.NewNotFound("Anime")
    }
    
    if err != nil {
        return  err
    }

	animeSlug := req.Slug

    if animeSlug == "" {
        animeSlug = slug.Make(req.TitleRomaji) 
		fmt.Println(animeSlug)
    }

	coverId := sql.NullString{String:req.CoverId, Valid:false}


	if req.CoverId != "" {
		coverId.Valid = true 

		cover, err := as.mediaRepository.FindById(ctx, req.CoverId)

		if err != nil && cover.Id == "" {
       		return utility.NewNotFound("Anime Media")
		} 
	}			

    // Update data sesuai request
	exist.Slug = animeSlug
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
	exist.AltTitles = req.AltTitles
	exist.ExternalIDs = req.ExternalIDs
	exist.UpdatedAt = sql.NullTime{Time: time.Now(), Valid: true}
	exist.CoverId = coverId

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
        return  utility.NewNotFound("Anime")
    }
    
    if err != nil {
        return err
    }

    return as.animeRepository.Delete(ctx, id)
}