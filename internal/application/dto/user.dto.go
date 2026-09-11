package dto

import (
	"time"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"github.com/google/uuid"
)

// dto/user.go
type CreateUserRequest struct {
	Username string      `json:"username" validate:"required,alphanum,min=3,max=50"`
	Email    string      `json:"email" validate:"required,email,max=150"`
	Password string      `json:"password" validate:"required,min=8,max=72"`
	FullName string      `json:"full_name" validate:"required,max=150"`
	Status   string      `json:"status" validate:"omitempty,oneof=ACTIVE INACTIVE"`
	RoleIDs  []uuid.UUID `json:"role_ids" validate:"omitempty,dive,required"`
}

type UpdateUserRequest struct {
	Username *string     `json:"username" validate:"omitempty,alphanum,min=3,max=50"`
	Email    *string     `json:"email" validate:"omitempty,email,max=150"`
	FullName *string     `json:"full_name" validate:"omitempty,max=150"`
	Status   *string     `json:"status" validate:"omitempty,oneof=ACTIVE INACTIVE"`
	RoleIDs  []uuid.UUID `json:"role_ids" validate:"omitempty,dive,required"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required,min=8,max=72,nefield=OldPassword"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AssignRolesRequest struct {
	RoleIDs []uuid.UUID `json:"role_ids" validate:"required,min=1,dive,required"`
}

type User struct {
	ID              uuid.UUID                `json:"id"`
	Username        string                   `json:"username"`
	Email           string                   `json:"email"`
	FullName        string                   `json:"full_name"`
	Status          string                   `json:"status"`
	Roles           []RoleResponse           `json:"roles,omitempty"`
	UserRoles       []UserRoleResponse       `json:"user_roles,omitempty"`
	UserPermissions []UserPermissionResponse `json:"user_permissions,omitempty"`
	CreatedAt       time.Time                `json:"created_at"`
	UpdatedAt       time.Time                `json:"updated_at"`
}

type UserRoleResponse struct {
	RoleID      uuid.UUID     `json:"role_id"`
	WarehouseID uuid.UUID     `json:"warehouse_id"`
	Role        *RoleResponse `json:"role,omitempty"`
	Warehouse   *Warehouse    `json:"warehouse,omitempty"`
	CreatedAt   time.Time     `json:"created_at"`
}

type UserPermissionResponse struct {
	PermissionID uuid.UUID      `json:"permission_id"`
	Permission   *PermissionDTO `json:"permission,omitempty"`
}

func ToUserDTO(u *entity.User) User {
	res := User{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		FullName:  u.FullName,
		Status:    u.Status,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
	for _, r := range u.Roles {
		res.Roles = append(res.Roles, NewRoleResponse(r))
	}
	for _, userRole := range u.UserRoles {
		userRoleResponse := UserRoleResponse{
			RoleID:      userRole.RoleID,
			WarehouseID: userRole.WarehouseID,
			CreatedAt:   userRole.CreatedAt,
		}
		if userRole.Role != nil {
			roleResponse := NewRoleResponse(*userRole.Role)
			userRoleResponse.Role = &roleResponse
		}
		if userRole.Warehouse != nil {
			warehouseResponse := NewWarehouseResponse(*userRole.Warehouse)
			userRoleResponse.Warehouse = &warehouseResponse
		}
		res.UserRoles = append(res.UserRoles, userRoleResponse)
	}
	for _, userPermission := range u.UserPermissions {
		userPermissionResponse := UserPermissionResponse{
			PermissionID: userPermission.PermissionID,
		}
		if userPermission.Permission != nil {
			userPermissionResponse.Permission = ToPermissionDTO(userPermission.Permission)
		}
		res.UserPermissions = append(res.UserPermissions, userPermissionResponse)
	}
	return res
}

func ToUserDTOs(users []*entity.User) []User {
	res := make([]User, 0, len(users))
	for _, u := range users {
		res = append(res, ToUserDTO(u))
	}
	return res
}
