package services

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/application/dto"
	"github.com/google/uuid"
)

type WarehouseService interface {
	FindAll(ctx context.Context, filter dto.FilterDTO) ([]dto.Warehouse, int64, error)
	FindByID(ctx context.Context, id uuid.UUID) (*dto.Warehouse, error)
	Create(ctx context.Context, request dto.CreateWarehouseRequest) error
	Update(ctx context.Context, id uuid.UUID, request dto.UpdateWarehouseRequest) error
	Delete(ctx context.Context, id uuid.UUID) error
}
