package services

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/application/dto"
	"github.com/google/uuid"
)

type ReceivingItemService interface {
	FindAll(context.Context, uuid.UUID) ([]dto.ReceivingItemResponse, error)
	FindByID(context.Context, uuid.UUID) (*dto.ReceivingItemResponse, error)
	Create(context.Context, uuid.UUID, dto.CreateReceivingItemRequest) error
	Update(context.Context, uuid.UUID, dto.UpdateReceivingItemRequest) error
	Delete(context.Context, uuid.UUID) error
}
