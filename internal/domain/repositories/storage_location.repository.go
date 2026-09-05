package repositories

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"github.com/google/uuid"
)

type StorageLocationRepository interface {
	FindAll(ctx context.Context, filter Filter) ([]entity.StorageLocation, int64, error)
	FindByID(ctx context.Context, id uuid.UUID) (*entity.StorageLocation, error)
	Create(ctx context.Context, location entity.StorageLocation) error
	Update(ctx context.Context, location *entity.StorageLocation) error
	Delete(ctx context.Context, id uuid.UUID) error
}
