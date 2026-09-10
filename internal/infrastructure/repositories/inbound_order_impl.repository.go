package repositories

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	domainrepositories "github.com/Ardnh/be-warehouse-management/internal/domain/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InboundOrderRepositoryImpl struct{ db *gorm.DB }

func NewInboundOrderRepository(db *gorm.DB) domainrepositories.InboundOrderRepository {
	return &InboundOrderRepositoryImpl{db: db}
}
func (r *InboundOrderRepositoryImpl) FindAll(ctx context.Context, filter domainrepositories.Filter) ([]entity.InboundOrder, int64, error) {
	var list []entity.InboundOrder
	var total int64
	q := r.db.WithContext(ctx).Model(&entity.InboundOrder{}).Preload("Customer").Preload("Items")
	if filter.Search != "" {
		q = q.Where("order_number ILIKE ?", "%"+filter.Search+"%")
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
func (r *InboundOrderRepositoryImpl) FindByID(ctx context.Context, id uuid.UUID) (*entity.InboundOrder, error) {
	var item entity.InboundOrder
	if err := r.db.WithContext(ctx).Preload("Customer").Preload("Items").First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}
func (r *InboundOrderRepositoryImpl) Create(ctx context.Context, item entity.InboundOrder) error {
	return r.db.WithContext(ctx).Session(&gorm.Session{FullSaveAssociations: true}).Create(&item).Error
}
func (r *InboundOrderRepositoryImpl) Update(ctx context.Context, item *entity.InboundOrder) error {
	return r.db.WithContext(ctx).Save(item).Error
}
func (r *InboundOrderRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&entity.InboundOrder{}, id).Error
}
