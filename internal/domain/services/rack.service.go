package services

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/application/dto"
	"github.com/google/uuid"
)

type RackService interface {
	FindAll(ctx context.Context, filter dto.FilterDTO) ([]dto.RackResponse, int64, error)
	FindByID(ctx context.Context, id uuid.UUID) (*dto.RackResponse, error)
	Create(ctx context.Context, request dto.CreateRackRequest) error
	Update(ctx context.Context, id uuid.UUID, request dto.UpdateRackRequest) error
	Delete(ctx context.Context, id uuid.UUID) error
}
