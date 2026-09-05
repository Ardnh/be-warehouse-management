package repositories

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"github.com/google/uuid"
)

type ZoneRepository interface {
	FindAll(ctx context.Context, filter Filter) ([]entity.Zone, int64, error)
	FindByID(ctx context.Context, id uuid.UUID) (*entity.Zone, error)
	Create(ctx context.Context, zone entity.Zone) error
	Update(ctx context.Context, zone *entity.Zone) error
	Delete(ctx context.Context, id uuid.UUID) error
}
