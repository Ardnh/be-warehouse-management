package dto

import (
	"time"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"github.com/google/uuid"
)

// dto/handling_unit_item.go
type AddHandlingUnitItemRequest struct {
	ProductID uuid.UUID `json:"product_id" validate:"required"`
	Quantity  int       `json:"quantity" validate:"required,gt=0"`
}

type UpdateHandlingUnitItemRequest struct {
	Quantity *int `json:"quantity" validate:"omitempty,gte=0"`
}

type HandlingUnitItemResponse struct {
	ID             uuid.UUID `json:"id"`
	HandlingUnitID uuid.UUID `json:"handling_unit_id"`
	ProductID      uuid.UUID `json:"product_id"`
	ProductSKU     string    `json:"product_sku,omitempty"`
	ProductName    string    `json:"product_name,omitempty"`
	Quantity       int       `json:"quantity"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func NewHandlingUnitItemResponse(i entity.HandlingUnitItem) HandlingUnitItemResponse {
	res := HandlingUnitItemResponse{
		ID:             i.ID,
		HandlingUnitID: i.HandlingUnitID,
		ProductID:      i.ProductID,
		Quantity:       i.Quantity,
		CreatedAt:      i.CreatedAt,
		UpdatedAt:      i.UpdatedAt,
	}
	if i.Product != nil {
		res.ProductSKU = i.Product.SKU
		res.ProductName = i.Product.Name
	}
	return res
}
