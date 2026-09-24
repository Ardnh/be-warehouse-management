package repositories

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	domainrepositories "github.com/Ardnh/be-warehouse-management/internal/domain/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LocationRepositoryImpl struct {
	db *gorm.DB
}

func NewLocationRepository(db *gorm.DB) domainrepositories.LocationRepository {
	return &LocationRepositoryImpl{db: db}
}

func (r *LocationRepositoryImpl) FindAll(ctx context.Context, filter domainrepositories.Filter) ([]entity.Location, int64, error) {
	var locations []entity.Location
	var total int64

	query := Conn(ctx, r.db).Model(&entity.Location{})
	if filter.Search != "" {
		pattern := "%" + filter.Search + "%"
		query = query.Where("code ILIKE ? OR name ILIKE ? OR city ILIKE ? OR province ILIKE ?", pattern, pattern, pattern, pattern)
	}
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := applyMasterListFilter(query, filter).Find(&locations).Error; err != nil {
		return nil, 0, err
	}
	return locations, total, nil
}

func (r *LocationRepositoryImpl) FindByID(ctx context.Context, id uuid.UUID) (*entity.Location, error) {
	var location entity.Location
	if err := Conn(ctx, r.db).First(&location, id).Error; err != nil {
		return nil, err
	}
	return &location, nil
}

func (r *LocationRepositoryImpl) Create(ctx context.Context, location entity.Location) error {
	return Conn(ctx, r.db).Create(&location).Error
}

func (r *LocationRepositoryImpl) Update(ctx context.Context, location *entity.Location) error {
	return Conn(ctx, r.db).Save(location).Error
}

func (r *LocationRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return Conn(ctx, r.db).Delete(&entity.Location{}, id).Error
}
