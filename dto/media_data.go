package dto

type MediaData struct {
	Id    string  `json:"id"`
	Path  string  `json:"path"`
	Url   string  `json:"url"`
	OldPath string `json:"-"`
}

type CreateMediaRequest struct {
	Path string `json:"path" validate:"required"`
}

type UpdateMediaRequest struct {
	Id   string `json:"id"`
	Path string `json:"path" validate:"required"`
}