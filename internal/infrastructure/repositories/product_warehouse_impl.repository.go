package repositories

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"github.com/Ardnh/be-warehouse-management/internal/domain/repositories"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ProductWarehouseRepositoryImpl struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewProductWarehouseRepository(db *gorm.DB, redis *redis.Client) repositories.ProductWarehouseRepository {
	return &ProductWarehouseRepositoryImpl{
		db:    db,
		redis: redis,
	}
}

func (r *ProductWarehouseRepositoryImpl) FindByProductAndWarehouse(ctx context.Context, productID uuid.UUID, warehouseID uuid.UUID) (*entity.ProductWarehouse, error) {
	var productWarehouse entity.ProductWarehouse
	err := r.db.WithContext(ctx).Where("product_id = ? AND warehouse_id = ?", productID, warehouseID).First(&productWarehouse).Error
	if err != nil {
		return nil, err
	}
	return &productWarehouse, nil
}

func (r *ProductWarehouseRepositoryImpl) Create(ctx context.Context, customer entity.ProductWarehouse) error {
	err := r.db.WithContext(ctx).Create(&customer).Error
	return err
}
