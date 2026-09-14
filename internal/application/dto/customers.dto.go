package dto

import (
	"time"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"github.com/google/uuid"
)

type Customer struct {
	ID            uuid.UUID  `json:"id"`
	WarehouseID   uuid.UUID  `json:"warehouse_id"`
	WarehouseCode string     `json:"warehouse_code,omitempty"`
	Code          string     `json:"code"`
	Name          string     `json:"name"`
	Email         string     `json:"email"`
	Address       string     `json:"address"`
	Status        string     `json:"status"`
	Phone         string     `json:"phone"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at,omitempty"`
}

type CreateCustomerRequest struct {
	WarehouseID uuid.UUID `json:"warehouse_id" validate:"required"`
	Code        string    `json:"code" validate:"required,max=50"`
	Name        string    `json:"name" validate:"required,max=255"`
	Email       string    `json:"email" validate:"required,email,max=50"`
	Address     string    `json:"address" validate:"required"`
	Status      string    `json:"status" validate:"required"`
	Phone       string    `json:"phone" validate:"required,max=20"`
}

type UpdateCustomerRequest struct {
	Name    string `json:"name" validate:"required,max=255"`
	Address string `json:"address" validate:"required"`
	Status  string `json:"status" validate:"required"`
	Email   string `json:"email" validate:"required,email,max=50"`
	Phone   string `json:"phone" validate:"required,max=20"`
}

func ToCustomerDTO(c *entity.Customer) Customer {
	if c == nil {
		return Customer{}
	}
	res := Customer{
		ID:          c.ID,
		WarehouseID: c.WarehouseID,
		Code:        c.Code,
		Name:        c.Name,
		Email:       c.Email,
		Address:     c.Address,
		Status:      c.Status,
		Phone:       c.Phone,
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
	}
	if c.Warehouse != nil {
		res.WarehouseCode = c.Warehouse.Code
	}
	if c.DeletedAt.Valid {
		deletedAt := c.DeletedAt.Time
		res.DeletedAt = &deletedAt
	}
	return res
}

func ToCustomerDTOs(customers []entity.Customer) []Customer {
	customerDTOs := make([]Customer, 0, len(customers))
	for _, c := range customers {
		customerDTOs = append(customerDTOs, ToCustomerDTO(&c))
	}
	return customerDTOs
}
