package dto

import (
	"time"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"github.com/google/uuid"
)

type CreateReceivingRequest struct {
	InboundOrderID uuid.UUID                    `json:"inbound_order_id" validate:"required"`
	ReceivedAt     *time.Time                   `json:"received_at"`
	ReceivedBy     *uuid.UUID                   `json:"received_by"`
	Notes          *string                      `json:"notes" validate:"omitempty,max=1000"`
	Items          []CreateReceivingItemRequest `json:"items" validate:"required,min=1,dive"`
}

type CreateReceivingItemRequest struct {
	InboundOrderItemID uuid.UUID `json:"inbound_order_item_id" validate:"required"`
	ProductID          uuid.UUID `json:"product_id" validate:"required"`
	ReceivedQty        int       `json:"received_qty" validate:"required,gt=0"`
}

type UpdateReceivingRequest struct {
	Status     *string                      `json:"status" validate:"omitempty,oneof=draft partial completed cancelled"`
	ReceivedAt *time.Time                   `json:"received_at"`
	ReceivedBy *uuid.UUID                   `json:"received_by"`
	Notes      *string                      `json:"notes" validate:"omitempty,max=1000"`
	Items      []UpdateReceivingItemRequest `json:"items" validate:"omitempty,dive"`
}

type UpdateReceivingItemRequest struct {
	ID                 *uuid.UUID `json:"id"` // kosong = item baru
	InboundOrderItemID uuid.UUID  `json:"inbound_order_item_id" validate:"required"`
	ProductID          uuid.UUID  `json:"product_id" validate:"required"`
	ReceivedQty        int        `json:"received_qty" validate:"required,gt=0"`
}

type ListReceivingQuery struct {
	Page           int        `query:"page" validate:"omitempty,min=1"`
	Limit          int        `query:"limit" validate:"omitempty,min=1,max=100"`
	Search         string     `query:"search"`
	Status         string     `query:"status" validate:"omitempty,oneof=draft partial completed cancelled"`
	InboundOrderID *uuid.UUID `query:"inbound_order_id"`
	ReceivedFrom   *time.Time `query:"received_from"`
	ReceivedTo     *time.Time `query:"received_to"`
}

func ToReceivingResponse(e *entity.Receiving) *ReceivingResponse {
	if e == nil {
		return nil
	}

	res := &ReceivingResponse{
		ID:              e.ID,
		InboundOrderID:  e.InboundOrderID,
		ReceivingNumber: e.ReceivingNumber,
		Status:          e.Status,
		ReceivedAt:      e.ReceivedAt,
		ReceivedBy:      e.ReceivedBy,
		Notes:           e.Notes,
		CreatedAt:       e.CreatedAt,
		UpdatedAt:       e.UpdatedAt,
	}

	if len(e.Items) > 0 {
		res.Items = ToReceivingItemResponses(e.Items)
	}

	return res
}

func ToReceivingResponses(list []entity.Receiving) []ReceivingResponse {
	res := make([]ReceivingResponse, 0, len(list))
	for i := range list {
		res = append(res, *ToReceivingResponse(&list[i]))
	}
	return res
}
