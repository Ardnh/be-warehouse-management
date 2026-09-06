package entity

import "github.com/google/uuid"

type UserPermission struct {
	UserID       uuid.UUID `gorm:"type:uuid;not null"`
	PermissionID uuid.UUID `gorm:"type:uuid;not null"`
}

func (UserPermission) TableName() string {
	return "user_permissions"
}
