package seeder

import (
	"fmt"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"gorm.io/gorm"
)

// SeedLocations inserts the head office and baseline warehouse locations.
// Existing locations are matched by their unique code and left unchanged.
func SeedLocations(db *gorm.DB) (*entity.Location, error) {
	locations := []entity.Location{
		{Code: "HO", Name: "HO", Type: entity.LocationTypeHO, IsActive: true},
		{Code: "WH-JAKARTA", Name: "Warehouse Jakarta", Type: entity.LocationTypeWarehouse, IsActive: true},
		{Code: "WH-CIREBON", Name: "Warehouse Cirebon", Type: entity.LocationTypeWarehouse, IsActive: true},
		{Code: "WH-BANDUNG", Name: "Warehouse Bandung", Type: entity.LocationTypeWarehouse, IsActive: true},
		{Code: "WH-YOGYAKARTA", Name: "Warehouse Yogyakarta", Type: entity.LocationTypeWarehouse, IsActive: true},
	}

	for i := range locations {
		location := &locations[i]
		if err := db.Where("code = ?", location.Code).FirstOrCreate(location).Error; err != nil {
			return nil, fmt.Errorf("seed location %s: %w", location.Code, err)
		}
	}

	var location entity.Location
	if err := db.Where("code = ?", "HO").First(&location).Error; err != nil {
		return nil, fmt.Errorf("get head office: %w", err)
	}

	return &location, nil
}
