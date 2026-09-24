package seeder

import (
	"errors"
	"fmt"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"gorm.io/gorm"
)

func SeedUserAssignments(db *gorm.DB, user *entity.User, role *entity.Role, location *entity.Location) error {
	if user == nil || role == nil || location == nil {
		return errors.New("seed user assignment: user, role, and location must not be nil")
	}

	var assignment entity.UserAssignment
	result := db.
		Where("user_id = ?", user.ID).
		Attrs(entity.UserAssignment{
			RoleID:     role.ID,
			LocationID: location.ID,
		}).
		FirstOrCreate(&assignment, entity.UserAssignment{UserID: user.ID})

	if result.Error != nil {
		return fmt.Errorf("seed user assignment for user %s: %w", user.ID, result.Error)
	}

	return nil
}
