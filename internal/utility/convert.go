package utility

import (
	"github.com/nullsec45/golang-anime-restapi/dto"
	"database/sql"
	"encoding/json"
)

func ToString(ns sql.NullString) string {
    if ns.Valid {
        return ns.String
    }
    return ""
}

func ToSeasonPtr(s *string) *dto.Season {
	if s == nil {
		return nil
	}
	v := dto.Season(*s) 
	return &v
}

func ToAgeRatingPtr(s *string) *dto.AgeRating {
	if s == nil {
		return nil
	}
	v := dto.AgeRating(*s) 
	return &v
}



func ToTimePtr(nt sql.NullTime) *dto.FlexibleTime {
	if !nt.Valid {
		return nil
	}
	ft := dto.FlexibleTime{}
	ft.Time = nt.Time.UTC()
	return &ft
}

func ToAnimeType(v any) dto.AnimeType {
	switch x := v.(type) {
	case dto.AnimeType:
		return x
	case *dto.AnimeType:
		if x == nil {
			return ""
		}
		return *x
	case string:
		return dto.AnimeType(x)
	case *string:
		if x == nil {
			return ""
		}
		return dto.AnimeType(*x)
	default:
		return ""
	}
}

func ToAnimeStatus(v any) dto.AnimeStatus {
	switch x := v.(type) {
	case dto.AnimeStatus:
		return x
	case *dto.AnimeStatus:
		if x == nil { return "" }
		return *x
	case string:
		return dto.AnimeStatus(x)
	case *string:
		if x == nil { return "" }
		return dto.AnimeStatus(*x)
	default:
		return ""
	}
}

func ToSeason(v any) *dto.Season {
	switch x := v.(type) {
	case nil:
		return nil
	case dto.Season:
		s := x
		return &s
	case *dto.Season:
		return x
	case string:
		s := dto.Season(x)
		return &s
	case *string:
		if x == nil {
			return nil
		}
		s := dto.Season(*x)
		return &s
	default:
		return nil
	}
}

func ToAgeRating(v any) *dto.AgeRating {
	switch x := v.(type) {
	case nil:
		return nil
	case dto.AgeRating:
		s := x
		return &s
	case *dto.AgeRating:
		return x
	case string:
		s := dto.AgeRating(x)
		return &s
	case *string:
		if x == nil {
			return nil
		}
		s := dto.AgeRating(*x)
		return &s
	default:
		return nil
	}
}

func nstr(p *string) sql.NullString {
	if p == nil {
		return sql.NullString{}
	}
	return sql.NullString{String: *p, Valid: true}
}


func nint16(p *int16) sql.NullInt16 {
	if p == nil {
		return sql.NullInt16{}
	}
	return sql.NullInt16{Int16: *p, Valid: true}
}

func ToSqlNullTime(ft *dto.FlexibleTime) sql.NullTime {
	if ft == nil || ft.Time.IsZero() {
		return sql.NullTime{}
	}
	return sql.NullTime{Time: ft.Time.UTC(), Valid: true}
}

func nfloat32(p *float32) sql.NullFloat64 {
	if p == nil {
		return sql.NullFloat64{}
	}
	return sql.NullFloat64{Float64: float64(*p), Valid: true}
}

func ToJson(m map[string]any) json.RawMessage {
	if len(m) == 0 {
		return nil
	}
	b, _ := json.Marshal(m)
	return json.RawMessage(b)
}

func SeasonToString(s *dto.Season) *string {
	if s == nil {
		return nil
	}
	v := string(*s) 
	return &v
}

func AgeToString(a *dto.AgeRating) *string {
	if a == nil {
		return nil
	}
	v := string(*a)
	return &v
}

func FirstNonEmpty(vals ...string) string {
	for _, v := range vals {
		if v != "" {
			return v
		}
	}
	return ""
}
func PtrToString(p *string) string {
	if p == nil {
		return ""
	}
	return *p
}
func PtrToInt(p *int) int {
	if p == nil {
		return 0
	}
	return *p
}
