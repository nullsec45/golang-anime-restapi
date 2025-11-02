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

type CharacterService struct {
	characterRepository domain.CharacterRepository
}

func NewCharacter(
	characterRepository domain.CharacterRepository,
) domain.CharacterService {
	return &CharacterService{
		characterRepository: characterRepository,
	}
}

func (cs CharacterService) Index(ctx context.Context, opts domain.CharacterListOptions) (dto.Paginated[dto.CharacterData], error) {
	items, total, err := cs.characterRepository.FindAll(ctx, opts)

	if err != nil {
		return dto.Paginated[dto.CharacterData]{}, err
	}

	var characterData []dto.CharacterData

	for _, v:= range items {

		characterData = append(characterData, dto.CharacterData{
			Id:                     v.Id,
			Slug:                   v.Slug,
			NameNative:             v.NameNative,
			Name:                   v.Name,       
			Description:            v.Description,        
		})
	}

	return dto.Paginated[dto.CharacterData]{
		Data: characterData,
		Meta:opts.Pagination.BuildMeta(total),
	}, nil
}

func (cs CharacterService) Show (ctx context.Context, param string) (dto.CharacterData, error) {
	exist, err := func() (domain.Character, error) {
		if utility.IsUUID(param) {
			return cs.characterRepository.FindById(ctx, param)
		}
		return cs.characterRepository.FindBySlug(ctx, param)
	}()


    if err != nil && exist.Id == "" {
        return dto.CharacterData{}, domain.CharacterNotFound
    }

    return dto.CharacterData{
		Id:                     exist.Id,
		Slug:                   exist.Slug,
		Name:                   exist.Name,       
		NameNative:             exist.NameNative,
		Description:            exist.Description,    
    }, nil
}

func (cs CharacterService) Create(ctx context.Context, req dto.CreateCharacterRequest) error {
	characterSlug := req.Slug

    if characterSlug == "" {
        characterSlug = slug.Make(req.Name) 
    }

 	character := domain.Character{
		Id:uuid.New().String(),
		Slug: characterSlug,
		Name: req.Name,       
		NameNative: req.NameNative,
		Description:req.Description,
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
    }

	return cs.characterRepository.Save(ctx, &character)
}

func (cs CharacterService) Update(ctx context.Context, req dto.UpdateCharacterRequest)  error {
    exist, err := cs.characterRepository.FindById(ctx, req.Id)

    if err != nil && exist.Id == "" {
        return domain.CharacterNotFound
    }
    
    if err != nil {
        return  err
    }

	peopleSlug := req.Slug

    if peopleSlug == "" {
        peopleSlug = slug.Make(req.Name) 
    }


	exist.Slug = peopleSlug
	exist.NameNative=req.NameNative
	exist.Name= req.Name 
	exist.Description=req.Description    

	exist.UpdatedAt = sql.NullTime{Time: time.Now(), Valid: true}

	return cs.characterRepository.Update(ctx, &exist)
}

func (cs CharacterService) Delete (ctx context.Context, id string) error {
    exist, err := cs.characterRepository.FindById(ctx, id)

    if err != nil && exist.Id == "" {
        return  domain.CharacterNotFound
    }
    
    if err != nil {
        return err
    }

    return cs.characterRepository.Delete(ctx, id)
}