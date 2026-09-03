package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Receiving struct {
	ID              uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	InboundOrderID  uuid.UUID `gorm:"type:uuid;not null"`
	ReceivingNumber string    `gorm:"type:varchar(50);not null"`
	Status          string    `gorm:"type:varchar(30);not null"`
	ReceivedAt      time.Time `gorm:"not null"`
	ReceivedBy      uuid.UUID `gorm:"type:uuid;not null"`
	Notes           *string   `gorm:"type:text"`
	CreatedAt       time.Time `gorm:"not null;autoCreateTime"`
	UpdatedAt       time.Time `gorm:"not null;autoUpdateTime"`
}

func (Receiving) TableName() string {
	return "receivings"
}

func (r *Receiving) BeforeCreate(tx *gorm.DB) error {
	if r.ID == uuid.Nil {
		r.ID = uuid.New()
	}
	return nil
}
