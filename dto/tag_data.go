package dto

type AnimeTagData struct {
	Id    string  `json:"id"`
	Slug  string  `json:"slug"`
	Name  string  `json:"name"`
}

type CreateAnimeTagRequest struct {
	Id       string  `json:"id"`
	Slug     string  `json:"slug,omitempty" validate:"omitempty"`
	Name     string  `json:"name"`                              
}

type UpdateAnimeTagRequest struct {
	Id       string  `json:"id"`
	Slug     string  `json:"slug,omitempty" validate:"omitempty"`
	Name     string  `json:"name"`                              
}
