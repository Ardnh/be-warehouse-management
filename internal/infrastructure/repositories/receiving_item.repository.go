package repositories

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	dr "github.com/Ardnh/be-warehouse-management/internal/domain/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReceivingItemRepositoryImpl struct{ db *gorm.DB }

func NewReceivingItemRepository(db *gorm.DB) dr.ReceivingItemRepository {
	return &ReceivingItemRepositoryImpl{db: db}
}
func (r *ReceivingItemRepositoryImpl) FindAllByReceiving(c context.Context, id uuid.UUID) ([]entity.ReceivingItem, error) {
	var x []entity.ReceivingItem
	e := r.db.WithContext(c).Where("receiving_id = ?", id).Preload("Product").Preload("InboundOrderItem").Find(&x).Error
	return x, e
}
func (r *ReceivingItemRepositoryImpl) FindByID(c context.Context, id uuid.UUID) (*entity.ReceivingItem, error) {
	var x entity.ReceivingItem
	e := r.db.WithContext(c).Preload("Product").Preload("InboundOrderItem").First(&x, id).Error
	if e != nil {
		return nil, e
	}
	return &x, nil
}
func (r *ReceivingItemRepositoryImpl) Create(c context.Context, x entity.ReceivingItem) error {
	return r.db.WithContext(c).Create(&x).Error
}
func (r *ReceivingItemRepositoryImpl) Update(c context.Context, x *entity.ReceivingItem) error {
	return r.db.WithContext(c).Save(x).Error
}
func (r *ReceivingItemRepositoryImpl) Delete(c context.Context, id uuid.UUID) error {
	return r.db.WithContext(c).Delete(&entity.ReceivingItem{}, id).Error
}
