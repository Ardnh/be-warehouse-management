package repositories

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"github.com/google/uuid"
)

type ReceivingItemRepository interface {
	FindAllByReceiving(context.Context, uuid.UUID) ([]entity.ReceivingItem, error)
	FindByID(context.Context, uuid.UUID) (*entity.ReceivingItem, error)
	Create(context.Context, entity.ReceivingItem) error
	Update(context.Context, *entity.ReceivingItem) error
	Delete(context.Context, uuid.UUID) error
}
