package repositories

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	domainrepositories "github.com/Ardnh/be-warehouse-management/internal/domain/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type HandlingUnitRepositoryImpl struct{ db *gorm.DB }

func NewHandlingUnitRepository(db *gorm.DB) domainrepositories.HandlingUnitRepository {
	return &HandlingUnitRepositoryImpl{db: db}
}
func (r *HandlingUnitRepositoryImpl) FindAll(ctx context.Context, filter domainrepositories.Filter) ([]entity.HandlingUnit, int64, error) {
	var list []entity.HandlingUnit
	var total int64
	q := r.db.WithContext(ctx).Model(&entity.HandlingUnit{}).Preload("Warehouse").Preload("Items").Preload("Items.Product")
	if filter.Search != "" {
		q = q.Where("code ILIKE ?", "%"+filter.Search+"%")
	}
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	q = applyMasterListFilter(q, filter)
	if err := q.Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}
func (r *HandlingUnitRepositoryImpl) FindByID(ctx context.Context, id uuid.UUID) (*entity.HandlingUnit, error) {
	var item entity.HandlingUnit
	if err := r.db.WithContext(ctx).Preload("Warehouse").Preload("Items").Preload("Items.Product").First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}
func (r *HandlingUnitRepositoryImpl) Create(ctx context.Context, item entity.HandlingUnit) error {
	return r.db.WithContext(ctx).Create(&item).Error
}
func (r *HandlingUnitRepositoryImpl) Update(ctx context.Context, item *entity.HandlingUnit) error {
	return r.db.WithContext(ctx).Save(item).Error
}
func (r *HandlingUnitRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&entity.HandlingUnit{}, id).Error
}
