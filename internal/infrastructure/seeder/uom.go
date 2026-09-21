package seeder

import (
	"fmt"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"gorm.io/gorm"
)

// SeedUoms inserts the standard units of measure used by local development.
// Existing UOMs are matched by their unique code and left unchanged.
func SeedUoms(db *gorm.DB) error {
	uoms := []entity.Uom{
		{Code: "PCS", Name: "Pieces", Type: "COUNT"},
		{Code: "BOX", Name: "Box", Type: "PACKAGING"},
		{Code: "KG", Name: "Kilogram", Type: "WEIGHT"},
		{Code: "PALLET", Name: "Pallet", Type: "PACKAGING"},
	}

	for i := range uoms {
		uom := &uoms[i]
		var existing entity.Uom
		result := db.Where("code = ?", uom.Code).First(&existing)
		if result.Error == nil {
			continue
		}
		if result.Error != gorm.ErrRecordNotFound {
			return fmt.Errorf("find UOM %s: %w", uom.Code, result.Error)
		}
		if err := db.Create(uom).Error; err != nil {
			return fmt.Errorf("create UOM %s: %w", uom.Code, err)
		}
	}

	return nil
}
