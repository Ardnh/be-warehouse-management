package services

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/application/dto"
	"github.com/google/uuid"
)

type HandlingUnitService interface {
	FindAll(ctx context.Context, filter dto.FilterDTO) ([]dto.HandlingUnitResponse, int64, error)
	FindByID(ctx context.Context, id uuid.UUID) (*dto.HandlingUnitResponse, error)
	Create(ctx context.Context, request dto.CreateHandlingUnitRequest) error
	Update(ctx context.Context, id uuid.UUID, request dto.UpdateHandlingUnitRequest) error
	Delete(ctx context.Context, id uuid.UUID) error
}
