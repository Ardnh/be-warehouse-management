package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// entity/handling_unit.go
type HandlingUnit struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	Code        string    `gorm:"type:varchar(50);uniqueIndex:uq_handling_units_code;not null"`
	WarehouseID uuid.UUID `gorm:"type:uuid;not null;index"`
	Status      string    `gorm:"type:varchar(30);not null;default:'EMPTY'"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Warehouse *Warehouse         `gorm:"foreignKey:WarehouseID"`
	Items     []HandlingUnitItem `gorm:"foreignKey:HandlingUnitID"`
}

func (h *HandlingUnit) BeforeCreate(tx *gorm.DB) error {
	if h.ID == uuid.Nil {
		h.ID = uuid.New()
	}
	return nil
}
