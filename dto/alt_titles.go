// dto/alt_titles.go
package dto

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type AltTitles struct {
	En      string   `json:"en,omitempty"`
	Ja      string   `json:"ja,omitempty"`
	Romaji  string   `json:"romaji,omitempty"`
	Aliases []string `json:"aliases,omitempty"`
}

func (a AltTitles) Value() (driver.Value, error) {
	b, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (a *AltTitles) Scan(src any) error {
	switch v := src.(type) {
	case []byte:
		return json.Unmarshal(v, a)
	case string:
		return json.Unmarshal([]byte(v), a)
	case nil:
		*a = AltTitles{}
		return nil
	default:
		return fmt.Errorf("AltTitles.Scan: unsupported type %T", src)
	}
}
