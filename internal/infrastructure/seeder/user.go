package seeder

import (
	"time"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedUser(db *gorm.DB) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("123123"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Create a user
	user := entity.User{
		Username:     "adam",
		Email:        "adam@example.com",
		PasswordHash: string(hashedPassword),
		FullName:     "adam 1",
		Status:       "ACTIVE",
		CreatedAt:    time.Now(),
	}

	result := db.Create(&user)

	// Create a role
	role := entity.Role{
		Name:        "admin",
		Code:        "ADMIN",
		Description: "Admin role",
	}
	db.Create(&role)

	// Create a user-role association
	userRole := entity.UserRole{
		UserID: user.ID,
		RoleID: role.ID,
	}
	db.Create(&userRole)

	return result.Error
}
