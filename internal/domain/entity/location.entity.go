package entity

import (
	"time"

	"github.com/google/uuid"
)

type LocationType string

const (
	LocationTypeHO        LocationType = "HO"
	LocationTypeWarehouse LocationType = "WAREHOUSE"
)

type Location struct {
	ID         uuid.UUID    `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Code       string       `gorm:"type:varchar(30);not null;uniqueIndex"`
	Name       string       `gorm:"type:varchar(100);not null"`
	Type       LocationType `gorm:"type:varchar(20);not null;index"`
	Address    *string      `gorm:"type:text"`
	City       *string      `gorm:"type:varchar(100)"`
	Province   *string      `gorm:"type:varchar(100)"`
	PostalCode *string      `gorm:"type:varchar(10)"`

	IsActive bool `gorm:"not null;default:true"`

	CreatedAt time.Time `gorm:"type:timestamptz;not null;default:now()"`
	UpdatedAt time.Time `gorm:"type:timestamptz;not null;default:now()"`
}
