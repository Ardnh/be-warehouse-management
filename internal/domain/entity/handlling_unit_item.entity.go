package entity // entity/handling_unit_item.go
import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type HandlingUnitItem struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
	HandlingUnitID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:uq_hu_items_hu_product"`
	ProductID      uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:uq_hu_items_hu_product"`
	Quantity       int       `gorm:"not null;default:0"`
	CreatedAt      time.Time
	UpdatedAt      time.Time

	HandlingUnit *HandlingUnit `gorm:"foreignKey:HandlingUnitID"`
	Product      *Product      `gorm:"foreignKey:ProductID"`
}

func (i *HandlingUnitItem) BeforeCreate(tx *gorm.DB) error {
	if i.ID == uuid.Nil {
		i.ID = uuid.New()
	}
	return nil
}
