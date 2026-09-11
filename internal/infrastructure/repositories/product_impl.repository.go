package repositories

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	domainrepositories "github.com/Ardnh/be-warehouse-management/internal/domain/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct{ db *gorm.DB }

func NewProductRepository(db *gorm.DB) domainrepositories.ProductRepository {
	return &ProductRepositoryImpl{db: db}
}

func (r *ProductRepositoryImpl) FindAll(ctx context.Context, filter domainrepositories.Filter) ([]entity.Product, int64, error) {
	var items []entity.Product
	var total int64
	q := Conn(ctx, r.db).Model(&entity.Product{}).Preload("Customer").Preload("Uom")
	if filter.Search != "" {
		q = q.Where("sku ILIKE ? OR name ILIKE ?", "%"+filter.Search+"%", "%"+filter.Search+"%")
	}
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := applyMasterListFilter(q, filter).Find(&items).Error; err != nil {
		return nil, 0, err
	}
	return items, total, nil
}

func (r *ProductRepositoryImpl) FindByID(ctx context.Context, id uuid.UUID) (*entity.Product, error) {
	var item entity.Product
	if err := Conn(ctx, r.db).Preload("Customer").Preload("Uom").First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *ProductRepositoryImpl) Create(ctx context.Context, item entity.Product) error {
	return Conn(ctx, r.db).Create(&item).Error
}

func (r *ProductRepositoryImpl) Update(ctx context.Context, item *entity.Product) error {
	return Conn(ctx, r.db).Save(item).Error
}

func (r *ProductRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return Conn(ctx, r.db).Delete(&entity.Product{}, id).Error
}
