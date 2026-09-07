package seeder

import (
	"fmt"
	"strings"

	"github.com/Ardnh/be-warehouse-management/pkg/constants"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Permission struct {
	Resource    string
	Action      string
	Description string
}

// semua resource yang punya CRUD standar
var permissionResources = []string{
	constants.ResourceUser,
	constants.ResourceUserRole,
	constants.ResourceUserPermission,
	constants.ResourceCustomer,
	constants.ResourceWarehouse,
	constants.ResourceUom,
	constants.ResourceZone,
	constants.ResourceRack,
	constants.ResourceStorageLocation,
	constants.ResourceInboundOrder,
	constants.ResourceInboundOrderItem,
	constants.ResourceHandlingUnit,
	constants.ResourceHandlingUnitItem,
	constants.ResourceProduct,
	constants.ResourceRole,
	constants.ResourceReceiving,
	constants.ResourceReceivingItem,
}

var crudActions = []string{
	constants.ActionRead,
	constants.ActionCreate,
	constants.ActionUpdate,
	constants.ActionDelete,
}

// resource yang boleh di-export/download (sesuaikan dengan kebutuhan)
var downloadableResources = map[string]bool{
	constants.ResourceUser:         true,
	constants.ResourceCustomer:     true,
	constants.ResourceProduct:      true,
	constants.ResourceInboundOrder: true,
	constants.ResourceReceiving:    true,
	constants.ResourceHandlingUnit: true,
}

func describe(action, resource string) string {
	label := strings.ReplaceAll(resource, "-", " ")
	switch action {
	case constants.ActionRead:
		return "Melihat data " + label
	case constants.ActionCreate:
		return "Membuat data " + label
	case constants.ActionUpdate:
		return "Mengubah data " + label
	case constants.ActionDelete:
		return "Menghapus data " + label
	case constants.ActionDownload:
		return "Mengunduh data " + label
	default:
		return action + " " + label
	}
}

func buildPermissions() []Permission {
	perms := make([]Permission, 0, len(permissionResources)*len(crudActions)+len(downloadableResources))

	for _, resource := range permissionResources {
		for _, action := range crudActions {
			perms = append(perms, Permission{Resource: resource, Action: action, Description: describe(action, resource)})
		}
		if downloadableResources[resource] {
			perms = append(perms, Permission{Resource: resource, Action: constants.ActionDownload, Description: describe(constants.ActionDownload, resource)})
		}
	}

	return perms
}

func SeedPermissions(db *gorm.DB) error {
	permissions := buildPermissions()

	err := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "resource"}, {Name: "action"}},
		DoNothing: true,
	}).CreateInBatches(&permissions, 100).Error
	if err != nil {
		return fmt.Errorf("gagal seed permissions: %w", err)
	}

	return nil
}
