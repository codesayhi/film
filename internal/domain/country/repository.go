package country

import "context"

type ListFilter struct {
	Search  string
	Page    int
	PerPage int
}

type Repository interface {
	Create(ctx context.Context, c *Country) error
	Update(ctx context.Context, c *Country) error
	Delete(ctx context.Context, id string) error
	FindByID(ctx context.Context, id string) (*Country, error)
	FindBySlug(ctx context.Context, slug string) (*Country, error)
	List(ctx context.Context, filter ListFilter) (items []*Country, total int64, err error)
}
