package seeder

import (
	"errors"
	"time"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedUser(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {

		// --- User ---
		var user entity.User
		err := tx.Where("username = ?", "system-admin").First(&user).Error
		switch {
		case err == nil:
		case errors.Is(err, gorm.ErrRecordNotFound):
			hashed, hashErr := bcrypt.GenerateFromPassword([]byte("123123"), bcrypt.DefaultCost)
			if hashErr != nil {
				return hashErr
			}
			user = entity.User{
				Username:     "system-admin",
				Email:        "system-admin@example.com",
				PasswordHash: string(hashed),
				FullName:     "System Admin Uhuy",
				Status:       "ACTIVE",
				CreatedAt:    time.Now(),
			}
			if err := tx.Create(&user).Error; err != nil {
				return err
			}
		default:
			return err
		}

		// --- Role ---
		role := entity.Role{
			Name:        "system-admin",
			Code:        "SYSTEM_ADMIN",
			Description: "System admin role",
		}
		if err := tx.Where("code = ?", "SYSTEM_ADMIN").
			FirstOrCreate(&role, entity.Role{Code: "SYSTEM_ADMIN"}).Error; err != nil {
			return err
		}

		return nil
	})
}
