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
	entry := s.log.WithFields(logrus.Fields{
		"component":        "inbound_order_item_service",
		"method":           "FindAll",
		"inbound_order_id": orderID.String(),
	})
	list, err := s.repository.FindAllByInboundOrder(ctx, orderID)
	if err != nil {
		entry.WithError(err).Error("failed to list inbound order items")
		return nil, err
	}
	entry.WithField("result_count", len(list)).Debug("inbound order items listed")
	return dto.ToInboundOrderItemResponses(list), nil
}
func (s *InboundOrderItemServiceImpl) FindByID(ctx context.Context, id uuid.UUID) (*dto.InboundOrderItemResponse, error) {
	entry := s.log.WithFields(logrus.Fields{
		"component":             "inbound_order_item_service",
		"method":                "FindByID",
		"inbound_order_item_id": id.String(),
	})
	item, err := s.repository.FindByID(ctx, id)
	if err != nil {
		entry.WithError(err).Error("failed to retrieve inbound order item")
		return nil, err
	}
	entry.Debug("inbound order item retrieved")
	result := dto.ToInboundOrderItemResponse(item)
	return &result, nil
}
func (s *InboundOrderItemServiceImpl) Create(ctx context.Context, orderID uuid.UUID, request dto.CreateInboundOrderItemRequest) error {
	entry := s.log.WithFields(logrus.Fields{
		"component":        "inbound_order_item_service",
		"method":           "Create",
		"inbound_order_id": orderID.String(),
		"product_id":       request.ProductID.String(),
	})
	if err := s.repository.Create(ctx, entity.InboundOrderItem{ID: uuid.New(), InboundOrderID: orderID, ProductID: request.ProductID, ExpectedQty: request.ExpectedQty, CreatedAt: time.Now()}); err != nil {
		entry.WithError(err).Error("failed to create inbound order item")
		return err
	}
	entry.Info("inbound order item created")
	return nil
}
func (s *InboundOrderItemServiceImpl) Update(ctx context.Context, id uuid.UUID, request dto.UpdateInboundOrderItemRequest) error {
	entry := s.log.WithFields(logrus.Fields{
		"component":             "inbound_order_item_service",
		"method":                "Update",
		"inbound_order_item_id": id.String(),
		"product_id":            request.ProductID.String(),
	})
	item, err := s.repository.FindByID(ctx, id)
	if err != nil {
		entry.WithError(err).Error("failed to find inbound order item for update")
		return err
	}
	item.ProductID = request.ProductID
	item.ExpectedQty = request.ExpectedQty
	if err := s.repository.Update(ctx, item); err != nil {
		entry.WithError(err).Error("failed to persist inbound order item update")
		return err
	}
	entry.Info("inbound order item updated")
	return nil
}
func (s *InboundOrderItemServiceImpl) Delete(ctx context.Context, id uuid.UUID) error {
	entry := s.log.WithFields(logrus.Fields{
		"component":             "inbound_order_item_service",
		"method":                "Delete",
		"inbound_order_item_id": id.String(),
	})
	if err := s.repository.Delete(ctx, id); err != nil {
		entry.WithError(err).Error("failed to delete inbound order item")
		return err
	}
	entry.Info("inbound order item deleted")
	return nil
}
