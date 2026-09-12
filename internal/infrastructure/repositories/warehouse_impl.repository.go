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
	base := Conn(ctx, r.db).Model(&entity.Warehouse{})
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
	if err := Conn(ctx, r.db).Preload("Zones").First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *WarehouseRepositoryImpl) Create(ctx context.Context, item entity.Warehouse) error {
	return Conn(ctx, r.db).Create(&item).Error
}

func (r *WarehouseRepositoryImpl) Update(ctx context.Context, item *entity.Warehouse) error {
	return Conn(ctx, r.db).Save(item).Error
}

func (r *WarehouseRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return Conn(ctx, r.db).Delete(&entity.Warehouse{}, id).Error
}
