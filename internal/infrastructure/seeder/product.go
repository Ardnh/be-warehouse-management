package seeder

import (
	"fmt"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"gorm.io/gorm"
)

// SeedProducts inserts baseline products for the seeded customers.
// Customer and UOM records are resolved by their unique business codes.
func SeedProducts(db *gorm.DB) error {
	seedProducts := []struct {
		CustomerCode string
		UomCode      string
		SKU          string
		Name         string
		Barcode      string
		Status       string
	}{
		{
			CustomerCode: "CUST-001",
			UomCode:      "PCS",
			SKU:          "ACME-001",
			Name:         "Standard carton",
			Barcode:      "100000000001",
			Status:       "active",
		},
		{
			CustomerCode: "CUST-001",
			UomCode:      "PCS",
			SKU:          "ACME-002",
			Name:         "Heavy-duty carton",
			Barcode:      "100000000002",
			Status:       "active",
		},
		{
			CustomerCode: "CUST-002",
			UomCode:      "PCS",
			SKU:          "NSTAR-001",
			Name:         "Retail display box",
			Barcode:      "100000000003",
			Status:       "active",
		},
		{
			CustomerCode: "CUST-003",
			UomCode:      "PCS",
			SKU:          "SUMMIT-001",
			Name:         "Bulk storage container",
			Barcode:      "100000000004",
			Status:       "active",
		},
	}

	for _, seed := range seedProducts {
		var customer entity.Customer
		if err := db.Where("code = ?", seed.CustomerCode).First(&customer).Error; err != nil {
			return fmt.Errorf("find customer %s for product %s: %w", seed.CustomerCode, seed.SKU, err)
		}

		var uom entity.Uom
		if err := db.Where("code = ?", seed.UomCode).First(&uom).Error; err != nil {
			return fmt.Errorf("find UOM %s for product %s: %w", seed.UomCode, seed.SKU, err)
		}

		barcode := seed.Barcode
		product := entity.Product{
			CustomerID: customer.ID,
			SKU:        seed.SKU,
			Name:       seed.Name,
			UomID:      uom.ID,
			Barcode:    &barcode,
			Status:     seed.Status,
		}

		if err := db.Where("sku = ?", seed.SKU).FirstOrCreate(&product).Error; err != nil {
			return fmt.Errorf("create product %s: %w", seed.SKU, err)
		}
	}

	return nil
}
