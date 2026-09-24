package services

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/application/dto"
	"github.com/google/uuid"
)

type LocationService interface {
	FindAll(ctx context.Context, filter dto.FilterDTO) ([]dto.LocationDTO, int64, error)
	FindByID(ctx context.Context, id uuid.UUID) (*dto.LocationDTO, error)
	Create(ctx context.Context, request dto.CreateLocationRequest) error
	Update(ctx context.Context, id uuid.UUID, request dto.UpdateLocationRequest) error
	Delete(ctx context.Context, id uuid.UUID) error
}
