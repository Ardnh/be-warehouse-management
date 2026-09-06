package repositories

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"github.com/google/uuid"
)

type ProductRepository interface {
	FindAll(context.Context, Filter) ([]entity.Product, int64, error)
	FindByID(context.Context, uuid.UUID) (*entity.Product, error)
	Create(context.Context, entity.Product) error
	Update(context.Context, *entity.Product) error
	Delete(context.Context, uuid.UUID) error
}
