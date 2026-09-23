package seeder

import (
	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"gorm.io/gorm"
)

func SeedSystemAdminPermissions(db *gorm.DB, role *entity.Role) error {

	var permissions []entity.Permission

	if err := db.Find(&permissions).Error; err != nil {
		return err
	}

	for _, permission := range permissions {
		rolePermission := entity.RolePermission{
			RoleID:       role.ID,
			PermissionID: permission.ID,
		}

		if err := db.
			Where(
				"role_id = ? AND permission_id = ?",
				role.ID,
				permission.ID,
			).
			FirstOrCreate(&rolePermission).Error; err != nil {
			return err
		}
	}

	return nil
}
