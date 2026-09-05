package repositories

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	domainrepositories "github.com/Ardnh/be-warehouse-management/internal/domain/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StorageLocationRepositoryImpl struct{ db *gorm.DB }

func NewStorageLocationRepository(db *gorm.DB) domainrepositories.StorageLocationRepository {
	return &StorageLocationRepositoryImpl{db: db}
}
func (r *StorageLocationRepositoryImpl) FindAll(ctx context.Context, filter domainrepositories.Filter) ([]entity.StorageLocation, int64, error) {
	var items []entity.StorageLocation
	var total int64
	base := r.db.WithContext(ctx).Model(&entity.StorageLocation{}).Preload("Rack")
	if filter.Search != "" {
		base = base.Where("code ILIKE ?", "%"+filter.Search+"%")
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
func (r *StorageLocationRepositoryImpl) FindByID(ctx context.Context, id uuid.UUID) (*entity.StorageLocation, error) {
	var item entity.StorageLocation
	if err := r.db.WithContext(ctx).Preload("Rack").First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}
func (r *StorageLocationRepositoryImpl) Create(ctx context.Context, item entity.StorageLocation) error {
	return r.db.WithContext(ctx).Create(&item).Error
}
func (r *StorageLocationRepositoryImpl) Update(ctx context.Context, item *entity.StorageLocation) error {
	return r.db.WithContext(ctx).Save(item).Error
}
func (r *StorageLocationRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&entity.StorageLocation{}, id).Error
}
