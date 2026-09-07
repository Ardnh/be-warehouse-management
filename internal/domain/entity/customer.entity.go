package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Customer struct {
	ID          uuid.UUID      `gorm:"type:uuid;primaryKey"`
	WarehouseID uuid.UUID      `gorm:"type:uuid;not null"`
	Code        string         `gorm:"type:varchar(50);uniqueIndex:uq_customer_code;not null"`
	Name        string         `gorm:"type:varchar(255);not null"`
	Email       string         `gorm:"type:varchar(50);uniqueIndex:uq_customer_email;not null"`
	Address     string         `gorm:"type:text"`
	Status      string         `gorm:"type:varchar(20);not null;default:'active'"`
	Phone       string         `gorm:"type:varchar(20)"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`

	Warehouse *Warehouse `gorm:"foreignKey:WarehouseID"`
}

func (Customer) TableName() string { return "customers" }

func (c *Customer) BeforeCreate(tx *gorm.DB) error {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}
	return nil
}
