package country

import (
	"context"

	"github.com/codesayhi/golang-clean/internal/domain/shared"
)

type ListFilter struct {
	FilterBasic shared.FilterBasic
}

type Repository interface {
	FindSimilarSlugs(ctx context.Context, base string, ignoreID *string) ([]string, error)
	Create(ctx context.Context, c *Country) error
	Update(ctx context.Context, c *Country) error
	Delete(ctx context.Context, id string) error
	FindByID(ctx context.Context, id string) (*Country, error)
	FindBySlug(ctx context.Context, slug string) (*Country, error)
	List(ctx context.Context, filter ListFilter) (items []*Country, total int64, err error)
}
