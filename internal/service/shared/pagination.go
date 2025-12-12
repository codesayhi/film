package sharedservice

type Pagination struct {
	Total   int64
	Page    int
	PerPage int
}

func NewPagination(total int64, page, perPage int) Pagination {
	return Pagination{
		Total:   total,
		Page:    page,
		PerPage: perPage,
	}
}
