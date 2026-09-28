package dto

import (
	"time"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"github.com/google/uuid"
)

type CreateWarehouseRequest struct {
	LocationID uuid.UUID `json:"location_id" validate:"required"`
	Status     string    `json:"status" validate:"omitempty,oneof=ACTIVE INACTIVE"`
}

type UpdateWarehouseRequest struct {
	LocationID *uuid.UUID `json:"location_id" validate:"omitempty"`
	Status     *string    `json:"status" validate:"omitempty,oneof=ACTIVE INACTIVE"`
}

type Warehouse struct {
	ID         uuid.UUID `json:"id"`
	LocationID uuid.UUID `json:"location_id"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`

	Location *LocationResponse `json:"location,omitempty"`
	Zones    []ZoneResponse    `json:"zones,omitempty"`
}

func ToWarehouseResponse(w *entity.Warehouse) Warehouse {
	warehouse := Warehouse{
		ID:         w.ID,
		LocationID: w.LocationID,
		Status:     w.Status,
		CreatedAt:  w.CreatedAt,
		UpdatedAt:  w.UpdatedAt,
	}

	if w.Location != nil {
		warehouse.Location = ToLocationResponse(w.Location)
	}

	if w.Zones != nil {
		warehouse.Zones = ToZoneResponses(w.Zones)
	}

	return warehouse
}
