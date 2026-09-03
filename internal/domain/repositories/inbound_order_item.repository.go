package repositories

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"github.com/google/uuid"
)

type InboundOrderItemRepository interface {
	FindAll(ctx context.Context) ([]entity.InboundOrderItem, error)
	FindByID(ctx context.Context, id uuid.UUID) (entity.InboundOrderItem, error)
	Create(ctx context.Context, order entity.InboundOrderItem) error
	Update(ctx context.Context, order entity.InboundOrderItem) error
	Delete(ctx context.Context, order entity.InboundOrderItem) error
}
