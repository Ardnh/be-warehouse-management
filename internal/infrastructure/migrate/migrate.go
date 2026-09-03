package migrate

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
	)

	if err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}

	log.Println("Migration completed successfully")
	return nil
}
