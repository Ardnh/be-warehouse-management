package dto

import (
	"time"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"github.com/google/uuid"
)

// dto/handling_unit.go
type CreateHandlingUnitRequest struct {
	Code        string    `json:"code" validate:"required,max=50"`
	WarehouseID uuid.UUID `json:"warehouse_id" validate:"required"`
	Status      string    `json:"status" validate:"omitempty,oneof=EMPTY IN_USE STAGED SHIPPED DAMAGED"`
}

type UpdateHandlingUnitRequest struct {
	Status *string `json:"status" validate:"omitempty,oneof=EMPTY IN_USE STAGED SHIPPED DAMAGED"`
}

type HandlingUnitResponse struct {
	ID            uuid.UUID                  `json:"id"`
	Code          string                     `json:"code"`
	WarehouseID   uuid.UUID                  `json:"warehouse_id"`
	WarehouseCode string                     `json:"warehouse_code,omitempty"`
	Status        string                     `json:"status"`
	Items         []HandlingUnitItemResponse `json:"items,omitempty"`
	CreatedAt     time.Time                  `json:"created_at"`
	UpdatedAt     time.Time                  `json:"updated_at"`
}

func NewHandlingUnitResponse(h entity.HandlingUnit) HandlingUnitResponse {
	res := HandlingUnitResponse{
		ID:          h.ID,
		Code:        h.Code,
		WarehouseID: h.WarehouseID,
		Status:      h.Status,
		CreatedAt:   h.CreatedAt,
		UpdatedAt:   h.UpdatedAt,
	}
	if h.Warehouse != nil {
		res.WarehouseCode = h.Warehouse.Code
	}
	for _, item := range h.Items {
		res.Items = append(res.Items, NewHandlingUnitItemResponse(item))
	}
	return res
}

func NewHandlingUnitResponses(hus []entity.HandlingUnit) []HandlingUnitResponse {
	res := make([]HandlingUnitResponse, 0, len(hus))
	for _, h := range hus {
		res = append(res, NewHandlingUnitResponse(h))
	}
	return res
}
