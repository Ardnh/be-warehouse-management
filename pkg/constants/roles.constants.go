package constants

const (
	RoleAdminPlatform       = "admin-platform"
	RoleWarehouseAdmin      = "warehouse-admin"
	RoleReceivingOperator   = "receiving-operator"
	RolePutawayOperator     = "putaway-operator"
	RolePickingOperator     = "picking-operator"
	RoleDispatchOperator    = "dispatch-operator"
	RoleWarehouseSupervisor = "warehouse-supervisor"
)

var ValidRoles = map[string]bool{
	RoleAdminPlatform:       true,
	RoleWarehouseSupervisor: true,
	RoleWarehouseAdmin:      true,
	RoleReceivingOperator:   true,
	RolePutawayOperator:     true,
	RolePickingOperator:     true,
	RoleDispatchOperator:    true,
}

func IsValidRole(role string) bool {
	return ValidRoles[role]
}
