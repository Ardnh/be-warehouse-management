package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LoginSession struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;index"`
	ExpiresAt time.Time `gorm:"not null;index"`
	UsedAt    *time.Time
	CreatedAt time.Time `gorm:"not null"`
}

func (LoginSession) TableName() string { return "login_sessions" }

func (c *LoginSession) BeforeCreate(tx *gorm.DB) error {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}

	c.CreatedAt = time.Now()
	return nil
}
