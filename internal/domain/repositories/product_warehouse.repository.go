package repositories

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"github.com/google/uuid"
)

type ProductWarehouseRepository interface {
	FindByProductAndWarehouse(ctx context.Context, productID uuid.UUID, warehouseID uuid.UUID) (*entity.ProductWarehouse, error)
	Create(ctx context.Context, product entity.ProductWarehouse) error
}
