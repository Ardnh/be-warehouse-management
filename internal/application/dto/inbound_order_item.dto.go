package dto

import (
	"time"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"github.com/google/uuid"
)

type CreateInboundOrderItemRequest struct {
	ProductID   uuid.UUID `json:"product_id" validate:"required"`
	ExpectedQty int       `json:"expected_qty" validate:"required,gt=0"`
}

type UpdateInboundOrderItemRequest struct {
	ID          *uuid.UUID `json:"id"` // kosong = item baru
	ProductID   uuid.UUID  `json:"product_id" validate:"required"`
	ExpectedQty int        `json:"expected_qty" validate:"required,gt=0"`
}

type InboundOrderItemResponse struct {
	ID             uuid.UUID `json:"id"`
	InboundOrderID uuid.UUID `json:"inbound_order_id"`
	ProductID      uuid.UUID `json:"product_id"`
	ExpectedQty    int       `json:"expected_qty"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func ToInboundOrderItemResponse(e *entity.InboundOrderItem) InboundOrderItemResponse {
	return InboundOrderItemResponse{
		ID:             e.ID,
		InboundOrderID: e.InboundOrderID,
		ProductID:      e.ProductID,
		ExpectedQty:    e.ExpectedQty,
		CreatedAt:      e.CreatedAt,
		UpdatedAt:      e.UpdatedAt,
	}
}

func ToInboundOrderItemResponses(list []entity.InboundOrderItem) []InboundOrderItemResponse {
	res := make([]InboundOrderItemResponse, 0, len(list))
	for i := range list {
		res = append(res, ToInboundOrderItemResponse(&list[i]))
	}
	return res
}
