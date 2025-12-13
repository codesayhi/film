package httpcountry

import "github.com/codesayhi/golang-clean/pkg/utils"

// CreateCountryRequest dùng cho POST /countries
type CreateCountryRequest struct {
	Name     string `json:"name" binding:"required,min=2,max=255"`
	Code     string `json:"code" binding:"required,min=2,max=20"`
	Position int    `json:"position" binding:"omitempty,gte=0"`
}

// UpdateCountryRequest dùng cho PATCH /countries/:id
// Null[T] sẽ cho phép phân biệt: không gửi / gửi null / gửi giá trị.
type UpdateCountryRequest struct {
	Name     utils.Null[string] `json:"name"`
	Slug     utils.Null[string] `json:"slug"`
	Code     utils.Null[string] `json:"code"`
	Position utils.Null[int]    `json:"position"`
}

// ListCountriesRequest cho GET /countries
type ListCountriesRequest struct {
	Search  string `form:"q"`
	Page    int    `form:"page,default=1" binding:"gte=1"`
	PerPage int    `form:"per_page,default=20" binding:"gte=1,lte=100"`
}
