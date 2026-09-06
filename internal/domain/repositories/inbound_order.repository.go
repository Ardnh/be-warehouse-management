package repositories

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"github.com/google/uuid"
)

type InboundOrderRepository interface {
	FindAll(ctx context.Context, filter Filter) ([]entity.InboundOrder, int64, error)
	FindByID(ctx context.Context, id uuid.UUID) (*entity.InboundOrder, error)
	Create(ctx context.Context, order entity.InboundOrder) error
	Update(ctx context.Context, order *entity.InboundOrder) error
	Delete(ctx context.Context, id uuid.UUID) error
}
