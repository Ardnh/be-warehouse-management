package services

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/application/dto"
	"github.com/google/uuid"
)

type RoleService interface {
	FindAll(context.Context, dto.FilterDTO) ([]dto.RoleResponse, int64, error)
	FindByID(context.Context, uuid.UUID) (*dto.RoleResponse, error)
	Create(context.Context, dto.CreateRoleRequest) error
	Update(context.Context, uuid.UUID, dto.UpdateRoleRequest) error
	Delete(context.Context, uuid.UUID) error
}
