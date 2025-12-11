package country

import (
	"errors"
	"time"
)

type Country struct {
	ID        string
	Name      string
	Slug      string
	Code      string
	Position  int
	DeletedAt *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

// ErrNotFound dùng chung giữa repo / usecase / handler.
var ErrNotFound = errors.New("country not found")
