package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Warehouse struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	LocationID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex"`
	Status     string    `gorm:"type:varchar(20);not null;default:'ACTIVE'"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`

	Zones    []*Zone   `gorm:"foreignKey:WarehouseID"`
	Location *Location `gorm:"foreignKey:LocationID;references:ID"`
}

func (w *Warehouse) BeforeCreate(tx *gorm.DB) error {
	if w.ID == uuid.Nil {
		w.ID = uuid.New()
	}
	return nil
}
