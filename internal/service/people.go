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
	"fmt"
)

type PeopleService struct {
	peopleRepository domain.PeopleRepository
	mediaRepository  domain.MediaRepository
	config *config.Config
}

func NewPeople(
	peopleRepository domain.PeopleRepository,
	mediaRepository domain.MediaRepository,
	config *config.Config,
) domain.PeopleService {
	return &PeopleService{
		peopleRepository: peopleRepository,
		mediaRepository:mediaRepository,
		config:config,
	}
}

func (ps PeopleService) Index(ctx context.Context, opts domain.PeopleListOptions) (dto.Paginated[dto.PeopleData], error) {
	items, total, err := ps.peopleRepository.FindAll(ctx, opts)

	if err != nil {
		return dto.Paginated[dto.PeopleData]{}, err
	}

	var peopleData []dto.PeopleData

	for _, v:= range items {
		personImage := ""
		if v.PersonImage.Valid {
			personImage = ps.config.Server.Asset + "/" + v.PersonImage.String
		}

		peopleData = append(peopleData, dto.PeopleData{
			Id:         v.Id,
			Slug:       v.Slug,
			NameNative: v.NameNative,
			Name:       v.Name,       
			Birthday:   utility.ToTimePtr(v.Birthday),        
			Gender:     dto.GenderType(v.Gender),            	
			Country:    v.Country,
			SiteURL:    v.SiteURL,         
			Biography:  v.Biography,    
			PersonImage: personImage,
		})
	}

	return dto.Paginated[dto.PeopleData]{
		Data: peopleData,
		Meta:opts.Pagination.BuildMeta(total),
	}, nil
}

func (ps PeopleService) Show (ctx context.Context, param string) (dto.PeopleData, error) {
	exist, err := func() (domain.People, error) {
		if utility.IsUUID(param) {
			return ps.peopleRepository.FindById(ctx, param)
		}
		return ps.peopleRepository.FindBySlug(ctx, param)
	}()

	if err != nil {
		return dto.PeopleData{}, err
	}

    if err != nil && exist.Id == "" {
        return dto.PeopleData{}, utility.NewNotFound("People")
    }

	personImage := ""
	if exist.PersonImage.Valid {
		personImage = ps.config.Server.Asset + "/" + exist.PersonImage.String
	}

    return dto.PeopleData{
		Id:          exist.Id,
		Slug:        exist.Slug,
		NameNative:  exist.NameNative,
		Name:        exist.Name,       
		Birthday:    utility.ToTimePtr(exist.Birthday),        
		Gender:      dto.GenderType(exist.Gender),            	
		Country:     exist.Country,
		SiteURL:     exist.SiteURL,         
		Biography:   exist.Biography,  
		PersonImage: personImage,  
    }, nil
}

func (ps PeopleService) Create(ctx context.Context, req dto.CreatePeopleRequest) error {
	peopleSlug := req.Slug

    if peopleSlug == "" {
        peopleSlug = slug.Make(req.Name) 
    }

	image := sql.NullString{String:req.PersonImage, Valid:false}

	if req.PersonImage != "" {
		image.Valid = true 

		media, err := ps.mediaRepository.FindById(ctx, req.PersonImage)

		if err != nil && media.Id == "" {
       		return utility.NewNotFound("Image Character")
		} 
	}

 	people := domain.People{
		Id:uuid.New().String(),
		Slug: peopleSlug,
		NameNative: req.NameNative,
		Name: req.Name,       
		Birthday: utility.ToSqlNullTime(req.Birthday),        
		Gender: utility.ToGenderType(req.Gender),            	
		Country: req.Country,
		SiteURL: req.SiteURL,         
		Biography: req.Biography,   
		PersonImage:image, 
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
    }

	return ps.peopleRepository.Save(ctx, &people)
}

// func (ps PeopleService) Update(ctx context.Context, req dto.UpdatePeopleRequest)  error {
//     exist, err := ps.peopleRepository.FindById(ctx, req.Id)

//     if err != nil && exist.Id == "" {
//         return utility.NewNotFound("People")
//     }
    
//     if err != nil {
//         return  err
//     }

// 	peopleSlug := req.Slug

//     if peopleSlug == "" {
//         peopleSlug = slug.Make(req.Name) 
//     }


// 	image := sql.NullString{String:req.PersonImage, Valid:false}

// 	if req.PersonImage != "" {
// 		image.Valid = true 

// 		media, err := ps.mediaRepository.FindById(ctx, req.PersonImage)

// 		if err != nil && media.Id == "" {
//        		return utility.NewNotFound("Image Character")
// 		} 
// 	}


// 	exist.Slug = peopleSlug
// 	exist.NameNative=req.NameNative
// 	exist.Name= req.Name
// 	exist.Birthday=utility.ToSqlNullTime(req.Birthday)        
// 	exist.Gender=utility.ToGenderType(req.Gender)            	
// 	exist.Country=req.Country
// 	exist.SiteURL=req.SiteURL         
// 	exist.Biography=req.Biography    
// 	exist.PersonImage=image

// 	exist.UpdatedAt = sql.NullTime{Time: time.Now(), Valid: true}

// 	return ps.peopleRepository.Update(ctx, &exist)
// }

func (ps PeopleService) Update(ctx context.Context, req dto.UpdatePeopleRequest) error {
    exist, err := ps.peopleRepository.FindById(ctx, req.Id)
    if err != nil || exist.Id == "" {
        return utility.NewNotFound("People")
    }

    peopleSlug := req.Slug
    if peopleSlug == "" {
        peopleSlug = slug.Make(req.Name)
    }

    if peopleSlug != exist.Slug {
		fmt.Println(req.Id)
		fmt.Println(peopleSlug)
		fmt.Println(exist.Slug)
        other, _ := ps.peopleRepository.FindBySlug(ctx, peopleSlug)
        if other.Id != "" && other.Id != exist.Id {
            return utility.NewAlreadyExist("Slug")
        }
    }

    // --- Cek unik name ---
    if req.Name != exist.Name {
		fmt.Println(req.Name)
		fmt.Println(exist.Name)
        other, _ := ps.peopleRepository.FindByName(ctx, req.Name)
        if other.Id != "" && other.Id != exist.Id {
            return utility.NewAlreadyExist("Name")
        }
    }

    image := sql.NullString{String: req.PersonImage, Valid: req.PersonImage != ""}

    if image.Valid {
        media, err := ps.mediaRepository.FindById(ctx, req.PersonImage)
        if err != nil || media.Id == "" {
            return utility.NewNotFound("Image Character")
        }
    }

    exist.Slug        = peopleSlug
    exist.NameNative  = req.NameNative
    exist.Name        = req.Name
    exist.Birthday    = utility.ToSqlNullTime(req.Birthday)
    exist.Gender      = utility.ToGenderType(req.Gender)
    exist.Country     = req.Country
    exist.SiteURL     = req.SiteURL
    exist.Biography   = req.Biography
    exist.PersonImage = image
    exist.UpdatedAt   = sql.NullTime{Time: time.Now(), Valid: true}

    return ps.peopleRepository.Update(ctx, &exist)
}

func (ps PeopleService) Delete (ctx context.Context, id string) error {
    exist, err := ps.peopleRepository.FindById(ctx, id)

    if err != nil && exist.Id == "" {
        return  utility.NewNotFound("People")
    }
    
    if err != nil {
        return err
    }

    return ps.peopleRepository.Delete(ctx, id)
}