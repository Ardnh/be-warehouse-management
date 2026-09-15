package entity

import (
	"time"

	"github.com/google/uuid"
)

type CustomerWarehouse struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	CustomerID  uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:uq_customer_warehouse" json:"customer_id"`
	WarehouseID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:uq_customer_warehouse" json:"warehouse_id"`
	Status      string    `gorm:"type:varchar(20);not null;default:'active'" json:"status"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relations
	Customer  Customer  `gorm:"foreignKey:CustomerID;references:ID" json:"customer,omitempty"`
	Warehouse Warehouse `gorm:"foreignKey:WarehouseID;references:ID" json:"warehouse,omitempty"`
}
