package dto

import (
	"time"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"github.com/google/uuid"
)

type ReceivingItemResponse struct {
	ID                 uuid.UUID `json:"id"`
	ReceivingID        uuid.UUID `json:"receiving_id"`
	InboundOrderItemID uuid.UUID `json:"inbound_order_item_id"`
	ProductID          uuid.UUID `json:"product_id"`
	ReceivedQty        int       `json:"received_qty"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type CreateReceivingItemRequest struct {
	InboundOrderItemID uuid.UUID `json:"inbound_order_item_id" validate:"required"`
	ProductID          uuid.UUID `json:"product_id" validate:"required"`
	ReceivedQty        int       `json:"received_qty" validate:"required,gt=0"`
}

// CreateReceivingItem is kept as an alias for callers using the original name.
type CreateReceivingItem = CreateReceivingItemRequest

type UpdateReceivingItemRequest struct {
	ID                 *uuid.UUID `json:"id"`
	InboundOrderItemID uuid.UUID  `json:"inbound_order_item_id" validate:"required"`
	ProductID          uuid.UUID  `json:"product_id" validate:"required"`
	ReceivedQty        int        `json:"received_qty" validate:"required,gt=0"`
}

func ToReceivingItemResponse(e *entity.ReceivingItem) ReceivingItemResponse {
	return ReceivingItemResponse{
		ID:                 e.ID,
		ReceivingID:        e.ReceivingID,
		InboundOrderItemID: e.InboundOrderItemID,
		ProductID:          e.ProductID,
		ReceivedQty:        e.ReceivedQty,
		CreatedAt:          e.CreatedAt,
		UpdatedAt:          e.UpdatedAt,
	}
}

func ToReceivingItemResponses(list []entity.ReceivingItem) []ReceivingItemResponse {
	res := make([]ReceivingItemResponse, 0, len(list))
	for i := range list {
		res = append(res, ToReceivingItemResponse(&list[i]))
	}
	return res
}
