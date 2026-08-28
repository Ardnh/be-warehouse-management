package repository

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
)

type CustomerRepository interface {
	CreateCustomer(ctx context.Context, customer entity.Customer) error
}
