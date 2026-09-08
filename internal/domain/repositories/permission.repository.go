package repositories

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"github.com/google/uuid"
)

type PermissionRepository interface {
	FindAll(context.Context, Filter) ([]entity.Permission, int64, error)
	FindByID(context.Context, uuid.UUID) (*entity.Permission, error)
	Create(context.Context, entity.Permission) error
	Update(context.Context, *entity.Permission) error
	Delete(context.Context, uuid.UUID) error
}
