package seeder

import (
	"fmt"

	"github.com/Ardnh/be-warehouse-management/pkg/constants"
	"github.com/casbin/casbin/v3"
)

type Permission struct {
	Resource string
	Action   string
}

func SeedCasbinRules(enforcer *casbin.Enforcer) error {

	existingPolicies, _ := enforcer.GetPolicy()
	if len(existingPolicies) > 0 {
		return nil
	}

	permissions := []Permission{
		{Resource: constants.ResourceUser, Action: constants.ActionRead},
		{Resource: constants.ResourceUser, Action: constants.ActionCreate},
		{Resource: constants.ResourceUser, Action: constants.ActionUpdate},
		{Resource: constants.ResourceUser, Action: constants.ActionDelete},

		{Resource: constants.ResourceCustomer, Action: constants.ActionRead},
		{Resource: constants.ResourceCustomer, Action: constants.ActionCreate},
		{Resource: constants.ResourceCustomer, Action: constants.ActionUpdate},
		{Resource: constants.ResourceCustomer, Action: constants.ActionDelete},

		{Resource: constants.ResourceWarehouse, Action: constants.ActionRead},
		{Resource: constants.ResourceWarehouse, Action: constants.ActionCreate},
		{Resource: constants.ResourceWarehouse, Action: constants.ActionUpdate},
		{Resource: constants.ResourceWarehouse, Action: constants.ActionDelete},

		{Resource: constants.ResourceUom, Action: constants.ActionRead},
		{Resource: constants.ResourceUom, Action: constants.ActionCreate},
		{Resource: constants.ResourceUom, Action: constants.ActionUpdate},
		{Resource: constants.ResourceUom, Action: constants.ActionDelete},

		{Resource: constants.ResourceProduct, Action: constants.ActionRead},
		{Resource: constants.ResourceProduct, Action: constants.ActionCreate},
		{Resource: constants.ResourceProduct, Action: constants.ActionUpdate},
		{Resource: constants.ResourceProduct, Action: constants.ActionDelete},

		{Resource: constants.ResourceRole, Action: constants.ActionRead},
		{Resource: constants.ResourceRole, Action: constants.ActionCreate},
		{Resource: constants.ResourceRole, Action: constants.ActionUpdate},
		{Resource: constants.ResourceRole, Action: constants.ActionDelete},

		{Resource: constants.ResourceReceiving, Action: constants.ActionRead},
		{Resource: constants.ResourceReceiving, Action: constants.ActionCreate},
		{Resource: constants.ResourceReceiving, Action: constants.ActionUpdate},
		{Resource: constants.ResourceReceiving, Action: constants.ActionDelete},
	}

	_, err := enforcer.AddPolicies(permissions)
	if err != nil {
		return fmt.Errorf("gagal seed policies: %w", err)
	}

	// === Role Hierarchy ===
	for childRole, parentRole := range constants.RoleHierarchy {
		_, err := enforcer.AddGroupingPolicy(childRole, parentRole)
		if err != nil {
			return fmt.Errorf("gagal seed hierarchy %s -> %s: %w", childRole, parentRole, err)
		}
	}
	return nil
}
