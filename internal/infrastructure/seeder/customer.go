package seeder

import (
	"fmt"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"gorm.io/gorm"
)

// SeedCustomers inserts the baseline customers used by local development.
// Existing customers are matched by their unique code and left unchanged.
func SeedCustomers(db *gorm.DB) error {
	customers := []entity.Customer{
		{
			Code:    "CUST-001",
			Name:    "Acme Distribution",
			Email:   "acme@example.com",
			Address: "123 Main Street",
			Status:  "active",
			Phone:   "+1-555-0100",
		},
		{
			Code:    "CUST-002",
			Name:    "Northstar Retail",
			Email:   "northstar@example.com",
			Address: "456 Market Avenue",
			Status:  "active",
			Phone:   "+1-555-0101",
		},
		{
			Code:    "CUST-003",
			Name:    "Summit Wholesale",
			Email:   "summit@example.com",
			Address: "789 Industrial Road",
			Status:  "active",
			Phone:   "+1-555-0102",
		},
	}

	for i := range customers {
		customer := &customers[i]
		var existing entity.Customer
		result := db.Where("code = ?", customer.Code).First(&existing)
		if result.Error == nil {
			continue
		}
		if result.Error != gorm.ErrRecordNotFound {
			return fmt.Errorf("find customer %s: %w", customer.Code, result.Error)
		}
		if err := db.Create(customer).Error; err != nil {
			return fmt.Errorf("create customer %s: %w", customer.Code, err)
		}
	}

	return nil
}
