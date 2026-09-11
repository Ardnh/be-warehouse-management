package repositories

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	domainrepositories "github.com/Ardnh/be-warehouse-management/internal/domain/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UomRepositoryImpl struct{ db *gorm.DB }

func NewUomRepository(db *gorm.DB) domainrepositories.UomRepository {
	return &UomRepositoryImpl{db: db}
}

func (r *UomRepositoryImpl) FindAll(ctx context.Context, filter domainrepositories.Filter) ([]entity.Uom, int64, error) {
	var items []entity.Uom
	var total int64
	base := Conn(ctx, r.db).Model(&entity.Uom{})
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

func (r *UomRepositoryImpl) FindByID(ctx context.Context, id uuid.UUID) (*entity.Uom, error) {
	var item entity.Uom
	if err := Conn(ctx, r.db).First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *UomRepositoryImpl) Create(ctx context.Context, item entity.Uom) error {
	return Conn(ctx, r.db).Create(&item).Error
}

func (r *UomRepositoryImpl) Update(ctx context.Context, item *entity.Uom) error {
	return Conn(ctx, r.db).Save(item).Error
}

func (r *UomRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return Conn(ctx, r.db).Delete(&entity.Uom{}, id).Error
}
