package repositories

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"github.com/google/uuid"
)

type CustomerWarehouseRepository interface {
	FindByCustomerAndWarehouse(ctx context.Context, customerID uuid.UUID, warehouseID uuid.UUID) (*entity.CustomerWarehouse, error)
	Create(ctx context.Context, customer entity.CustomerWarehouse) error
}
