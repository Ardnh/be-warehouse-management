package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserAssignment struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID     uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:uq_user_assignment"`
	LocationID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:uq_user_assignment"`
	RoleID     uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:uq_user_assignment"`

	CreatedAt time.Time
	UpdatedAt time.Time

	User     *User     `gorm:"foreignKey:UserID;references:ID"`
	Location *Location `gorm:"foreignKey:LocationID;references:ID"`
	Role     *Role     `gorm:"foreignKey:RoleID;references:ID"`
}

func (UserAssignment) TableName() string {
	return "user_assignments"
}

func (u *UserAssignment) BeforeCreate(tx *gorm.DB) error {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}

	u.CreatedAt = time.Now()
	return nil
}
