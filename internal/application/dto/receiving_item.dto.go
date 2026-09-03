package dto

import (
	"time"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"github.com/google/uuid"
)

type ReceivingResponse struct {
	ID              uuid.UUID               `json:"id"`
	InboundOrderID  uuid.UUID               `json:"inbound_order_id"`
	ReceivingNumber string                  `json:"receiving_number"`
	Status          string                  `json:"status"`
	ReceivedAt      *time.Time              `json:"received_at"`
	ReceivedBy      *uuid.UUID              `json:"received_by"`
	Notes           *string                 `json:"notes"`
	CreatedAt       time.Time               `json:"created_at"`
	UpdatedAt       time.Time               `json:"updated_at"`
	Items           []ReceivingItemResponse `json:"items,omitempty"`
}

type ReceivingItemResponse struct {
	ID                 uuid.UUID `json:"id"`
	ReceivingID        uuid.UUID `json:"receiving_id"`
	InboundOrderItemID uuid.UUID `json:"inbound_order_item_id"`
	ProductID          uuid.UUID `json:"product_id"`
	ReceivedQty        int       `json:"received_qty"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
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
