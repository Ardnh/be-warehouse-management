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

type InboundOrderServiceImpl struct {
	repository repositories.InboundOrderRepository
	items      repositories.InboundOrderItemRepository
	log        *logrus.Logger
}

func NewInboundOrderService(repository repositories.InboundOrderRepository, items repositories.InboundOrderItemRepository, log *logrus.Logger) domainservices.InboundOrderService {
	return &InboundOrderServiceImpl{repository: repository, items: items, log: log}
}
func (s *InboundOrderServiceImpl) FindAll(ctx context.Context, filter dto.FilterDTO) ([]dto.InboundOrderResponse, int64, error) {
	list, total, err := s.repository.FindAll(ctx, repositories.Filter{Search: filter.Search, Page: filter.Page, PageSize: filter.Size, SortBy: filter.SortBy, SortDir: filter.SortDir})
	if err != nil {
		return nil, 0, err
	}
	return dto.ToInboundOrderResponses(list), total, nil
}
func (s *InboundOrderServiceImpl) FindByID(ctx context.Context, id uuid.UUID) (*dto.InboundOrderResponse, error) {
	item, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return dto.ToInboundOrderResponse(item), nil
}
func (s *InboundOrderServiceImpl) Create(ctx context.Context, request dto.CreateInboundOrderRequest) error {
	order := entity.InboundOrder{ID: uuid.New(), CustomerID: request.CustomerID, OrderNumber: "IN-" + uuid.NewString()[:8], Status: entity.InboundStatusDraft, ExpectedArrivalAt: request.ExpectedArrivalAt, Notes: request.Notes, CreatedAt: time.Now()}
	for _, item := range request.Items {
		order.Items = append(order.Items, entity.InboundOrderItem{ID: uuid.New(), InboundOrderID: order.ID, ProductID: item.ProductID, ExpectedQty: item.ExpectedQty, CreatedAt: time.Now()})
	}
	return s.repository.Create(ctx, order)
}
func (s *InboundOrderServiceImpl) Update(ctx context.Context, id uuid.UUID, request dto.UpdateInboundOrderRequest) error {
	order, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if !order.Status.IsEditable() {
		return fiberErr("inbound order cannot be edited in its current status")
	}
	if request.CustomerID != nil {
		order.CustomerID = *request.CustomerID
	}
	if request.ExpectedArrivalAt != nil {
		order.ExpectedArrivalAt = request.ExpectedArrivalAt
	}
	if request.Notes != nil {
		order.Notes = request.Notes
	}
	for _, item := range request.Items {
		if item.ID != nil {
			existing, findErr := s.items.FindByID(ctx, *item.ID)
			if findErr != nil {
				return findErr
			}
			existing.ProductID = item.ProductID
			existing.ExpectedQty = item.ExpectedQty
			if err = s.items.Update(ctx, existing); err != nil {
				return err
			}
		} else if err = s.items.Create(ctx, entity.InboundOrderItem{ID: uuid.New(), InboundOrderID: id, ProductID: item.ProductID, ExpectedQty: item.ExpectedQty, CreatedAt: time.Now()}); err != nil {
			return err
		}
	}
	return s.repository.Update(ctx, order)
}
func (s *InboundOrderServiceImpl) UpdateStatus(ctx context.Context, id uuid.UUID, request dto.UpdateInboundOrderStatusRequest) error {
	order, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return err
	}
	status := entity.InboundOrderStatus(request.Status)
	if !status.IsValid() {
		return fiberErr("invalid inbound order status")
	}
	order.Status = status
	if request.Notes != nil {
		order.Notes = request.Notes
	}
	return s.repository.Update(ctx, order)
}
func (s *InboundOrderServiceImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repository.Delete(ctx, id)
}

// fiberErr is kept local to avoid coupling the application service to HTTP details.
func fiberErr(message string) error { return &serviceError{message: message} }

type serviceError struct{ message string }

func (e *serviceError) Error() string { return e.message }
