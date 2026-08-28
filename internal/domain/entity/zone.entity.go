package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// entity/zone.go
type Zone struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	WarehouseID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:uq_zones_warehouse_code" json:"warehouse_id"`
	Code        string    `gorm:"type:varchar(50);not null;uniqueIndex:uq_zones_warehouse_code" json:"code"`
	Name        string    `gorm:"type:varchar(100);not null" json:"name"`
	Type        string    `gorm:"type:varchar(30);not null" json:"type"`
	Status      string    `gorm:"type:varchar(20);not null;default:'ACTIVE'" json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	Warehouse *Warehouse `gorm:"foreignKey:WarehouseID" json:"warehouse,omitempty"`
	Racks     []Rack     `gorm:"foreignKey:ZoneID" json:"racks,omitempty"`
}

func (z *Zone) BeforeCreate(tx *gorm.DB) error {
	if z.ID == uuid.Nil {
		z.ID = uuid.New()
	}
	return nil
}
