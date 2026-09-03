package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReceivingItem struct {
	ID                 uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	ReceivingID        uuid.UUID `gorm:"type:uuid;not null"`
	InboundOrderItemID uuid.UUID `gorm:"type:uuid;not null"`
	ProductID          uuid.UUID `gorm:"type:uuid;not null"`
	ReceivedQty        int       `gorm:"not null"`
	CreatedAt          time.Time `gorm:"not null;autoCreateTime"`
	UpdatedAt          time.Time `gorm:"not null;autoUpdateTime"`
}

func (ReceivingItem) TableName() string {
	return "receiving_items"
}

func (r *ReceivingItem) BeforeCreate(tx *gorm.DB) error {
	if r.ID == uuid.Nil {
		r.ID = uuid.New()
	}
	return nil
}
