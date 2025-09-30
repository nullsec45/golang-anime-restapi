package dto

import "math"

type PaginationQuery struct {
	Page    int    `query:"page" json:"page"`
	PerPage int    `query:"per_page" json:"per_page"`
	Search  string `query:"search" json:"search,omitempty"`
	Sort    string `query:"sort" json:"sort,omitempty"`   // contoh: "title_romaji" / "created_at"
	Order   string `query:"order" json:"order,omitempty"` // "asc" / "desc"
}

func (q *PaginationQuery) Normalize(defaultPage, defaultPerPage, maxPerPage int) {
	if q.Page < 1 {
		q.Page = defaultPage
	}
	if q.PerPage <= 0 {
		q.PerPage = defaultPerPage
	}
	if q.PerPage > maxPerPage {
		q.PerPage = maxPerPage
	}
	if q.Order == "" {
		q.Order = "asc"
	}
}

func (q PaginationQuery) LimitOffset() (limit, offset uint64) {
	return uint64(q.PerPage), uint64((q.Page-1)*q.PerPage)
}

type PageMeta struct {
	Page       int   `json:"page"`
	PerPage    int   `json:"per_page"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
	HasNext    bool  `json:"has_next"`
	HasPrev    bool  `json:"has_prev"`
	NextPage   *int  `json:"next_page,omitempty"`
	PrevPage   *int  `json:"prev_page,omitempty"`
}

func (q PaginationQuery) BuildMeta(total int64) PageMeta {
	totalPages := 0
	if q.PerPage > 0 && total > 0 {
		totalPages = int(math.Ceil(float64(total) / float64(q.PerPage)))
	}
	meta := PageMeta{
		Page:       q.Page,
		PerPage:    q.PerPage,
		Total:      total,
		TotalPages: totalPages,
		HasNext:    q.Page < totalPages,
		HasPrev:    q.Page > 1 && totalPages > 0,
	}
	if meta.HasNext {
		n := q.Page + 1
		meta.NextPage = &n
	}
	if meta.HasPrev {
		p := q.Page - 1
		meta.PrevPage = &p
	}
	return meta
}

type Paginated[T any] struct {
	Data []T     `json:"data"`
	Meta PageMeta `json:"meta"`
}
