package dto 

type AnimeStudioData struct {
	Id    string    `json:"id"`
	Slug  string    `json:"slug"`
	Name  string    `json:"name"`
	Country string  `json:"country"`
	SiteURL string  `json:"site_url"`
}

type CreateAnimeStudioRequest struct {
	Id       string  `json:"id"`
	Slug     string  `json:"slug" validate:"omitempty"`
	Name     string  `json:"name" validate:"required"`
	Country  string  `json:"country" validate:"required"`
	SiteURL  string  `json:"site_url" validate:"required"`                               
}

type UpdateAnimeStudioRequest struct {
	Id       string  `json:"id"`
	Slug     string  `json:"slug" validate:"omitempty"`
	Name     string  `json:"name" validate:"required"`    
	Country  string  `json:"country" validate:"required"`
	SiteURL  string  `json:"site_url" validate:"required"`                           
}
