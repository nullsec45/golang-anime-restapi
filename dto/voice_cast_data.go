package dto

type VoiceCastData struct {
	Id           string  `json:"id"`
	AnimeId      string  `json:"anime_id"`
	CharacterId  string  `json:"character_id"`
	PersonId     string  `json:"person_id"`
	Language     string  `json:"language`
	RoleNote     string  `json:"role_note"`
}

type CreateVoiceCastRequest struct {
	Id           string  `json:"id"`
	AnimeId      string  `json:"anime_id" validate:"required,uuid4"`
	CharacterId  string  `json:"character_id" validate:"required,uuid4"`
	PersonId     string  `json:"person_id" validate:"required,uuid4"`
	Language     string  `json:"language" validate:"required"`
	RoleNote     string  `json:"role_note"`
}


type UpdateVoiceCastRequest struct {
	Id           string  `json:"id"`
	AnimeId      string  `json:"anime_id" validate:"required,uuid4"`
	CharacterId  string  `json:"character_id" validate:"required,uuid4"`
	PersonId     string  `json:"person_id" validate:"required,uuid4"`
	Language     string  `json:"language" validate:"required"`
	RoleNote     string  `json:"role_note"`
}