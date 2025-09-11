package dto

import (
	"encoding/json"
	"strings"
	"time"
	"fmt"
)

type FlexibleTime struct{ 
	time.Time 
	Layout string

}

var ftLayouts = []string{
	time.RFC3339,
	"2006-01-02",
	"2006-01-02 15:04:05",
}

func (ft FlexibleTime) IsZero() bool { return ft.Time.IsZero() }


func (ft *FlexibleTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	if s == "" || s == "null" {
		*ft = FlexibleTime{}
		return nil
	}

	allowed := []string{
		"2006-01-02",         
		time.RFC3339,         
		"2006-01-02 15:04:05",
	}

	var lastErr error
	for _, layout := range allowed {
		if t, err := time.Parse(layout, s); err == nil {
			*ft = FlexibleTime{Time: t, Layout: layout}
			return nil
		} else {
			lastErr = err
		}
	}
	return fmt.Errorf("invalid date %q: must be YYYY-MM-DD (or one of allowed layouts): %w", s, lastErr)
}


func (ft FlexibleTime) MarshalJSON() ([]byte, error) {
	if ft.Time.IsZero() {
		return []byte("null"), nil
	}
	return json.Marshal(ft.Time.UTC().Format(time.RFC3339))
}
