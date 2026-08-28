package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// entity/rack.go
type Rack struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	ZoneID         uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:uq_racks_zone_code" json:"zone_id"`
	Code           string    `gorm:"type:varchar(50);not null;uniqueIndex:uq_racks_zone_code" json:"code"`
	BayCount       int       `gorm:"not null;default:0" json:"bay_count"`
	LevelCount     int       `gorm:"not null;default:0" json:"level_count"`
	PalletCapacity int       `gorm:"not null;default:0" json:"pallet_capacity"`
	Status         string    `gorm:"type:varchar(20);not null;default:'ACTIVE'" json:"status"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`

	Zone      *Zone             `gorm:"foreignKey:ZoneID" json:"zone,omitempty"`
	Locations []StorageLocation `gorm:"foreignKey:RackID" json:"locations,omitempty"`
}

func (r *Rack) BeforeCreate(tx *gorm.DB) error {
	if r.ID == uuid.Nil {
		r.ID = uuid.New()
	}
	return nil
}
