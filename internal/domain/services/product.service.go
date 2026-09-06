package services

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/application/dto"
	"github.com/google/uuid"
)

type ProductService interface {
	FindAll(context.Context, dto.FilterDTO) ([]dto.Product, int64, error)
	FindByID(context.Context, uuid.UUID) (*dto.Product, error)
	Create(context.Context, dto.CreateProductRequest) error
	Update(context.Context, uuid.UUID, dto.UpdateProductRequest) error
	Delete(context.Context, uuid.UUID) error
}
