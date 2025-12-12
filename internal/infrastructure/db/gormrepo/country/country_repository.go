package country

import (
	"context"
	"errors"
	"strings"
	"time"

	domain "github.com/codesayhi/golang-clean/internal/domain/country"
	"gorm.io/gorm"
)

type Model struct {
	ID        string         `gorm:"type:uuid;primaryKey;column:id"`
	Name      string         `gorm:"type:varchar(255);not null;column:name"`
	Slug      string         `gorm:"type:varchar(255);not null;uniqueIndex;column:slug"`
	Code      string         `gorm:"type:varchar(20);not null;column:code"`
	Position  int            `gorm:"not null;default:0;column:position"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime"`
}

func (Model) TableName() string {
	return "countries"
}

type Repository struct {
	db *gorm.DB
}

func NewCountryRepository(db *gorm.DB) domain.Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, c *domain.Country) error {
	model := fromDomain(c)
	if err := r.db.WithContext(ctx).Create(model).Error; err != nil {
		return err
	}
	c.CreatedAt = model.CreatedAt
	c.UpdatedAt = model.UpdatedAt
	return nil
}

func (r *Repository) Update(ctx context.Context, c *domain.Country) error {
	model := fromDomain(c)
	if err := r.db.WithContext(ctx).Save(model).Error; err != nil { // dùng Save, GORM auto update UpdatedAt
		return err
	}
	c.UpdatedAt = model.UpdatedAt
	return nil
}

func (r *Repository) Delete(ctx context.Context, id string) error {
	result := r.db.WithContext(ctx).Delete(&Model{}, "id = ?", id) // soft delete
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domain.ErrNotFound
	}
	return nil
}

func (r *Repository) FindByID(ctx context.Context, id string) (*domain.Country, error) {
	var m Model
	if err := r.db.WithContext(ctx).First(&m, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}
	return toDomain(&m), nil
}

func (r *Repository) FindBySlug(ctx context.Context, slug string) (*domain.Country, error) {
	var m Model
	if err := r.db.WithContext(ctx).First(&m, "slug = ?", slug).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}
	return toDomain(&m), nil
}

func (r *Repository) List(ctx context.Context, filter domain.ListFilter) ([]*domain.Country, int64, error) {
	var (
		models []*Model
		total  int64
	)

	db := r.db.WithContext(ctx).Model(&Model{})

	if s := strings.TrimSpace(filter.FilterBasic.Search); s != "" {
		like := "%" + s + "%"
		db = db.Where("name ILIKE ? OR code ILIKE ?", like, like)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if filter.FilterBasic.Page <= 0 {
		filter.FilterBasic.Page = 1
	}
	if filter.FilterBasic.PerPage <= 0 {
		filter.FilterBasic.PerPage = 20
	}
	offset := (filter.FilterBasic.Page - 1) * filter.FilterBasic.PerPage

	if err := db.
		Order("position ASC, name ASC").
		Limit(filter.FilterBasic.PerPage).
		Offset(offset).
		Find(&models).Error; err != nil {
		return nil, 0, err
	}

	items := make([]*domain.Country, 0, len(models))
	for _, m := range models {
		items = append(items, toDomain(m))
	}

	return items, total, nil
}
