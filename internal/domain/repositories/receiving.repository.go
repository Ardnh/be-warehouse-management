package repositories

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"github.com/google/uuid"
)

type ReceivingRepository interface {
	FindAll(context.Context, Filter) ([]entity.Receiving, int64, error)
	FindByID(context.Context, uuid.UUID) (*entity.Receiving, error)
	Create(context.Context, entity.Receiving) error
	Update(context.Context, *entity.Receiving) error
	Delete(context.Context, uuid.UUID) error
}
