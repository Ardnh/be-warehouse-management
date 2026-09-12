package repositories

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	domainrepositories "github.com/Ardnh/be-warehouse-management/internal/domain/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RackRepositoryImpl struct{ db *gorm.DB }

func NewRackRepository(db *gorm.DB) domainrepositories.RackRepository {
	return &RackRepositoryImpl{db: db}
}

func (r *RackRepositoryImpl) FindAll(ctx context.Context, filter domainrepositories.Filter) ([]entity.Rack, int64, error) {
	var items []entity.Rack
	var total int64
	base := Conn(ctx, r.db).Model(&entity.Rack{}).Preload("Zone")
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

func (r *RackRepositoryImpl) FindByID(ctx context.Context, id uuid.UUID) (*entity.Rack, error) {
	var item entity.Rack
	if err := Conn(ctx, r.db).Preload("Zone").Preload("Locations").First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *RackRepositoryImpl) Create(ctx context.Context, item entity.Rack) error {
	return Conn(ctx, r.db).Create(&item).Error
}

func (r *RackRepositoryImpl) Update(ctx context.Context, item *entity.Rack) error {
	return Conn(ctx, r.db).Save(item).Error
}

func (r *RackRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return Conn(ctx, r.db).Delete(&entity.Rack{}, id).Error
}
