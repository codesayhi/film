package genre

import "time"

type Genre struct {
	ID        string
	Name      string
	Slug      string
	Position  int
	DeletedAt *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
