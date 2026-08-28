package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Uom struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Code      string    `gorm:"type:varchar(20);uniqueIndex:uq_uoms_code;not null"`
	Name      string    `gorm:"type:varchar(100);not null"`
	Type      string    `gorm:"type:varchar(20);not null"`
	CreatedAt time.Time `gorm:"index"`
	UpdatedAt time.Time `gorm:"index"`
}

func (u *Uom) BeforeCreate(tx *gorm.DB) error {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return nil
}
