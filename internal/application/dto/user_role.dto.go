package dto

import (
	"time"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
)

type UserRoleDto struct {
	UserID      string    `json:"user_id"`
	RoleID      string    `json:"role_id"`
	WarehouseID string    `json:"warehouse_id"`
	CreatedAt   time.Time `json:"created_at"`
	Status      string    `json:"status"`

	User      *User         `json:"user,omitempty"`
	Role      *RoleResponse `json:"role,omitempty"`
	Warehouse *Warehouse    `json:"warehouse,omitempty"`
}

func ToUserRoleDto(userRole *entity.UserRole) *UserRoleDto {
	if userRole == nil {
		return nil
	}

	result := &UserRoleDto{
		UserID:      userRole.UserID.String(),
		RoleID:      userRole.RoleID.String(),
		WarehouseID: userRole.WarehouseID.String(),
		CreatedAt:   userRole.CreatedAt,
		Status:      userRole.Status,
	}

	if userRole.User != nil {
		user := ToUserDTO(userRole.User)
		result.User = &user
	}
	if userRole.Role != nil {
		role := NewRoleResponse(*userRole.Role)
		result.Role = &role
	}
	if userRole.Warehouse != nil {
		warehouse := NewWarehouseResponse(*userRole.Warehouse)
		result.Warehouse = &warehouse
	}

	return result
}

func ToUserRoleDtos(userRoles []entity.UserRole) []*UserRoleDto {
	result := make([]*UserRoleDto, 0, len(userRoles))
	for _, userRole := range userRoles {
		result = append(result, ToUserRoleDto(&userRole))
	}
	return result
}
