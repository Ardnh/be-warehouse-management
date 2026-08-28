package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// entity/storage_location.go
type StorageLocation struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	RackID    uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:uq_locations_rack_bay_level" json:"rack_id"`
	Code      string    `gorm:"type:varchar(50);uniqueIndex:uq_locations_code;not null" json:"code"`
	Bay       int       `gorm:"not null;uniqueIndex:uq_locations_rack_bay_level" json:"bay"`
	Level     int       `gorm:"not null;uniqueIndex:uq_locations_rack_bay_level" json:"level"`
	Capacity  int       `gorm:"not null;default:0" json:"capacity"`
	Status    string    `gorm:"type:varchar(20);not null;default:'ACTIVE'" json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Rack *Rack `gorm:"foreignKey:RackID" json:"rack,omitempty"`
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
