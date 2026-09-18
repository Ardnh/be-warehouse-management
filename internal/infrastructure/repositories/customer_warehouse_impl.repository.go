package repositories

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"github.com/Ardnh/be-warehouse-management/internal/domain/repositories"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type CustomerWarehouseRepositoryImpl struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewCustomerWarehouseRepository(db *gorm.DB, redis *redis.Client) repositories.CustomerWarehouseRepository {
	return &CustomerWarehouseRepositoryImpl{
		db:    db,
		redis: redis,
	}
}

func (r *CustomerWarehouseRepositoryImpl) FindByCustomerAndWarehouse(ctx context.Context, customerID uuid.UUID, warehouseID uuid.UUID) (*entity.CustomerWarehouse, error) {
	var customer entity.CustomerWarehouse
	err := r.db.WithContext(ctx).Where("customer_id = ? AND warehouse_id = ?", customerID, warehouseID).First(&customer).Error
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *CustomerWarehouseRepositoryImpl) Create(ctx context.Context, customer entity.CustomerWarehouse) error {
	err := r.db.WithContext(ctx).Create(&customer).Error
	return err
}
