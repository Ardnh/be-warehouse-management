package repositories

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	dr "github.com/Ardnh/be-warehouse-management/internal/domain/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReceivingRepositoryImpl struct{ db *gorm.DB }

func NewReceivingRepository(db *gorm.DB) dr.ReceivingRepository {
	return &ReceivingRepositoryImpl{db: db}
}
func (r *ReceivingRepositoryImpl) FindAll(c context.Context, f dr.Filter) ([]entity.Receiving, int64, error) {
	var x []entity.Receiving
	var t int64
	q := r.db.WithContext(c).Model(&entity.Receiving{}).Preload("Items")
	if f.Search != "" {
		q = q.Where("receiving_number ILIKE ?", "%"+f.Search+"%")
	}
	if e := q.Count(&t).Error; e != nil {
		return nil, 0, e
	}
	if e := applyMasterListFilter(q, f).Find(&x).Error; e != nil {
		return nil, 0, e
	}
	return x, t, nil
}
func (r *ReceivingRepositoryImpl) FindByID(c context.Context, id uuid.UUID) (*entity.Receiving, error) {
	var x entity.Receiving
	e := r.db.WithContext(c).Preload("Items").First(&x, id).Error
	if e != nil {
		return nil, e
	}
	return &x, nil
}
func (r *ReceivingRepositoryImpl) Create(c context.Context, x entity.Receiving) error {
	return r.db.WithContext(c).Session(&gorm.Session{FullSaveAssociations: true}).Create(&x).Error
}
func (r *ReceivingRepositoryImpl) Update(c context.Context, x *entity.Receiving) error {
	return r.db.WithContext(c).Save(x).Error
}
func (r *ReceivingRepositoryImpl) Delete(c context.Context, id uuid.UUID) error {
	return r.db.WithContext(c).Delete(&entity.Receiving{}, id).Error
}
