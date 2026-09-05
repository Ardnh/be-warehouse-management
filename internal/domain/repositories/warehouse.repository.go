package repositories

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"github.com/google/uuid"
)

type WarehouseRepository interface {
	FindAll(ctx context.Context, filter Filter) ([]entity.Warehouse, int64, error)
	FindByID(ctx context.Context, id uuid.UUID) (*entity.Warehouse, error)
	Create(ctx context.Context, warehouse entity.Warehouse) error
	Update(ctx context.Context, warehouse *entity.Warehouse) error
	Delete(ctx context.Context, id uuid.UUID) error
}
