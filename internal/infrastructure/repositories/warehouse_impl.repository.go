package repositories

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	domainrepositories "github.com/Ardnh/be-warehouse-management/internal/domain/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WarehouseRepositoryImpl struct{ db *gorm.DB }

func NewWarehouseRepository(db *gorm.DB) domainrepositories.WarehouseRepository {
	return &WarehouseRepositoryImpl{db: db}
}

func (r *WarehouseRepositoryImpl) FindAll(ctx context.Context, filter domainrepositories.Filter) ([]entity.Warehouse, int64, error) {
	var items []entity.Warehouse
	var total int64
	base := r.db.WithContext(ctx).Model(&entity.Warehouse{})
	if filter.Search != "" {
		base = base.Where("code ILIKE ? OR name ILIKE ?", "%"+filter.Search+"%", "%"+filter.Search+"%")
	}
	if err := base.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	base = applyMasterListFilter(base, filter)
	if err := base.Find(&items).Error; err != nil {
		return nil, 0, err
	}
	return items, total, nil
}
func (r *WarehouseRepositoryImpl) FindByID(ctx context.Context, id uuid.UUID) (*entity.Warehouse, error) {
	var item entity.Warehouse
	if err := r.db.WithContext(ctx).Preload("Zones").First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}
func (r *WarehouseRepositoryImpl) Create(ctx context.Context, item entity.Warehouse) error {
	return r.db.WithContext(ctx).Create(&item).Error
}
func (r *WarehouseRepositoryImpl) Update(ctx context.Context, item *entity.Warehouse) error {
	return r.db.WithContext(ctx).Save(item).Error
}
func (r *WarehouseRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&entity.Warehouse{}, id).Error
}
