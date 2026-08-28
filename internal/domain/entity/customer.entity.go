package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Customer struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey"`
	Code      string         `gorm:"type:varchar(50);uniqueIndex;not null"`
	Name      string         `gorm:"type:varchar(255);not null"`
	Address   string         `gorm:"type:text"`
	Status    string         `gorm:"type:varchar(20);not null;default:'active'"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Customer) TableName() string { return "customers" }

func (c *Customer) BeforeCreate(tx *gorm.DB) error {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}
	return nil
}
