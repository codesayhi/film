package httpcountry

import (
	domain "github.com/codesayhi/golang-clean/internal/domain/country"
)

type CountryResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Slug     string `json:"slug"`
	Code     string `json:"code"`
	Position int    `json:"position"`
}

type ListCountriesResponse struct {
	Items []CountryResponse `json:"items"`
	Total int64             `json:"total"`
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

func toListCountriesResponse(items []*domain.Country, total int64) ListCountriesResponse {
	out := ListCountriesResponse{
		Items: make([]CountryResponse, 0, len(items)),
		Total: total,
	}
	for _, c := range items {
		out.Items = append(out.Items, toCountryResponse(c))
	}
	return out
}
