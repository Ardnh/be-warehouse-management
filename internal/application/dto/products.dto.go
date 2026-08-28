package dto

import (
	"time"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

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
