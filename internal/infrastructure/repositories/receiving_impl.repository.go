package repositories

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	domainrepositories "github.com/Ardnh/be-warehouse-management/internal/domain/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReceivingRepositoryImpl struct{ db *gorm.DB }

func NewReceivingRepository(db *gorm.DB) domainrepositories.ReceivingRepository {
	return &ReceivingRepositoryImpl{db: db}
}

func (r *ReceivingRepositoryImpl) FindAll(ctx context.Context, filter domainrepositories.Filter) ([]entity.Receiving, int64, error) {
	var items []entity.Receiving
	var total int64
	q := Conn(ctx, r.db).Model(&entity.Receiving{}).Preload("Items")
	if filter.Search != "" {
		q = q.Where("receiving_number ILIKE ?", "%"+filter.Search+"%")
	}
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := applyMasterListFilter(q, filter).Find(&items).Error; err != nil {
		return nil, 0, err
	}
	return items, total, nil
}

func (r *ReceivingRepositoryImpl) FindByID(ctx context.Context, id uuid.UUID) (*entity.Receiving, error) {
	var item entity.Receiving
	if err := Conn(ctx, r.db).Preload("Items").First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *ReceivingRepositoryImpl) Create(ctx context.Context, item entity.Receiving) error {
	return Conn(ctx, r.db).Session(&gorm.Session{FullSaveAssociations: true}).Create(&item).Error
}

func (r *ReceivingRepositoryImpl) Update(ctx context.Context, item *entity.Receiving) error {
	return Conn(ctx, r.db).Save(item).Error
}

func (r *ReceivingRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return Conn(ctx, r.db).Delete(&entity.Receiving{}, id).Error
}
