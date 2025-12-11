package genre

import "context"

type Filter struct {
	Search  string
	Page    int
	PerPage int
}

type Repository interface {
	Create(ctx context.Context, g *Genre) error
	Update(ctx context.Context, g *Genre) error
	Delete(ctx context.Context, id string) error
	FindByID(ctx context.Context, id string) (*Genre, error)
	FindBySlug(ctx context.Context, slug string) (*Genre, error)
	List(ctx context.Context, filter Filter) ([]*Genre, error)
}
