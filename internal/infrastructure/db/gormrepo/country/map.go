package country

import (
	"time"

	domain "github.com/codesayhi/golang-clean/internal/domain/country"
)

// mapping helpers

func toDomain(m *Model) *domain.Country {
	if m == nil {
		return nil
	}

	var deletedAt *time.Time
	if m.DeletedAt.Valid {
		deletedAt = &m.DeletedAt.Time
	}

	return &domain.Country{
		ID:        m.ID,
		Name:      m.Name,
		Slug:      m.Slug,
		Code:      m.Code,
		Position:  m.Position,
		DeletedAt: deletedAt,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func fromDomain(c *domain.Country) *Model {
	m := &Model{
		ID:       c.ID,
		Name:     c.Name,
		Slug:     c.Slug,
		Code:     c.Code,
		Position: c.Position,
	}

	if !c.CreatedAt.IsZero() {
		m.CreatedAt = c.CreatedAt
	}
	if !c.UpdatedAt.IsZero() {
		m.UpdatedAt = c.UpdatedAt
	}
	if c.DeletedAt != nil {
		m.DeletedAt.Time = *c.DeletedAt
		m.DeletedAt.Valid = true
	}

	return m
}
