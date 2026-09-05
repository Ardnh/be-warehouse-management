package dto

import (
	"time"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"github.com/google/uuid"
)

type Customer struct {
	ID        uuid.UUID `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type CreateCustomerRequest struct {
	Code    string `json:"code" validate:"required"`
	Name    string `json:"name" validate:"required"`
	Address string `json:"address" validate:"required"`
	Status  string `json:"status" validate:"required"`
	Email   string `json:"email" validate:"required"`
	Phone   string `json:"phone" validate:"required"`
}

type UpdateCustomerRequest struct {
	Name    string `json:"name" validate:"required"`
	Address string `json:"address" validate:"required"`
	Status  string `json:"status" validate:"required"`
	Email   string `json:"email" validate:"required"`
	Phone   string `json:"phone" validate:"required"`
}

func ToCustomerDTO(c *entity.Customer) Customer {
	if c == nil {
		return Customer{}
	}
	return Customer{
		ID:        c.ID,
		Code:      c.Code,
		Name:      c.Name,
		Address:   c.Address,
		Status:    c.Status,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
		DeletedAt: c.DeletedAt.Time,
	}
}

func ToCustomerDTOs(customers []entity.Customer) []Customer {
	customerDTOs := make([]Customer, 0, len(customers))
	for _, c := range customers {
		customerDTOs = append(customerDTOs, ToCustomerDTO(&c))
	}
	return customerDTOs
}
