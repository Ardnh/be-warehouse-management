package repositories

type Filter struct {
	Search   string
	Page     int
	PageSize int
	SortBy   string
	SortDir  string
}
