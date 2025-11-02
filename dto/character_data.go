package dto

type CharacterData struct {
	Id          string `json:"id"`
	Slug        string `json:"slug"`
	Name        string `json:"name"`
	NameNative  string `json:"name_native"`
	Description string `json:"description"`
}

type CreateCharacterRequest struct {
	Id          string `json:"id"`
	Slug        string `json:"slug"`
	Name        string `json:"name"`
	NameNative  string `json:"name_native"`
	Description string `json:"description"`
}

type UpdateCharacterRequest struct {
	Id          string `json:"id"`
	Slug        string `json:"slug"`
	Name        string `json:"name"`
	NameNative  string `json:"name_native"`
	Description string `json:"description"`
}