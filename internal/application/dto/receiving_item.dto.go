package dto

import "github.com/google/uuid"

type ReceivingItemResponse struct {
	ID                 uuid.UUID `json:"id"`
	InboundOrderItemID uuid.UUID `json:"inbound_order_item_id"`
	ProductID          uuid.UUID `json:"product_id"`
	ReceivedQty        int       `json:"received_qty"`
}

type CreateReceivingItem struct {
	InboundOrderItemID uuid.UUID `json:"inbound_order_item_id" validate:"required"`
	ProductID          uuid.UUID `json:"product_id" validate:"required"`
	ReceivedQty        int       `json:"received_qty" validate:"required,gt=0"`
}
