package country

import (
	"github.com/codesayhi/golang-clean/internal/domain/country"
	"github.com/codesayhi/golang-clean/pkg/utils"
)

// CreateCountryInput: dùng cho tạo mới.
type CreateCountryInput struct {
	Name     string
	Slug     string
	Code     string
	Position int
}

// UpdateCountryInput: dùng cho PATCH, hỗ trợ Null[T].
type UpdateCountryInput struct {
	ID       string
	Name     utils.Null[string]
	Slug     utils.Null[string]
	Code     utils.Null[string]
	Position utils.Null[int]
}

type ListCountriesInput struct {
	Search  string
	Page    int
	PerPage int
}

type ListCountriesOutput struct {
	Items []*country.Country
	Total int64
}
