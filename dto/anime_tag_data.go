package dto

type AnimeTagsData struct {
	Id       string  `json:"id"`
	AnimeId  string  `json:"anime_id"`
	TagId    string  `json:"tag_id"`
}

type CreateAnimeTagsRequest struct {
	Id       string `json:"id"`
	AnimeId  string `json:"anime_id" validate:"required,uuid4"`
	TagId    string `json:"tag_id" validate:"required,uuid4"`
}


type UpdateAnimeTagsRequest struct {
	Id       string `json:"id"`
	AnimeId  string `json:"anime_id" validate:"required,uuid4"`
	TagId    string `json:"tag_id" validate:"required,uuid4"`
}