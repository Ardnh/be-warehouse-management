package seeder

import (
	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"gorm.io/gorm"
)

func SeedSystemAdminRole(db *gorm.DB) (*entity.Role, error) {
	role := entity.Role{
		Name:        "system-admin",
		Code:        "SYSTEM_ADMIN",
		Description: "System admin role",
	}
	if err := db.Where("code = ?", "SYSTEM_ADMIN").
		FirstOrCreate(&role, entity.Role{Code: "SYSTEM_ADMIN"}).Error; err != nil {
		return nil, err
	}
	return &role, nil
}
