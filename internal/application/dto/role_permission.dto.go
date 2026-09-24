package dto

import "github.com/Ardnh/be-warehouse-management/internal/domain/entity"

type RolePermissionResponse struct {
	RoleID       string `json:"role_id"`
	PermissionID string `json:"permission_id"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`

	Role       *RoleResponse       `json:"role"`
	Permission *PermissionResponse `json:"permission"`
}

func ToRolePermissionsResponse(permissions []*entity.RolePermission) []*PermissionResponse {
	result := make([]*PermissionResponse, 0, len(permissions))

	for i := range permissions {
		result = append(result, ToPermissionResponse(permissions[i].Permission))
	}

	return result
}

func ToFormattedRolePermissionsResponse(permissions []*entity.RolePermission) []string {
	result := make([]string, 0, len(permissions))

	for i := range permissions {
		result = append(result, ToFormattedPermissionResponse(permissions[i].Permission))
	}

	return result
}
