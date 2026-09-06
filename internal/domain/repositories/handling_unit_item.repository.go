package repositories

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"github.com/google/uuid"
)

type HandlingUnitItemRepository interface {
	FindAllByHandlingUnit(ctx context.Context, handlingUnitID uuid.UUID) ([]entity.HandlingUnitItem, error)
	FindByID(ctx context.Context, id uuid.UUID) (*entity.HandlingUnitItem, error)
	Create(ctx context.Context, item entity.HandlingUnitItem) error
	Update(ctx context.Context, item *entity.HandlingUnitItem) error
	Delete(ctx context.Context, id uuid.UUID) error
}
