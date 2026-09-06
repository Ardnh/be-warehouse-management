package services

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/application/dto"
	"github.com/google/uuid"
)

type InboundOrderItemService interface {
	FindAll(ctx context.Context, inboundOrderID uuid.UUID) ([]dto.InboundOrderItemResponse, error)
	FindByID(ctx context.Context, id uuid.UUID) (*dto.InboundOrderItemResponse, error)
	Create(ctx context.Context, inboundOrderID uuid.UUID, request dto.CreateInboundOrderItemRequest) error
	Update(ctx context.Context, id uuid.UUID, request dto.UpdateInboundOrderItemRequest) error
	Delete(ctx context.Context, id uuid.UUID) error
}
