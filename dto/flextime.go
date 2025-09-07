package dto

import (
	"bytes"
	"encoding/json"
	"reflect"
	"strings"
	"time"
)

type FlexibleTime struct{ time.Time }

var ftLayouts = []string{
	time.RFC3339,
	"2006-01-02",
	"2006-01-02 15:04:05",
}

func (ft *FlexibleTime) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, []byte("null")) {
		*ft = FlexibleTime{}
		return nil
	}
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	s = strings.TrimSpace(s)
	if s == "" {
		*ft = FlexibleTime{}
		return nil
	}
	for _, layout := range ftLayouts {
		if t, err := time.Parse(layout, s); err == nil {
			ft.Time = t.UTC()
			return nil
		}
	}
	return &json.UnmarshalTypeError{Value: "string (bad date format)", Type: reflect.TypeOf(time.Time{})}
}

func (ft FlexibleTime) MarshalJSON() ([]byte, error) {
	if ft.Time.IsZero() {
		return []byte("null"), nil
	}
	return json.Marshal(ft.Time.UTC().Format(time.RFC3339))
}
