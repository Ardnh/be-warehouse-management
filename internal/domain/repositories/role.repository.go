package repositories

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"github.com/google/uuid"
)

type RoleRepository interface {
	FindAll(context.Context, Filter) ([]entity.Role, int64, error)
	FindByID(context.Context, uuid.UUID) (*entity.Role, error)
	Create(context.Context, entity.Role) error
	Update(context.Context, *entity.Role) error
	Delete(context.Context, uuid.UUID) error
}
