package repositories

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	dr "github.com/Ardnh/be-warehouse-management/internal/domain/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct{ db *gorm.DB }

func NewProductRepository(db *gorm.DB) dr.ProductRepository { return &ProductRepositoryImpl{db: db} }
func (r *ProductRepositoryImpl) FindAll(c context.Context, f dr.Filter) ([]entity.Product, int64, error) {
	var x []entity.Product
	var t int64
	q := r.db.WithContext(c).Model(&entity.Product{}).Preload("Customer").Preload("Uom")
	if f.Search != "" {
		q = q.Where("sku ILIKE ? OR name ILIKE ?", "%"+f.Search+"%", "%"+f.Search+"%")
	}
	if e := q.Count(&t).Error; e != nil {
		return nil, 0, e
	}
	if e := applyMasterListFilter(q, f).Find(&x).Error; e != nil {
		return nil, 0, e
	}
	return x, t, nil
}
func (r *ProductRepositoryImpl) FindByID(c context.Context, id uuid.UUID) (*entity.Product, error) {
	var x entity.Product
	e := r.db.WithContext(c).Preload("Customer").Preload("Uom").First(&x, id).Error
	if e != nil {
		return nil, e
	}
	return &x, nil
}
func (r *ProductRepositoryImpl) Create(c context.Context, x entity.Product) error {
	return r.db.WithContext(c).Create(&x).Error
}
func (r *ProductRepositoryImpl) Update(c context.Context, x *entity.Product) error {
	return r.db.WithContext(c).Save(x).Error
}
func (r *ProductRepositoryImpl) Delete(c context.Context, id uuid.UUID) error {
	return r.db.WithContext(c).Delete(&entity.Product{}, id).Error
}
func xOrNil(x entity.Product, e error) *entity.Product {
	if e != nil {
		return nil
	}
	return &x
}
