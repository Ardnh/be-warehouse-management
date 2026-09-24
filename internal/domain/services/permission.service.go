package services

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/application/dto"
	"github.com/google/uuid"
)

type PermissionService interface {
	FindAll(context.Context, dto.FilterDTO) ([]*dto.PermissionResponse, int64, error)
	FindByID(context.Context, uuid.UUID) (*dto.PermissionResponse, error)
	HasPermission(ctx context.Context, userID uuid.UUID, resource string, action string) (bool, error)
	Create(context.Context, dto.CreatePermissionRequest) error
	Update(context.Context, uuid.UUID, dto.UpdatePermissionRequest) error
	Delete(context.Context, uuid.UUID) error
}
