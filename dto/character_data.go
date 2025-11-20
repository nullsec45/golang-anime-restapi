package dto

type CharacterData struct {
	Id             string `json:"id"`
	Slug           string `json:"slug"`
	Name           string `json:"name"`
	NameNative     string `json:"name_native"`
	Description    string `json:"description"`
	CharacterImage string `json:"character_image"`
}

type CreateCharacterRequest struct {
	Id             string `json:"id"`
	Slug           string `json:"slug" validation:"required"`
	Name           string `json:"name" validation:"required"`
	NameNative     string `json:"name_native" validation:"omitempty"`
	Description    string `json:"description" validation:"omitempty"`
	CharacterImage string `json:"character_image,omitempty"`
}

type UpdateCharacterRequest struct {
	Id             string `json:"id"`
	Slug           string `json:"slug"`
	Name           string `json:"name"`
	NameNative     string `json:"name_native" validation:"omitempty"`
	Description    string `json:"description" validation:"omitempty"`
	CharacterImage string `json:"character_image,omitempty"`
}