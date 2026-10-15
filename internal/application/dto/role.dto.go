package dto

import (
	"time"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"github.com/google/uuid"
)

// dto/role.go
type CreateRoleRequest struct {
	Code        string `json:"code" validate:"required,max=50"`
	Name        string `json:"name" validate:"required,max=100"`
	Description string `json:"description" validate:"omitempty"`
	Status        string      `json:"status" validate:"omitempty,oneof=ACTIVE INACTIVE"`
	PermissionIDs []uuid.UUID `json:"permission_ids" validate:"required,min=1,dive,required"`
}

type UpdateRoleRequest struct {
	Name        *string `json:"name" validate:"omitempty,max=100"`
	Description *string `json:"description"`
	Status        *string     `json:"status" validate:"omitempty,oneof=ACTIVE INACTIVE"`
	PermissionIDs []uuid.UUID `json:"permission_ids" validate:"omitempty,dive,required"`
}

type RoleResponse struct {
	ID          uuid.UUID `json:"id"`
	Code        string    `json:"code"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time             `json:"updated_at"`
	Permissions []*PermissionResponse `json:"permissions,omitempty"`
}

func ToRoleResponse(r *entity.Role) *RoleResponse {
	if r == nil {
		return nil
	}
	response := &RoleResponse{
		ID:          r.ID,
		Code:        r.Code,
		Name:        r.Name,
		Description: r.Description,
		Status:      r.Status,
		CreatedAt:   r.CreatedAt,
		UpdatedAt:   r.UpdatedAt,
	}
	for _, rolePermission := range r.RolePermissions {
		if rolePermission.Permission != nil {
			response.Permissions = append(response.Permissions, ToPermissionResponse(rolePermission.Permission))
		}
	}
	return response
}
