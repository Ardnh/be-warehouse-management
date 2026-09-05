package dto

func NewPagination(filter FilterDTO, total int64) Pagination {
	pageSize := filter.Size
	if pageSize <= 0 {
		pageSize = 30
	}
	page := filter.Page
	if page <= 0 {
		page = 1
	}
	totalItems := int(total)
	return Pagination{
		CurrentPage: page,
		PageSize:    pageSize,
		TotalItems:  totalItems,
		TotalPages:  (totalItems + pageSize - 1) / pageSize,
		HasNext:     page*pageSize < totalItems,
		HasPrevious: page > 1,
	}
}

type Pagination struct {
	CurrentPage int  `json:"current_page"`
	PageSize    int  `json:"page_size"`
	TotalItems  int  `json:"total_items"`
	TotalPages  int  `json:"total_pages"`
	HasNext     bool `json:"has_next"`
	HasPrevious bool `json:"has_previous"`
}
