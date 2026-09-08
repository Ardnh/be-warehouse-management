package casbin

import (
	"github.com/casbin/casbin/v3"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

func InitCasbin(modelPath string, db *gorm.DB) (*casbin.SyncedEnforcer, error) {
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		return nil, err
	}
	return casbin.NewSyncedEnforcer(modelPath, adapter)
}

func GetUserPermissions(
	enforcer *casbin.Enforcer,
	userID string,
	warehouseID string,
) []string {
	permissions, err := enforcer.GetImplicitPermissionsForUser(
		userID,
		warehouseID,
	)

	if err != nil {
		return []string{}
	}

	seen := make(map[string]bool)
	result := []string{}

	for _, p := range permissions {
		// p = [subject, domain, object, action]
		if len(p) < 4 {
			continue
		}

		perm := p[2] + ":" + p[3]

		if !seen[perm] {
			seen[perm] = true
			result = append(result, perm)
		}
	}

	return result
}

func GetUserRoles(
	enforcer *casbin.Enforcer,
	userID string,
	warehouseID string,
) []string {
	roles, err := enforcer.GetRolesForUser(
		userID,
		warehouseID,
	)

	if err != nil {
		return []string{}
	}

	return roles
}
