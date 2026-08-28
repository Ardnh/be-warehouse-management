package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Warehouse struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Code      string    `gorm:"type:varchar(50);uniqueIndex:uq_warehouses_code;not null" json:"code"`
	Name      string    `gorm:"type:varchar(150);not null" json:"name"`
	Address   string    `gorm:"type:text" json:"address"`
	Status    string    `gorm:"type:varchar(20);not null;default:'ACTIVE'" json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Zones []Zone `gorm:"foreignKey:WarehouseID" json:"zones,omitempty"`
}

func (w *Warehouse) BeforeCreate(tx *gorm.DB) error {
	if w.ID == uuid.Nil {
		w.ID = uuid.New()
	}
	return nil
}
