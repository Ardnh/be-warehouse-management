package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// entity/user.go
type User struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey"`
	Username     string    `gorm:"type:varchar(50);uniqueIndex:uq_users_username;not null"`
	Email        string    `gorm:"type:varchar(150);uniqueIndex:uq_users_email;not null"`
	PasswordHash string    `gorm:"type:varchar(255);not null"`
	FullName     string    `gorm:"type:varchar(150);not null"`
	Status       string    `gorm:"type:varchar(20);not null;default:'ACTIVE'"`
	CreatedAt    time.Time
	UpdatedAt    time.Time

	Roles []Role `gorm:"many2many:user_roles;joinForeignKey:UserID;joinReferences:RoleID"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return nil
}
