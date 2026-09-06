package entity

import "github.com/google/uuid"

type Permission struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Resource    string    `gorm:"type:varchar(50);not null"`
	Action      string    `gorm:"type:varchar(50);not null"`
	Description string    `gorm:"type:varchar(255);not null"`
}

func (Permission) TableName() string {
	return "permissions"
}
