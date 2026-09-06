package entity

import (
	"time"

	"github.com/google/uuid"
)

// entity/user_role.go
type UserRole struct {
	UserID      uuid.UUID `gorm:"type:uuid;primaryKey"`
	RoleID      uuid.UUID `gorm:"type:uuid;primaryKey"`
	WarehouseID uuid.UUID `gorm:"type:uuid"`
	CreatedAt   time.Time

	User      *User      `gorm:"foreignKey:UserID"`
	Role      *Role      `gorm:"foreignKey:RoleID"`
	Warehouse *Warehouse `gorm:"foreignKey:WarehouseID"`
}

func (UserRole) TableName() string {
	return "user_roles"
}
