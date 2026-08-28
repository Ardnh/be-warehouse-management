package dto

import (
	"time"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"github.com/google/uuid"
)

// dto/zone.go
type CreateZoneRequest struct {
	WarehouseID uuid.UUID `json:"warehouse_id" validate:"required"`
	Code        string    `json:"code" validate:"required,max=50"`
	Name        string    `json:"name" validate:"required,max=100"`
	Type        string    `json:"type" validate:"required,oneof=RECEIVING STORAGE PACKING OUTBOUND DAMAGED STAGING"`
	Status      string    `json:"status" validate:"omitempty,oneof=ACTIVE INACTIVE"`
}

type UpdateZoneRequest struct {
	Name   *string `json:"name" validate:"omitempty,max=100"`
	Type   *string `json:"type" validate:"omitempty,oneof=RECEIVING STORAGE PACKING OUTBOUND DAMAGED STAGING"`
	Status *string `json:"status" validate:"omitempty,oneof=ACTIVE INACTIVE"`
}

type ZoneResponse struct {
	ID            uuid.UUID `json:"id"`
	WarehouseID   uuid.UUID `json:"warehouse_id"`
	WarehouseCode string    `json:"warehouse_code,omitempty"`
	Code          string    `json:"code"`
	Name          string    `json:"name"`
	Type          string    `json:"type"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func NewZoneResponse(z entity.Zone) ZoneResponse {
	res := ZoneResponse{
		ID:          z.ID,
		WarehouseID: z.WarehouseID,
		Code:        z.Code,
		Name:        z.Name,
		Type:        z.Type,
		Status:      z.Status,
		CreatedAt:   z.CreatedAt,
		UpdatedAt:   z.UpdatedAt,
	}
	if z.Warehouse != nil {
		res.WarehouseCode = z.Warehouse.Code
	}
	return res
}
