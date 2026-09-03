package entity

import (
	"time"

	"github.com/google/uuid"
)

type InboundOrderItem struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	InboundOrderID uuid.UUID `gorm:"type:uuid;not null;index"`
	ProductID      uuid.UUID `gorm:"type:uuid;not null;index"`
	ExpectedQty    int       `gorm:"type:integer;not null;default:0"`
	CreatedAt      time.Time `gorm:"type:timestamptz;not null;default:now()"`
	UpdatedAt      time.Time `gorm:"type:timestamptz;not null;default:now()"`

	InboundOrder   *InboundOrder   `gorm:"foreignKey:InboundOrderID"`
	Product        *Product        `gorm:"foreignKey:ProductID"`
	ReceivingItems []ReceivingItem `gorm:"foreignKey:InboundOrderItemID"`
}

func (InboundOrderItem) TableName() string { return "inbound_order_items" }
