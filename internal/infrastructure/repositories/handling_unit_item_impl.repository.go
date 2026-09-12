package repositories

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	domainrepositories "github.com/Ardnh/be-warehouse-management/internal/domain/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type HandlingUnitItemRepositoryImpl struct{ db *gorm.DB }

func NewHandlingUnitItemRepository(db *gorm.DB) domainrepositories.HandlingUnitItemRepository {
	return &HandlingUnitItemRepositoryImpl{db: db}
}

func (r *HandlingUnitItemRepositoryImpl) FindAllByHandlingUnit(ctx context.Context, unitID uuid.UUID) ([]entity.HandlingUnitItem, error) {
	var list []entity.HandlingUnitItem
	if err := Conn(ctx, r.db).Where("handling_unit_id = ?", unitID).Preload("Product").Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *HandlingUnitItemRepositoryImpl) FindByID(ctx context.Context, id uuid.UUID) (*entity.HandlingUnitItem, error) {
	var item entity.HandlingUnitItem
	if err := Conn(ctx, r.db).Preload("Product").First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *HandlingUnitItemRepositoryImpl) Create(ctx context.Context, item entity.HandlingUnitItem) error {
	return Conn(ctx, r.db).Create(&item).Error
}

func (r *HandlingUnitItemRepositoryImpl) Update(ctx context.Context, item *entity.HandlingUnitItem) error {
	return Conn(ctx, r.db).Save(item).Error
}

func (r *HandlingUnitItemRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return Conn(ctx, r.db).Delete(&entity.HandlingUnitItem{}, id).Error
}
