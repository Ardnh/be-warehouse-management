package entity

import "github.com/google/uuid"

type UserPermission struct {
	UserID       uuid.UUID   `gorm:"type:uuid;not null"`
	PermissionID uuid.UUID   `gorm:"type:uuid;not null"`
	Permission   *Permission `gorm:"foreignKey:PermissionID;references:ID"`
}

func (UserPermission) TableName() string {
	return "user_permissions"
}
