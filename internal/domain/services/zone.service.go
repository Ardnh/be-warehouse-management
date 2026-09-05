package services

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/application/dto"
	"github.com/google/uuid"
)

type ZoneService interface {
	FindAll(ctx context.Context, filter dto.FilterDTO) ([]dto.ZoneResponse, int64, error)
	FindByID(ctx context.Context, id uuid.UUID) (*dto.ZoneResponse, error)
	Create(ctx context.Context, request dto.CreateZoneRequest) error
	Update(ctx context.Context, id uuid.UUID, request dto.UpdateZoneRequest) error
	Delete(ctx context.Context, id uuid.UUID) error
}
