package repositories

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"github.com/google/uuid"
)

type CustomerRepository interface {
	FindAll(ctx context.Context) ([]entity.Customer, error)
	FindByID(ctx context.Context, id uuid.UUID) (entity.Customer, error)
	Create(ctx context.Context, customer entity.Customer) error
	Update(ctx context.Context, customer entity.Customer) error
	Delete(ctx context.Context, customer entity.Customer) error
}
