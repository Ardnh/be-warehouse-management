package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID         uuid.UUID      `gorm:"type:uuid;primaryKey"`
	CustomerID uuid.UUID      `gorm:"type:uuid;not null;index"`
	SKU        string         `gorm:"column:sku;type:varchar(100);not null"`
	Name       string         `gorm:"type:varchar(255);not null"`
	UomID      uuid.UUID      `gorm:"type:uuid;not null;index"`
	Barcode    *string        `gorm:"type:varchar(100);index"`
	Status     string         `gorm:"type:varchar(20);not null;default:'active'"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`

	Customer *Customer `gorm:"foreignKey:CustomerID"`
	Uom      *Uom      `gorm:"foreignKey:UomID"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) error {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return nil
}
