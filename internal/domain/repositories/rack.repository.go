package repositories

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"github.com/google/uuid"
)

type RackRepository interface {
	FindAll(ctx context.Context, filter Filter) ([]entity.Rack, int64, error)
	FindByID(ctx context.Context, id uuid.UUID) (*entity.Rack, error)
	Create(ctx context.Context, rack entity.Rack) error
	Update(ctx context.Context, rack *entity.Rack) error
	Delete(ctx context.Context, id uuid.UUID) error
}
