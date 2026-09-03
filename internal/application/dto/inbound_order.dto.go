package dto

import (
	"time"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"github.com/google/uuid"
)

type CreateInboundOrderRequest struct {
	CustomerID        uuid.UUID                       `json:"customer_id" validate:"required"`
	ExpectedArrivalAt *time.Time                      `json:"expected_arrival_at"`
	Notes             *string                         `json:"notes" validate:"omitempty,max=1000"`
	Items             []CreateInboundOrderItemRequest `json:"items" validate:"required,min=1,dive"`
}

type UpdateInboundOrderRequest struct {
	CustomerID        *uuid.UUID                      `json:"customer_id"`
	ExpectedArrivalAt *time.Time                      `json:"expected_arrival_at"`
	Notes             *string                         `json:"notes" validate:"omitempty,max=1000"`
	Items             []UpdateInboundOrderItemRequest `json:"items" validate:"omitempty,dive"`
}

type UpdateInboundOrderStatusRequest struct {
	Status string  `json:"status" validate:"required,oneof=DRAFT OPEN PARTIALLY_RECEIVED RECEIVED CANCELLED"`
	Notes  *string `json:"notes" validate:"omitempty,max=1000"`
}

type ListInboundOrderQuery struct {
	Page        int        `query:"page" validate:"omitempty,min=1"`
	Limit       int        `query:"limit" validate:"omitempty,min=1,max=100"`
	Search      string     `query:"search"` // order_number
	Status      string     `query:"status" validate:"omitempty,oneof=DRAFT OPEN PARTIALLY_RECEIVED RECEIVED CANCELLED"`
	CustomerID  *uuid.UUID `query:"customer_id"`
	ArrivalFrom *time.Time `query:"arrival_from"`
	ArrivalTo   *time.Time `query:"arrival_to"`
	SortBy      string     `query:"sort_by" validate:"omitempty,oneof=created_at expected_arrival_at order_number"`
	SortDir     string     `query:"sort_dir" validate:"omitempty,oneof=asc desc"`
}

type InboundOrderResponse struct {
	ID                uuid.UUID                  `json:"id"`
	CustomerID        uuid.UUID                  `json:"customer_id"`
	OrderNumber       string                     `json:"order_number"`
	Status            string                     `json:"status"`
	ExpectedArrivalAt *time.Time                 `json:"expected_arrival_at"`
	Notes             *string                    `json:"notes"`
	CreatedAt         time.Time                  `json:"created_at"`
	UpdatedAt         time.Time                  `json:"updated_at"`
	Items             []InboundOrderItemResponse `json:"items,omitempty"`
}

func ToInboundOrderResponse(e *entity.InboundOrder) *InboundOrderResponse {
	if e == nil {
		return nil
	}

	res := &InboundOrderResponse{
		ID:                e.ID,
		CustomerID:        e.CustomerID,
		OrderNumber:       e.OrderNumber,
		Status:            e.Status.String(),
		ExpectedArrivalAt: e.ExpectedArrivalAt,
		Notes:             e.Notes,
		CreatedAt:         e.CreatedAt,
		UpdatedAt:         e.UpdatedAt,
	}

	if len(e.Items) > 0 {
		res.Items = ToInboundOrderItemResponses(e.Items)
	}

	return res
}

func ToInboundOrderResponses(list []entity.InboundOrder) []InboundOrderResponse {
	res := make([]InboundOrderResponse, 0, len(list))
	for i := range list {
		res = append(res, *ToInboundOrderResponse(&list[i]))
	}
	return res
}
