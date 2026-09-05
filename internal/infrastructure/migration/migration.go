package migration

import (
	"fmt"
	"log"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	log.Println("Running migrations...")

	err := db.AutoMigrate(
		&entity.User{},
		&entity.Role{},
		&entity.UserRole{},

		// Master data
		&entity.Customer{},
		&entity.Product{},
		&entity.Warehouse{},
		&entity.Zone{},
		&entity.Uom{},
		&entity.Rack{},
		&entity.StorageLocation{},

		// Bussiness Process
		&entity.InboundOrder{},
		&entity.InboundOrderItem{},
		&entity.Receiving{},
		&entity.ReceivingItem{},
		&entity.HandlingUnit{},
		&entity.HandlingUnitItem{},
	)

	if err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}

	if err := db.Exec(`
		CREATE SEQUENCE IF NOT EXISTS customer_code_seq
		START WITH 1
		INCREMENT BY 1
	`).Error; err != nil {
		return err
	}

	log.Println("Migration completed successfully")
	return nil
}
