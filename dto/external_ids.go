// dto/external_ids.go
package dto

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type ExternalIDs struct {
	MyAnimeList int    `json:"myanimelist,omitempty"`
	AniList     int    `json:"anilist,omitempty"`
	AniDB       int    `json:"anidb,omitempty"`
	IMDB        string `json:"imdb,omitempty"`
}

func (e ExternalIDs) Value() (driver.Value, error) {
	b, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (e *ExternalIDs) Scan(src any) error {
	switch v := src.(type) {
	case []byte:
		return json.Unmarshal(v, e)
	case string:
		return json.Unmarshal([]byte(v), e)
	case nil:
		*e = ExternalIDs{}
		return nil
	default:
		return fmt.Errorf("ExternalIDs.Scan: unsupported type %T", src)
	}
}
