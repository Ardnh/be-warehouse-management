package dto

import (
	"time"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CreateProductRequest struct {
	CustomerID uuid.UUID `json:"customer_id" validate:"required"`
	SKU        string    `json:"sku" validate:"required,max=100"`
	Name       string    `json:"name" validate:"required,max=255"`
	UomID      uuid.UUID `json:"uom_id" validate:"required"`
	Barcode    *string   `json:"barcode" validate:"omitempty,max=100"`
	Status     string    `json:"status" validate:"omitempty,oneof=active inactive"`
}

type UpdateProductRequest struct {
	Name    *string `json:"name" validate:"omitempty,max=255"`
	Barcode *string `json:"barcode" validate:"omitempty,max=100"`
	Status  *string `json:"status" validate:"omitempty,oneof=active inactive"`
}

type Product struct {
	ID         uuid.UUID      `json:"id"`
	CustomerID uuid.UUID      `json:"customer_id"`
	SKU        string         `json:"sku"`
	Name       string         `json:"name"`
	UomID      uuid.UUID      `json:"uom_id"`
	Barcode    *string        `json:"barcode"`
	Status     string         `json:"status"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-"`

	Customer *Customer `json:"customer,omitempty"`
	Uom      *Uom      `json:"uom,omitempty"`
}

func NewProductResponse(p entity.Product) Product {
	return Product{
		ID:         p.ID,
		CustomerID: p.CustomerID,
		SKU:        p.SKU,
		Name:       p.Name,
		UomID:      p.UomID,
		Barcode:    p.Barcode,
		Status:     p.Status,
		CreatedAt:  p.CreatedAt,
		UpdatedAt:  p.UpdatedAt,
		DeletedAt:  p.DeletedAt,
	}
}
