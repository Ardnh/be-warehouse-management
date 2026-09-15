package entity

import (
	"time"

	"github.com/google/uuid"
)

type ProductWarehouse struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	ProductID   uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:uq_product_warehouse"`
	WarehouseID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:uq_product_warehouse"`
	Status      string    `gorm:"type:varchar(20);not null;default:'active'"`

	CreatedAt time.Time `gorm:"type:timestamp;default:now()"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:now()"`

	// Relations
	Product   Product   `gorm:"foreignKey:ProductID;references:ID"`
	Warehouse Warehouse `gorm:"foreignKey:WarehouseID;references:ID"`
}
