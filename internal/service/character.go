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
	"github.com/nullsec45/golang-anime-restapi/internal/config"
	// "fmt"
)

type CharacterService struct {
	characterRepository domain.CharacterRepository
	mediaRepository domain.MediaRepository
	config *config.Config
}

func NewCharacter(
	characterRepository domain.CharacterRepository,
	mediaRepository  domain.MediaRepository,
	config *config.Config) domain.CharacterService {
	return &CharacterService{
		characterRepository: characterRepository,
		mediaRepository:mediaRepository,
		config:config,
	}
}

func (cs CharacterService) Index(ctx context.Context, opts domain.CharacterListOptions) (dto.Paginated[dto.CharacterData], error) {
	items, total, err := cs.characterRepository.FindAll(ctx, opts)

	if err != nil {
		return dto.Paginated[dto.CharacterData]{}, err
	}

	var characterData []dto.CharacterData


	for _, v:= range items {
		characterImage := ""
		if v.CharacterImage.Valid {
			characterImage = cs.config.Server.Asset + "/" + v.CharacterImage.String
		}

		characterData = append(characterData, dto.CharacterData{
			Id:             v.Id,
			Slug:           v.Slug,
			NameNative:     v.NameNative,
			Name:           v.Name,
			Description:    v.Description,
			CharacterImage: characterImage,
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

	if err != nil {
		return dto.CharacterData{}, err
	}

    if err != nil && exist.Id == "" {
        return dto.CharacterData{}, utility.NewNotFound("Character")
    }

	characterImage := ""
	if exist.CharacterImage.Valid {
		characterImage = cs.config.Server.Asset + "/" + exist.CharacterImage.String
	}

    return dto.CharacterData{
		Id:              exist.Id,
		Slug:            exist.Slug,
		Name:            exist.Name,       
		NameNative:      exist.NameNative,
		Description:     exist.Description,    
		CharacterImage:  characterImage,
    }, nil
}

func (cs CharacterService) Create(ctx context.Context, req dto.CreateCharacterRequest) error {
	characterSlug := req.Slug

    if characterSlug == "" {
        characterSlug = slug.Make(req.Name) 
    }

	image := sql.NullString{String:req.CharacterImage, Valid:false}

	if req.CharacterImage != "" {
		image.Valid = true 

		media, err := cs.mediaRepository.FindById(ctx, req.CharacterImage)

		if err != nil && media.Id == "" {
       		return utility.NewNotFound("Image Character")
		} 
	}

 	character := domain.Character{
		Id:uuid.New().String(),
		Slug: characterSlug,
		Name: req.Name,       
		NameNative: req.NameNative,
		Description:req.Description,
		CharacterImage:image,
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
    }

	return cs.characterRepository.Save(ctx, &character)
}

func (cs CharacterService) Update(ctx context.Context, req dto.UpdateCharacterRequest)  error {
    exist, err := cs.characterRepository.FindById(ctx, req.Id)

    if err != nil && exist.Id == "" {
        return utility.NewNotFound("Character")
    }
    
    if err != nil {
        return  err
    }

	characterSlug := req.Slug

    if characterSlug == "" {
        characterSlug = slug.Make(req.Name) 
    }


	image := sql.NullString{String:req.CharacterImage, Valid:false}

	if req.CharacterImage != "" {
		image.Valid = true 

		media, err := cs.mediaRepository.FindById(ctx, req.CharacterImage)

		if err != nil && media.Id == "" {
       		return utility.NewNotFound("Image Character")
		} 
	}

	exist.Slug = characterSlug
	exist.NameNative=req.NameNative
	exist.Name= req.Name 
	exist.Description=req.Description
	exist.CharacterImage=image    

	exist.UpdatedAt = sql.NullTime{Time: time.Now(), Valid: true}

	return cs.characterRepository.Update(ctx, &exist)
}


func (cs CharacterService) Delete (ctx context.Context, id string) error {
    exist, err := cs.characterRepository.FindById(ctx, id)

    if err != nil && exist.Id == "" {
        return  utility.NewNotFound("Character")
    }
    
    if err != nil {
        return err
    }

    return cs.characterRepository.Delete(ctx, id)
}