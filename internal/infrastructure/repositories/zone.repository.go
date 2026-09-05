package repositories

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	domainrepositories "github.com/Ardnh/be-warehouse-management/internal/domain/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ZoneRepositoryImpl struct{ db *gorm.DB }

func NewZoneRepository(db *gorm.DB) domainrepositories.ZoneRepository {
	return &ZoneRepositoryImpl{db: db}
}
func (r *ZoneRepositoryImpl) FindAll(ctx context.Context, filter domainrepositories.Filter) ([]entity.Zone, int64, error) {
	var items []entity.Zone
	var total int64
	base := r.db.WithContext(ctx).Model(&entity.Zone{}).Preload("Warehouse")
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
func (r *ZoneRepositoryImpl) FindByID(ctx context.Context, id uuid.UUID) (*entity.Zone, error) {
	var item entity.Zone
	if err := r.db.WithContext(ctx).Preload("Warehouse").Preload("Racks").First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}
func (r *ZoneRepositoryImpl) Create(ctx context.Context, item entity.Zone) error {
	return r.db.WithContext(ctx).Create(&item).Error
}
func (r *ZoneRepositoryImpl) Update(ctx context.Context, item *entity.Zone) error {
	return r.db.WithContext(ctx).Save(item).Error
}
func (r *ZoneRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&entity.Zone{}, id).Error
}
