package dto

import (
	"time"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"github.com/google/uuid"
)

type CreateWarehouseRequest struct {
	Code    string `json:"code" validate:"required,max=50"`
	Name    string `json:"name" validate:"required,max=150"`
	Address string `json:"address" validate:"omitempty"`
	Status  string `json:"status" validate:"omitempty,oneof=ACTIVE INACTIVE"`
}

type UpdateWarehouseRequest struct {
	Name    *string `json:"name" validate:"omitempty,max=150"`
	Address *string `json:"address"`
	Status  *string `json:"status" validate:"omitempty,oneof=ACTIVE INACTIVE"`
}

type Warehouse struct {
	ID        uuid.UUID `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewWarehouseResponse(w entity.Warehouse) Warehouse {
	return Warehouse{
		ID:        w.ID,
		Code:      w.Code,
		Name:      w.Name,
		Address:   w.Address,
		Status:    w.Status,
		CreatedAt: w.CreatedAt,
		UpdatedAt: w.UpdatedAt,
	}
}
