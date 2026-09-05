package services

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/application/dto"
	"github.com/google/uuid"
)

type CustomerService interface {
	FindAll(ctx context.Context, filter dto.FilterDTO) ([]dto.Customer, int64, error)
	FindByID(ctx context.Context, id uuid.UUID) (*dto.Customer, error)
	Create(ctx context.Context, customer dto.CreateCustomerRequest) error
	Update(ctx context.Context, id uuid.UUID, customer dto.UpdateCustomerRequest) error
	Delete(ctx context.Context, id uuid.UUID) error
}
