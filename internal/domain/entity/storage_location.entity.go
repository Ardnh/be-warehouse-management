package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// entity/storage_location.go
type StorageLocation struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	RackID    uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:uq_locations_rack_bay_level"`
	Code      string    `gorm:"type:varchar(50);uniqueIndex:uq_locations_code;not null"`
	Bay       int       `gorm:"not null;uniqueIndex:uq_locations_rack_bay_level"`
	Level     int       `gorm:"not null;uniqueIndex:uq_locations_rack_bay_level"`
	Capacity  int       `gorm:"not null;default:0"`
	Status    string    `gorm:"type:varchar(20);not null;default:'ACTIVE'"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Rack *Rack `gorm:"foreignKey:RackID"`
}

func (s *StorageLocation) BeforeCreate(tx *gorm.DB) error {
	if s.ID == uuid.Nil {
		s.ID = uuid.New()
	}
	return nil
}

func (StorageLocation) TableName() string {
	return "storage_locations"
}
