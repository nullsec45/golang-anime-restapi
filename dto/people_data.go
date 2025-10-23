package dto 


type GenderType string
const (
	TypeFemale GenderType="Female"
	TypeMale GenderType="Male"
)


type PeopleData struct {
	Id   		string        `json:"id"`
	Slug 		string        `json:"slug"`
	NameNative  string        `json:"name_native"`
	Name        string        `json:"name"`
	Birthday    *FlexibleTime `json:"birthday"`
	Gender      GenderType    `json:"gender"`
	Country     string        `json:"country"`
	SiteURL     string        `json:"site_url"`
	Biography   string        `json:"biography"`
}

type CreatePeopleRequest struct {
	Slug 		string        `json:"slug" validate:"omitempty"`
	NameNative  string        `json:"name_native" validate:"omitempty"`
	Name        string        `json:"name" validate:"required"`
	Birthday    *FlexibleTime `json:"birthday" validate:"required"`
	Gender      GenderType    `json:"gender" validate:"required,oneof=Female Male"`
	Country     string        `json:"country" validate:"required"`
	SiteURL     string        `json:"site_url" validate:"omitempty"`
	Biography   string        `json:"biography" validate:"required"`                         
}

type UpdatePeopleRequest struct {
	Id          string        `json:"id"`
	Slug 		string        `json:"slug" validate:"omitempty"`
	NameNative  string        `json:"name_native" validate:"omitempty"`
	Name        string        `json:"name" validate:"required"`
	Birthday    *FlexibleTime `json:"birthday" validate:"required"`
	Gender      GenderType    `json:"gender" validate:"required,oneof=Female Male"`
	Country     string        `json:"country" validate:"required"`
	SiteURL     string        `json:"site_url" validate:"omitempty"`
	Biography   string        `json:"biography" validate:"required"`                   
}
