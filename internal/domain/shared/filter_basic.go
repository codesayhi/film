package shared

type FilterBasic struct {
	Search  string
	Page    int
	PerPage int
}

func NewFilterBasic(search string, page int, perPage int) FilterBasic {
	return FilterBasic{
		Search:  search,
		Page:    page,
		PerPage: perPage,
	}
}
