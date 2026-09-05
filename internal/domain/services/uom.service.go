package services

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/application/dto"
	"github.com/google/uuid"
)

type UomService interface {
	FindAll(ctx context.Context, filter dto.FilterDTO) ([]dto.Uom, int64, error)
	FindByID(ctx context.Context, id uuid.UUID) (*dto.Uom, error)
	Create(ctx context.Context, request dto.CreateUomRequest) error
	Update(ctx context.Context, id uuid.UUID, request dto.UpdateUomRequest) error
	Delete(ctx context.Context, id uuid.UUID) error
}
