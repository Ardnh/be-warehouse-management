package services

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/application/dto"
	"github.com/google/uuid"
)

type HandlingUnitItemService interface {
	FindAll(ctx context.Context, handlingUnitID uuid.UUID) ([]dto.HandlingUnitItemResponse, error)
	FindByID(ctx context.Context, id uuid.UUID) (*dto.HandlingUnitItemResponse, error)
	Create(ctx context.Context, handlingUnitID uuid.UUID, request dto.AddHandlingUnitItemRequest) error
	Update(ctx context.Context, id uuid.UUID, request dto.UpdateHandlingUnitItemRequest) error
	Delete(ctx context.Context, id uuid.UUID) error
}
