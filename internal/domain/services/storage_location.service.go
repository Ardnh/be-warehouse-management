package services

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/application/dto"
	"github.com/google/uuid"
)

type StorageLocationService interface {
	FindAll(ctx context.Context, filter dto.FilterDTO) ([]dto.StorageLocationResponse, int64, error)
	FindByID(ctx context.Context, id uuid.UUID) (*dto.StorageLocationResponse, error)
	Create(ctx context.Context, request dto.CreateStorageLocationRequest) error
	Update(ctx context.Context, id uuid.UUID, request dto.UpdateStorageLocationRequest) error
	Delete(ctx context.Context, id uuid.UUID) error
}
