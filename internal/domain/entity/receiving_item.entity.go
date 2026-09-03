package entity

import (
	"time"

	"github.com/google/uuid"
)

type ReceivingItem struct {
	ID                 uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	ReceivingID        uuid.UUID `gorm:"type:uuid;not null;index"`
	InboundOrderItemID uuid.UUID `gorm:"type:uuid;not null;index"`
	ProductID          uuid.UUID `gorm:"type:uuid;not null;index"`
	ReceivedQty        int       `gorm:"type:integer;not null;default:0"`
	CreatedAt          time.Time `gorm:"type:timestamptz;not null;default:now()"`
	UpdatedAt          time.Time `gorm:"type:timestamptz;not null;default:now()"`

	Receiving        *Receiving        `gorm:"foreignKey:ReceivingID"`
	InboundOrderItem *InboundOrderItem `gorm:"foreignKey:InboundOrderItemID"`
	Product          *Product          `gorm:"foreignKey:ProductID"`
}

func (ReceivingItem) TableName() string { return "receiving_items" }
