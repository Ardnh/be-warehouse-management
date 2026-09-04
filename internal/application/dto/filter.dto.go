package dto

type FilterDTO struct {
	Page    int    `json:"page" validate:"required,omitempty"`
	Size    int    `json:"size" validate:"required,omitempty"`
	Search  string `json:"search" validate:"required,omitempty"`
	SortBy  string `json:"sortBy" validate:"required,omitempty"`
	SortDir string `json:"sortDir" validate:"required,omitempty"`
}
