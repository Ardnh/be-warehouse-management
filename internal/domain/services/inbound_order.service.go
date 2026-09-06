package services

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/application/dto"
	"github.com/google/uuid"
)

type InboundOrderService interface {
	FindAll(ctx context.Context, filter dto.FilterDTO) ([]dto.InboundOrderResponse, int64, error)
	FindByID(ctx context.Context, id uuid.UUID) (*dto.InboundOrderResponse, error)
	Create(ctx context.Context, request dto.CreateInboundOrderRequest) error
	Update(ctx context.Context, id uuid.UUID, request dto.UpdateInboundOrderRequest) error
	UpdateStatus(ctx context.Context, id uuid.UUID, request dto.UpdateInboundOrderStatusRequest) error
	Delete(ctx context.Context, id uuid.UUID) error
}
