package services

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/application/dto"
	"github.com/google/uuid"
)

type ReceivingService interface {
	FindAll(context.Context, dto.FilterDTO) ([]dto.ReceivingResponse, int64, error)
	FindByID(context.Context, uuid.UUID) (*dto.ReceivingResponse, error)
	Create(context.Context, dto.CreateReceivingRequest) error
	Update(context.Context, uuid.UUID, dto.UpdateReceivingRequest) error
	Delete(context.Context, uuid.UUID) error
}
