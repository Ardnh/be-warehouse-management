package entity

import (
	"time"

	"github.com/google/uuid"
)

type Receiving struct {
	ID              uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	InboundOrderID  uuid.UUID  `gorm:"type:uuid;not null;index"`
	ReceivingNumber string     `gorm:"type:varchar(50);not null;uniqueIndex"`
	Status          string     `gorm:"type:varchar(30);not null"`
	ReceivedAt      *time.Time `gorm:"type:timestamptz"`
	ReceivedBy      *uuid.UUID `gorm:"type:uuid;index"`
	Notes           *string    `gorm:"type:text"`
	CreatedAt       time.Time  `gorm:"type:timestamptz;not null;default:now()"`
	UpdatedAt       time.Time  `gorm:"type:timestamptz;not null;default:now()"`

	InboundOrder *InboundOrder   `gorm:"foreignKey:InboundOrderID"`
	Receiver     *User           `gorm:"foreignKey:ReceivedBy"`
	Items        []ReceivingItem `gorm:"foreignKey:ReceivingID"`
}

func (Receiving) TableName() string { return "receivings" }
