package services

import (
	"context"
	"time"

	"github.com/Ardnh/be-warehouse-management/internal/application/dto"
	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"github.com/Ardnh/be-warehouse-management/internal/domain/repositories"
	domainservices "github.com/Ardnh/be-warehouse-management/internal/domain/services"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type InboundOrderItemServiceImpl struct {
	repository repositories.InboundOrderItemRepository
	log        *logrus.Logger
}

func NewInboundOrderItemService(repository repositories.InboundOrderItemRepository, log *logrus.Logger) domainservices.InboundOrderItemService {
	return &InboundOrderItemServiceImpl{repository: repository, log: log}
}
func (s *InboundOrderItemServiceImpl) FindAll(ctx context.Context, orderID uuid.UUID) ([]dto.InboundOrderItemResponse, error) {
	list, err := s.repository.FindAllByInboundOrder(ctx, orderID)
	if err != nil {
		return nil, err
	}
	return dto.ToInboundOrderItemResponses(list), nil
}
func (s *InboundOrderItemServiceImpl) FindByID(ctx context.Context, id uuid.UUID) (*dto.InboundOrderItemResponse, error) {
	item, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	result := dto.ToInboundOrderItemResponse(item)
	return &result, nil
}
func (s *InboundOrderItemServiceImpl) Create(ctx context.Context, orderID uuid.UUID, request dto.CreateInboundOrderItemRequest) error {
	return s.repository.Create(ctx, entity.InboundOrderItem{ID: uuid.New(), InboundOrderID: orderID, ProductID: request.ProductID, ExpectedQty: request.ExpectedQty, CreatedAt: time.Now()})
}
func (s *InboundOrderItemServiceImpl) Update(ctx context.Context, id uuid.UUID, request dto.UpdateInboundOrderItemRequest) error {
	item, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return err
	}
	item.ProductID = request.ProductID
	item.ExpectedQty = request.ExpectedQty
	return s.repository.Update(ctx, item)
}
func (s *InboundOrderItemServiceImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repository.Delete(ctx, id)
}
