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

type VoiceCastService struct {
	animeRepository domain.AnimeRepository
	characterRepository domain.CharacterRepository
	peopleRepository domain.PeopleRepository
	voiceCastRepository domain.VoiceCastRepository
}

func NewVoiceCast(
	animeRepository domain.AnimeRepository,
	characterRepository domain.CharacterRepository,
	peopleRepository domain.PeopleRepository,
	voiceCastRepository domain.VoiceCastRepository) *VoiceCastService {
	return &VoiceCastService{
		animeRepository: animeRepository,
		characterRepository: characterRepository,
		peopleRepository: peopleRepository,
		voiceCastRepository: voiceCastRepository,
	}
}

func (vocs VoiceCastService) Create(ctx context.Context, req dto.CreateVoiceCastRequest) error {
	anime, errAnime := vocs.animeRepository.FindById(ctx, req.AnimeId)

	if anime.Id == "" {
		return utility.NewNotFound("Anime")
	}

	if errAnime != nil {
		return errAnime
	}

	character, errCharacter := vocs.characterRepository.FindById(ctx, req.CharacterId)

	if character.Id == "" {
		return utility.NewNotFound("Character")
	}

	if errCharacter != nil {
		return errCharacter
	}

	people, errPeople := vocs.peopleRepository.FindById(ctx, req.PersonId)

	if people.Id == "" {
		return utility.NewNotFound("People")
	}

	if errPeople != nil {
		return errPeople
	}

	_,  found, errVoiceCast := vocs.voiceCastRepository.FindUnique(ctx, req.AnimeId, req.CharacterId, req.PersonId)

	if errVoiceCast != nil {
		return errVoiceCast
	}

	if found {
		return utility.NewAlreadyExist("Voice Cast")
	}	
	
 	vc := domain.VoiceCast{
		Id:        uuid.NewString(),
		AnimeId:   req.AnimeId,
		CharacterId:     req.CharacterId,
		PersonId:req.PersonId,
		Language:req.Language,
		RoleNote:req.RoleNote,
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
	}

	return vocs.voiceCastRepository.Save(ctx, &vc)
}

func (vocs VoiceCastService) Update(ctx context.Context, req dto.UpdateVoiceCastRequest) error {
	
	exist, err := vocs.voiceCastRepository.FindById(ctx, req.Id)

	if exist.Id == "" {
		return utility.NewNotFound("Voice Cast")
	}

	if err != nil {
		return err
	}

	anime, errAnime := vocs.animeRepository.FindById(ctx, req.AnimeId)

	if anime.Id == "" {
		return utility.NewNotFound("Anime")
	}

	if errAnime != nil {
		return errAnime
	}

	character, errCharacter := vocs.characterRepository.FindById(ctx, req.CharacterId)

	if character.Id == "" {
		return utility.NewNotFound("Character")
	}

	if errCharacter != nil {
		return errCharacter
	}

	people, err := vocs.peopleRepository.FindById(ctx, req.PersonId)

    if err != nil && people.Id == "" {
        return utility.NewNotFound("People")
    }
    
    if err != nil {
        return  err
    }

	// _, found, errVoiceCast := vocs.voiceCastRepository.FindUnique(ctx, req.AnimeId, req.CharacterId, req.PersonId)

	// if errVoiceCast != nil {
	// 	return errVoiceCast
	// }

	// if found {
	// 	return utility.NewAlreadyExist("Voice Cast")
	// }

	exist.AnimeId = req.AnimeId
	exist.CharacterId = req.CharacterId
	exist.PersonId=req.PersonId
	exist.Language=req.Language
	exist.RoleNote=req.RoleNote
	exist.UpdatedAt = sql.NullTime{Time: time.Now(), Valid: true}

	return vocs.voiceCastRepository.Update(ctx, &exist)
}


// func (vocs VoiceCastService) DeleteByAnimeId (ctx context.Context, animeId string) error {
//     exist, err := vocs.animeRepository.FindById(ctx, animeId)

//     if err != nil && exist.Id == "" {
//         return  utility.NewNotFound("Anime")
//     }
    
//     if err != nil {
//         return err
//     }

//     return vocs.peopleRepository.DeleteByAnimeId(ctx, animeId)
// }

// func (vocs VoiceCastService) DeleteByTagId (ctx context.Context, tagId string) error {
//     exist, err := vocs.characterRepository.FindById(ctx, tagId)

//     if err != nil && exist.Id == "" {
//         return  utility.NewNotFound("Anime Tag")
//     }
    
//     if err != nil {
//         return err
//     }

//     return vocs.peopleRepository.DeleteByTagId(ctx, tagId)
// }

func (vocs VoiceCastService) DeleteById (ctx context.Context, Id string) error {
    exist, err := vocs.voiceCastRepository.FindById(ctx, Id)

    if err != nil && exist.Id == "" {
        return  utility.NewNotFound("Voice Cast")
    }
    
    if err != nil {
        return err
    }

    return vocs.voiceCastRepository.DeleteById(ctx, Id)
}