package repositories

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"github.com/google/uuid"
)

type LocationRepository interface {
	FindAll(ctx context.Context, filter Filter) ([]entity.Location, int64, error)
	FindByID(ctx context.Context, id uuid.UUID) (*entity.Location, error)
	Create(ctx context.Context, location entity.Location) error
	Update(ctx context.Context, location *entity.Location) error
	Delete(ctx context.Context, id uuid.UUID) error
}
