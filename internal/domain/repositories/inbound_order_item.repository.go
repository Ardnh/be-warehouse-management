package repositories

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"github.com/google/uuid"
)

type InboundOrderItemRepository interface {
	FindAllByInboundOrder(ctx context.Context, inboundOrderID uuid.UUID) ([]entity.InboundOrderItem, error)
	FindByID(ctx context.Context, id uuid.UUID) (*entity.InboundOrderItem, error)
	Create(ctx context.Context, item entity.InboundOrderItem) error
	Update(ctx context.Context, item *entity.InboundOrderItem) error
	Delete(ctx context.Context, id uuid.UUID) error
}
