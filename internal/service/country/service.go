package country

import (
	"context"
	"strings"

	domain "github.com/codesayhi/golang-clean/internal/domain/country"
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
}

func NewService(repo domain.Repository) Service {
	return &service{repo: repo}
}

func (s *service) Create(ctx context.Context, in CreateCountryInput) (*domain.Country, error) {
	in.Name = strings.TrimSpace(in.Name)
	in.Slug = strings.TrimSpace(in.Slug)
	in.Code = strings.TrimSpace(in.Code)

	country := &domain.Country{
		ID:       uuid.NewString(), // google/uuid tạo bên server
		Name:     in.Name,
		Slug:     in.Slug,
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
	if in.Page <= 0 {
		in.Page = 1
	}
	if in.PerPage <= 0 {
		in.PerPage = 20
	}

	filter := domain.ListFilter{
		Search:  strings.TrimSpace(in.Search),
		Page:    in.Page,
		PerPage: in.PerPage,
	}

	items, total, err := s.repo.List(ctx, filter)
	if err != nil {
		return nil, err
	}

	return &ListCountriesOutput{
		Items: items,
		Total: total,
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
