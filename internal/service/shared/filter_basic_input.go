package sharedservice

import "strings"

const (
	UnsetPage      = 0
	UnsetPerPage   = 0
	DefaultPage    = 1
	DefaultPerPage = 20
	MaxPerPage     = 100
)

type FilterBasicInput struct {
	Search  string
	Page    int
	PerPage int
}

func NewFilterBasicInput(search string, page int, perPage int) FilterBasicInput {
	if page <= UnsetPage {
		page = DefaultPage
	}
	if perPage <= UnsetPerPage {
		perPage = DefaultPerPage
	}
	if perPage > MaxPerPage {
		perPage = MaxPerPage
	}
	return FilterBasicInput{
		Search:  strings.TrimSpace(search),
		Page:    page,
		PerPage: perPage,
	}
}
