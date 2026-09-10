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
	Email    *string `json:"email" validate:"omitempty,email,max=150"`
	FullName *string `json:"full_name" validate:"omitempty,max=150"`
	Status   *string `json:"status" validate:"omitempty,oneof=ACTIVE INACTIVE"`
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
	ID        uuid.UUID      `json:"id"`
	Username  string         `json:"username"`
	Email     string         `json:"email"`
	FullName  string         `json:"full_name"`
	Status    string         `json:"status"`
	Roles     []RoleResponse `json:"roles,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
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
	return res
}

func ToUserDTOs(users []*entity.User) []User {
	res := make([]User, 0, len(users))
	for _, u := range users {
		res = append(res, ToUserDTO(u))
	}
	return res
}
