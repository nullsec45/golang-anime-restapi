package dto

type AnimeStudiosData struct {
	Id        string  `json:"id"`
	AnimeId   string  `json:"anime_id"`
	StudioId  string  `json:"studio_id"`
	Role      string  `json:"role"`
}

type CreateAnimeStudiosRequest struct {
	Id       string `json:"id"`
	AnimeId  string `json:"anime_id" validate:"required,uuid4"`
	StudioId string `json:"studio_id" validate:"required,uuid4"`
	Role     string `json:"role" validate:"required"`
}


type UpdateAnimeStudiosRequest struct {
	Id       string `json:"id"`
	AnimeId  string `json:"anime_id" validate:"required,uuid4"`
	StudioId string `json:"studio_id" validate:"required,uuid4"`
	Role     string `json:"role" validate:"required"`
}