package repositories

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	domainrepositories "github.com/Ardnh/be-warehouse-management/internal/domain/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InboundOrderItemRepositoryImpl struct{ db *gorm.DB }

func NewInboundOrderItemRepository(db *gorm.DB) domainrepositories.InboundOrderItemRepository {
	return &InboundOrderItemRepositoryImpl{db: db}
}

func (r *InboundOrderItemRepositoryImpl) FindAllByInboundOrder(ctx context.Context, orderID uuid.UUID) ([]entity.InboundOrderItem, error) {
	var list []entity.InboundOrderItem
	if err := Conn(ctx, r.db).Where("inbound_order_id = ?", orderID).Preload("Product").Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *InboundOrderItemRepositoryImpl) FindByID(ctx context.Context, id uuid.UUID) (*entity.InboundOrderItem, error) {
	var item entity.InboundOrderItem
	if err := Conn(ctx, r.db).Preload("Product").First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *InboundOrderItemRepositoryImpl) Create(ctx context.Context, item entity.InboundOrderItem) error {
	return Conn(ctx, r.db).Create(&item).Error
}

func (r *InboundOrderItemRepositoryImpl) Update(ctx context.Context, item *entity.InboundOrderItem) error {
	return Conn(ctx, r.db).Save(item).Error
}

func (r *InboundOrderItemRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return Conn(ctx, r.db).Delete(&entity.InboundOrderItem{}, id).Error
}
