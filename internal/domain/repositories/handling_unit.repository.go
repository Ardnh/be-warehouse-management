package repositories

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"github.com/google/uuid"
)

type HandlingUnitRepository interface {
	FindAll(ctx context.Context, filter Filter) ([]entity.HandlingUnit, int64, error)
	FindByID(ctx context.Context, id uuid.UUID) (*entity.HandlingUnit, error)
	Create(ctx context.Context, unit entity.HandlingUnit) error
	Update(ctx context.Context, unit *entity.HandlingUnit) error
	Delete(ctx context.Context, id uuid.UUID) error
}
