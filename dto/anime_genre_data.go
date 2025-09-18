package dto

type AnimeGenresData struct {
	Id       string  `json:"id"`
	AnimeId  string  `json:"anime_id"`
	GenreId  string  `json:"genre_id"`
}

type CreateAnimeGenresRequest struct {
	Id       string `json:"id"`
	AnimeId  string `json:"anime_id" validate:"required,uuid4"`
	GenreId  string `json:"genre_id" validate:"required,uuid4"`
}


type UpdateAnimeGenresRequest struct {
	Id       string `json:"id"`
	AnimeId  string `json:"anime_id" validate:"required,uuid4"`
	GenreId  string `json:"genre_id" validate:"required,uuid4"`
}