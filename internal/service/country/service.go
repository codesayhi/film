package country

import (
	"context"
	"strings"

	domain "github.com/codesayhi/golang-clean/internal/domain/country"
	shared2 "github.com/codesayhi/golang-clean/internal/domain/shared"
	"github.com/codesayhi/golang-clean/internal/service/shared"
	"github.com/codesayhi/golang-clean/internal/service/shared/slug"
	"github.com/codesayhi/golang-clean/pkg/utils"
	"github.com/google/uuid"
)

type Service interface {
	Create(ctx context.Context, in CreateCountryInput) (*domain.Country, error)
	GetByID(ctx context.Context, id string) (*domain.Country, error)
	GetBySlug(ctx context.Context, slug string) (*domain.Country, error)
	List(ctx context.Context, in ListCountriesInput) (*ListCountriesOutput, error)
	Update(ctx context.Context, in UpdateCountryInput) (*domain.Country, error)
	Delete(ctx context.Context, id string) error
}

type service struct {
	repo domain.Repository
	slug slug.Generator
}

func NewService(repo domain.Repository, slug slug.Generator) Service {
	return &service{repo: repo, slug: slug}
}

func (s *service) Create(ctx context.Context, in CreateCountryInput) (*domain.Country, error) {
	if err := validateCreate(in); err != nil {
		return nil, err
	}
	in.Name = strings.TrimSpace(in.Name)
	in.Code = strings.TrimSpace(in.Code)
	base := s.slug.Make(in.Name)

	slugs, err := s.repo.FindSimilarSlugs(ctx, base, nil)
	if err != nil {
		return nil, err
	}

	slugUnique := slug.MakeUnique(base, slugs)

	country := &domain.Country{
		ID:       uuid.NewString(), // google/uuid tạo bên server
		Name:     in.Name,
		Slug:     slugUnique,
		Code:     in.Code,
		Position: in.Position,
	}

	if err := s.repo.Create(ctx, country); err != nil {
		return nil, err
	}

	return country, nil
}

func (s *service) GetByID(ctx context.Context, id string) (*domain.Country, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *service) GetBySlug(ctx context.Context, slug string) (*domain.Country, error) {
	return s.repo.FindBySlug(ctx, slug)
}

func (s *service) List(ctx context.Context, in ListCountriesInput) (*ListCountriesOutput, error) {

	filter := domain.ListFilter{
		FilterBasic: shared2.NewFilterBasic(in.FilterBasicInput.Search, in.FilterBasicInput.Page, in.FilterBasicInput.PerPage),
	}

	items, total, err := s.repo.List(ctx, filter)
	if err != nil {
		return nil, err
	}

	return &ListCountriesOutput{
		Items:      items,
		Pagination: sharedservice.NewPagination(total, in.FilterBasicInput.Page, in.FilterBasicInput.PerPage),
	}, nil
}

func (s *service) Update(ctx context.Context, in UpdateCountryInput) (*domain.Country, error) {
	existing, err := s.repo.FindByID(ctx, in.ID)
	if err != nil {
		return nil, err
	}

	// Chỉ update field client gửi & không null.
	utils.ApplyValue(in.Name, &existing.Name)
	utils.ApplyValue(in.Slug, &existing.Slug)
	utils.ApplyValue(in.Code, &existing.Code)
	utils.ApplyValue(in.Position, &existing.Position)

	if err := s.repo.Update(ctx, existing); err != nil {
		return nil, err
	}
	return existing, nil
}

func (s *service) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
