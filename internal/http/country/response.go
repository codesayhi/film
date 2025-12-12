package httpcountry

import (
	domain "github.com/codesayhi/golang-clean/internal/domain/country"
	"github.com/codesayhi/golang-clean/internal/http/shared/pagination"
	usecase "github.com/codesayhi/golang-clean/internal/service/country"
)

type CountryResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Slug     string `json:"slug"`
	Code     string `json:"code"`
	Position int    `json:"position"`
}

type ListCountriesResponse struct {
	Items      []CountryResponse     `json:"items"`
	Pagination pagination.Pagination `json:"pagination"`
}

func toCountryResponse(c *domain.Country) CountryResponse {
	return CountryResponse{
		ID:       c.ID,
		Name:     c.Name,
		Slug:     c.Slug,
		Code:     c.Code,
		Position: c.Position,
	}
}

func toListCountriesResponse(out *usecase.ListCountriesOutput) ListCountriesResponse {
	// map Country → CountryResponse
	items := make([]CountryResponse, 0, len(out.Items))
	for _, c := range out.Items {
		items = append(items, toCountryResponse(c))
	}

	// tạo pagination JSON từ meta
	pg := pagination.New(out.Pagination.Total, out.Pagination.Page, out.Pagination.PerPage)

	return ListCountriesResponse{
		Items:      items,
		Pagination: pg,
	}
}
