package repositories

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	domainrepositories "github.com/Ardnh/be-warehouse-management/internal/domain/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReceivingItemRepositoryImpl struct{ db *gorm.DB }

func NewReceivingItemRepository(db *gorm.DB) domainrepositories.ReceivingItemRepository {
	return &ReceivingItemRepositoryImpl{db: db}
}

func (r *ReceivingItemRepositoryImpl) FindAllByReceiving(ctx context.Context, receivingID uuid.UUID) ([]entity.ReceivingItem, error) {
	var items []entity.ReceivingItem
	if err := Conn(ctx, r.db).Where("receiving_id = ?", receivingID).Preload("Product").Preload("InboundOrderItem").Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *ReceivingItemRepositoryImpl) FindByID(ctx context.Context, id uuid.UUID) (*entity.ReceivingItem, error) {
	var item entity.ReceivingItem
	if err := Conn(ctx, r.db).Preload("Product").Preload("InboundOrderItem").First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *ReceivingItemRepositoryImpl) Create(ctx context.Context, item entity.ReceivingItem) error {
	return Conn(ctx, r.db).Create(&item).Error
}

func (r *ReceivingItemRepositoryImpl) Update(ctx context.Context, item *entity.ReceivingItem) error {
	return Conn(ctx, r.db).Save(item).Error
}

func (r *ReceivingItemRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return Conn(ctx, r.db).Delete(&entity.ReceivingItem{}, id).Error
}
