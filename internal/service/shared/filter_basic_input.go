package sharedservice

type FilterBasicInput struct {
	Search  string
	Page    int
	PerPage int
}

func NewFilterBasicInput(search string, page int, perPage int) FilterBasicInput {
	if page <= 0 {
		page = 1
	}
	if perPage <= 0 {
		perPage = 20
	}
	return FilterBasicInput{
		Search:  search,
		Page:    page,
		PerPage: perPage,
	}
}
