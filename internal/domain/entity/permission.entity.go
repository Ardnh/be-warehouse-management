package entity

import "github.com/google/uuid"

type Permission struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Resource    string    `gorm:"size:64;uniqueIndex:idx_resource_action"`
	Action      string    `gorm:"size:32;uniqueIndex:idx_resource_action"`
	Description string    `gorm:"type:varchar(255);not null"`
}

func (Permission) TableName() string {
	return "permissions"
}
