package dto

import (
	"fmt"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
)

type CreatePermissionRequest struct {
	Resource    string `json:"resource" validate:"required"`
	Action      string `json:"action" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type UpdatePermissionRequest struct {
	Resource    string `json:"resource"`
	Action      string `json:"action"`
	Description string `json:"description"`
}

type PermissionResponse struct {
	ID          string `json:"id"`
	Resource    string `json:"resource"`
	Action      string `json:"action"`
	Description string `json:"description"`
}

type PermissionGroupResponse struct {
	Group  string                `json:"group"`
	Action []*PermissionResponse `json:"action"`
}

func ToPermissionResponse(permission *entity.Permission) *PermissionResponse {
	return &PermissionResponse{
		ID:          permission.ID.String(),
		Resource:    permission.Resource,
		Action:      permission.Action,
		Description: permission.Description,
	}
}

func ToFormattedPermissionResponse(permission *entity.Permission) string {
	return fmt.Sprintf("%s:%s:%s", permission.Resource, permission.Action, permission.Description)
}

func ToFormattedPermissionResponses(permissions []*entity.Permission) []string {
	result := make([]string, 0, len(permissions))
	for _, permission := range permissions {
		result = append(result, ToFormattedPermissionResponse(permission))
	}
	return result
}

func ToPermissionsResponse(permissions []*entity.Permission) []*PermissionResponse {
	result := make([]*PermissionResponse, 0, len(permissions))
	for _, permission := range permissions {
		result = append(result, ToPermissionResponse(permission))
	}
	return result
}
