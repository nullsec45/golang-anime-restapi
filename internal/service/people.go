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

type PeopleService struct {
	peopleRepository domain.PeopleRepository
}

func NewPeople(
	peopleRepository domain.PeopleRepository,
) domain.PeopleService {
	return &PeopleService{
		peopleRepository: peopleRepository,
	}
}

func (ps PeopleService) Index(ctx context.Context, opts domain.PeopleListOptions) (dto.Paginated[dto.PeopleData], error) {
	items, total, err := ps.peopleRepository.FindAll(ctx, opts)

	if err != nil {
		return dto.Paginated[dto.PeopleData]{}, err
	}

	var peopleData []dto.PeopleData

	for _, v:= range items {

		peopleData = append(peopleData, dto.PeopleData{
			Id:                     v.Id,
			Slug:                   v.Slug,
			NameNative:             v.NameNative,
			Name:                   v.Name,       
			Birthday:               utility.ToTimePtr(v.Birthday),        
			Gender:                 dto.GenderType(v.Gender),            	
			Country:                v.Country,
			SiteURL:                v.SiteURL,         
			Biography:              v.Biography,    
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


    if err != nil && exist.Id == "" {
        return dto.PeopleData{}, domain.PeopleNotFound
    }

    return dto.PeopleData{
		Id:                     exist.Id,
		Slug:                   exist.Slug,
		NameNative:             exist.NameNative,
		Name:                   exist.Name,       
		Birthday:               utility.ToTimePtr(exist.Birthday),        
		Gender:                 dto.GenderType(exist.Gender),            	
		Country:                exist.Country,
		SiteURL:                exist.SiteURL,         
		Biography:              exist.Biography,    
    }, nil
}

func (ps PeopleService) Create(ctx context.Context, req dto.CreatePeopleRequest) error {
	peopleSlug := req.Slug

    if peopleSlug == "" {
        peopleSlug = slug.Make(req.Name) 
    }

 	anime := domain.People{
		Id:uuid.New().String(),
		Slug: peopleSlug,
		NameNative: req.NameNative,
		Name: req.Name,       
		Birthday: utility.ToSqlNullTime(req.Birthday),        
		Gender: utility.ToGenderType(req.Gender),            	
		Country: req.Country,
		SiteURL: req.SiteURL,         
		Biography: req.Biography,    
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
    }

	return ps.peopleRepository.Save(ctx, &anime)
}

func (ps PeopleService) Update(ctx context.Context, req dto.UpdatePeopleRequest)  error {
    exist, err := ps.peopleRepository.FindById(ctx, req.Id)

    if err != nil && exist.Id == "" {
        return domain.PeopleNotFound
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
	exist.Birthday=utility.ToSqlNullTime(req.Birthday)        
	exist.Gender=utility.ToGenderType(req.Gender)            	
	exist.Country=req.Country
	exist.SiteURL=req.SiteURL         
	exist.Biography=req.Biography    

	exist.UpdatedAt = sql.NullTime{Time: time.Now(), Valid: true}

	return ps.peopleRepository.Update(ctx, &exist)
}

func (ps PeopleService) Delete (ctx context.Context, id string) error {
    exist, err := ps.peopleRepository.FindById(ctx, id)

    if err != nil && exist.Id == "" {
        return  domain.PeopleNotFound
    }
    
    if err != nil {
        return err
    }

    return ps.peopleRepository.Delete(ctx, id)
}