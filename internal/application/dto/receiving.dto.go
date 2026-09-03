package dto

import (
	"time"

	"github.com/google/uuid"
)

type ReceivingResponse struct {
	ID              uuid.UUID               `json:"id"`
	InboundOrderID  uuid.UUID               `json:"inbound_order_id"`
	ReceivingNumber string                  `json:"receiving_number"`
	Status          string                  `json:"status"`
	ReceivedAt      time.Time               `json:"received_at"`
	ReceivedBy      uuid.UUID               `json:"received_by"`
	Notes           *string                 `json:"notes,omitempty"`
	Items           []ReceivingItemResponse `json:"items"`
	CreatedAt       time.Time               `json:"created_at"`
	UpdatedAt       time.Time               `json:"updated_at"`
}

type CreateReceivingRequest struct {
	InboundOrderID uuid.UUID             `json:"inbound_order_id" validate:"required"`
	ReceivedBy     uuid.UUID             `json:"received_by" validate:"required"`
	Notes          *string               `json:"notes,omitempty"`
	Items          []CreateReceivingItem `json:"items" validate:"required,min=1"`
}
