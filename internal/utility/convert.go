package utility

import (
	"github.com/nullsec45/golang-anime-restapi/dto"
	"database/sql"
	"encoding/json"
	"time"
)

func toSeasonPtr(s *string) *dto.Season {
	if s == nil {
		return nil
	}
	v := dto.Season(*s) 
	return &v
}

func toAgeRatingPtr(s *string) *dto.AgeRating {
	if s == nil {
		return nil
	}
	v := dto.AgeRating(*s) 
	return &v
}

func toJSONMap[T any](v T) map[string]any {
	switch x := any(v).(type) {
	case nil:
		return map[string]any{}
	case map[string]any:
		return x
	case []byte:
		var m map[string]any
		_ = json.Unmarshal(x, &m)
		return m
	case json.RawMessage:
		var m map[string]any
		_ = json.Unmarshal(x, &m)
		return m
	default:
		return map[string]any{}
	}
}

// *string -> sql.NullString
func nstr(p *string) sql.NullString {
	if p == nil {
		return sql.NullString{}
	}
	return sql.NullString{String: *p, Valid: true}
}

// *int -> sql.NullInt32
func nint(p *int) sql.NullInt32 {
	if p == nil {
		return sql.NullInt32{}
	}
	return sql.NullInt32{Int32: int32(*p), Valid: true}
}

// *int16 -> sql.NullInt16
func nint16(p *int16) sql.NullInt16 {
	if p == nil {
		return sql.NullInt16{}
	}
	return sql.NullInt16{Int16: *p, Valid: true}
}

// *time.Time -> sql.NullTime
func ntime(p *time.Time) sql.NullTime {
	if p == nil {
		return sql.NullTime{}
	}
	return sql.NullTime{Time: *p, Valid: true}
}

// *float32 -> sql.NullFloat64 (cast ke float64)
func nfloat32(p *float32) sql.NullFloat64 {
	if p == nil {
		return sql.NullFloat64{}
	}
	return sql.NullFloat64{Float64: float64(*p), Valid: true}
}

// map[string]any -> json.RawMessage (untuk JSONB)
func njson(m map[string]any) json.RawMessage {
	if len(m) == 0 {
		return nil
	}
	b, _ := json.Marshal(m)
	return json.RawMessage(b)
}

// enum pointer -> *string
func seasonToString(s *dto.Season) *string {
	if s == nil {
		return nil
	}
	v := string(*s) // "Winter|Spring|Summer|Fall"
	return &v
}

func ageToString(a *dto.AgeRating) *string {
	if a == nil {
		return nil
	}
	v := string(*a) // "G|PG|PG-13|R|R+|Rx"
	return &v
}

// utilities kecil
func firstNonEmpty(vals ...string) string {
	for _, v := range vals {
		if v != "" {
			return v
		}
	}
	return ""
}
func ptrToString(p *string) string {
	if p == nil {
		return ""
	}
	return *p
}
func ptrToInt(p *int) int {
	if p == nil {
		return 0
	}
	return *p
}