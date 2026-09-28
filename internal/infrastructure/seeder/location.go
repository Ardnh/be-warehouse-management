package seeder

import (
	"fmt"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"gorm.io/gorm"
)

func ptr[T any](v T) *T { return &v }

// SeedLocations inserts the head office and baseline warehouse locations.
// Existing locations are matched by their unique code and left unchanged.
func SeedLocations(db *gorm.DB) (*entity.Location, error) {
	locations := []entity.Location{
		{
			Code:       "HO",
			Name:       "Head Office",
			Type:       entity.LocationTypeHO,
			Address:    ptr("Jl. Jend. Sudirman Kav. 52-53"),
			City:       ptr("Jakarta Selatan"),
			Province:   ptr("DKI Jakarta"),
			PostalCode: ptr("12190"),
			IsActive:   true,
		},
		{
			Code:       "WH-JAKARTA",
			Name:       "Warehouse Jakarta",
			Type:       entity.LocationTypeWarehouse,
			Address:    ptr("Jl. Raya Cakung Cilincing No. 10"),
			City:       ptr("Jakarta Timur"),
			Province:   ptr("DKI Jakarta"),
			PostalCode: ptr("13910"),
			IsActive:   true,
		},
		{
			Code:       "WH-CIREBON",
			Name:       "Warehouse Cirebon",
			Type:       entity.LocationTypeWarehouse,
			Address:    ptr("Jl. Brigjen Dharsono No. 25"),
			City:       ptr("Cirebon"),
			Province:   ptr("Jawa Barat"),
			PostalCode: ptr("45153"),
			IsActive:   true,
		},
		{
			Code:       "WH-BANDUNG",
			Name:       "Warehouse Bandung",
			Type:       entity.LocationTypeWarehouse,
			Address:    ptr("Jl. Soekarno-Hatta No. 101"),
			City:       ptr("Bandung"),
			Province:   ptr("Jawa Barat"),
			PostalCode: ptr("40286"),
			IsActive:   true,
		},
		{
			Code:       "WH-YOGYAKARTA",
			Name:       "Warehouse Yogyakarta",
			Type:       entity.LocationTypeWarehouse,
			Address:    ptr("Jl. Ring Road Utara No. 45"),
			City:       ptr("Sleman"),
			Province:   ptr("DI Yogyakarta"),
			PostalCode: ptr("55283"),
			IsActive:   true,
		},
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
