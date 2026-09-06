package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// entity/role.go
type Role struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	Code        string    `gorm:"type:varchar(50);uniqueIndex:uq_roles_code;not null"`
	Name        string    `gorm:"type:varchar(100);not null"`
	Description string    `gorm:"type:text"`
	Status      string    `gorm:"type:varchar(20);not null;default:'ACTIVE'"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Users []User `gorm:"many2many:user_roles;joinForeignKey:RoleID;joinReferences:UserID"`
}

func (r *Role) BeforeCreate(tx *gorm.DB) error {
	if r.ID == uuid.Nil {
		r.ID = uuid.New()
	}
	return nil
}

func (Role) TableName() string {
	return "roles"
}
