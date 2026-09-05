package repositories

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"github.com/google/uuid"
)

type UomRepository interface {
	FindAll(ctx context.Context, filter Filter) ([]entity.Uom, int64, error)
	FindByID(ctx context.Context, id uuid.UUID) (*entity.Uom, error)
	Create(ctx context.Context, uom entity.Uom) error
	Update(ctx context.Context, uom *entity.Uom) error
	Delete(ctx context.Context, id uuid.UUID) error
}
