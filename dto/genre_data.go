package dto 

type AnimeGenreData struct {
	Id    string  `json:"id"`
	Slug  string  `json:"slug"`
	Name  string  `json:"name"`
}

type CreateAnimeGenreRequest struct {
	Id       string  `json:"id"`
	Slug     string  `json:"slug,omitempty" validate:"omitempty"`
	Name     string  `json:"name"`                              
}

type UpdateAnimeGenreRequest struct {
	Id       string  `json:"id"`
	Slug     string  `json:"slug,omitempty" validate:"omitempty"`
	Name     string  `json:"name"`                              
}
