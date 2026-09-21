package services

import (
	"context"
	"errors"
	"time"

	"github.com/Ardnh/be-warehouse-management/internal/application/dto"
	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"github.com/Ardnh/be-warehouse-management/internal/domain/repositories"
	domainservices "github.com/Ardnh/be-warehouse-management/internal/domain/services"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type InboundOrderServiceImpl struct {
	inboundOrder                repositories.InboundOrderRepository
	inboundOrderItems           repositories.InboundOrderItemRepository
	customerWarehouseRepository repositories.CustomerWarehouseRepository
	productWarehouseRepository  repositories.ProductWarehouseRepository
	productRepository           repositories.ProductRepository
	customers                   repositories.CustomerRepository
	tx                          repositories.TxManager
	log                         *logrus.Logger
}

func NewInboundOrderService(repository repositories.InboundOrderRepository, items repositories.InboundOrderItemRepository, customerWarehouseRepository repositories.CustomerWarehouseRepository, productWarehouseRepository repositories.ProductWarehouseRepository, productRepository repositories.ProductRepository, customers repositories.CustomerRepository, tx repositories.TxManager, log *logrus.Logger) domainservices.InboundOrderService {
	return &InboundOrderServiceImpl{
		inboundOrder:                repository,
		inboundOrderItems:           items,
		customerWarehouseRepository: customerWarehouseRepository,
		productWarehouseRepository:  productWarehouseRepository,
		productRepository:           productRepository,
		customers:                   customers,
		tx:                          tx,
		log:                         log,
	}
}

func (s *InboundOrderServiceImpl) FindAll(ctx context.Context, filter dto.FilterDTO) ([]dto.InboundOrderResponse, int64, error) {
	list, total, err := s.inboundOrder.FindAll(ctx, repositories.Filter{Search: filter.Search, Page: filter.Page, PageSize: filter.Size, SortBy: filter.SortBy, SortDir: filter.SortDir})
	if err != nil {
		return nil, 0, err
	}
	return dto.ToInboundOrderResponses(list), total, nil
}

func (s *InboundOrderServiceImpl) FindByID(ctx context.Context, id uuid.UUID) (*dto.InboundOrderResponse, error) {
	item, err := s.inboundOrder.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return dto.ToInboundOrderResponse(item), nil
}

func (s *InboundOrderServiceImpl) Create(ctx context.Context, warehouseID uuid.UUID, request dto.CreateInboundOrderRequest) error {

	return s.tx.Do(ctx, func(txCtx context.Context) error {

		// 1. Validate Customer
		customer, err := s.customers.FindByID(txCtx, request.CustomerID)
		if err != nil {
			return err
		}

		if customer == nil {
			return fiber.NewError(
				fiber.StatusNotFound,
				"customer not found",
			)
		}

		// 2. Auto-register Customer -> Warehouse
		customerWarehouse, err := s.customerWarehouseRepository.FindByCustomerAndWarehouse(
			txCtx,
			request.CustomerID,
			warehouseID,
		)

		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		if customerWarehouse == nil {
			customerWarehouse = &entity.CustomerWarehouse{
				ID:          uuid.New(),
				CustomerID:  request.CustomerID,
				WarehouseID: warehouseID,
				Status:      "active",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			}

			if err := s.customerWarehouseRepository.Create(txCtx, *customerWarehouse); err != nil {
				return err
			}
		}

		// 3. Create Inbound Order
		order := entity.InboundOrder{
			ID:                uuid.New(),
			WarehouseID:       warehouseID,
			CustomerID:        request.CustomerID,
			OrderNumber:       "IN-" + uuid.NewString()[:8],
			Status:            entity.InboundStatusDraft,
			ExpectedArrivalAt: request.ExpectedArrivalAt,
			Notes:             request.Notes,
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
		}

		// 4. Validate + auto-register Products
		for _, item := range request.Items {

			product, err := s.productRepository.FindByID(
				txCtx,
				item.ProductID,
			)
			if err != nil {
				return err
			}

			if product == nil {
				return errors.New("product not found")
			}

			productWarehouse, err := s.productWarehouseRepository.FindByProductAndWarehouse(
				txCtx,
				item.ProductID,
				warehouseID,
			)

			if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				return err
			}

			if productWarehouse == nil {
				productWarehouse = &entity.ProductWarehouse{
					ID:          uuid.New(),
					ProductID:   item.ProductID,
					WarehouseID: warehouseID,
					Status:      "active",
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				}

				if err := s.productWarehouseRepository.Create(
					txCtx,
					*productWarehouse,
				); err != nil {
					return err
				}
			}

			// 5. Add Inbound Order Item
			order.Items = append(
				order.Items,
				entity.InboundOrderItem{
					ID:             uuid.New(),
					InboundOrderID: order.ID,
					ProductID:      item.ProductID,
					ExpectedQty:    item.ExpectedQty,
					CreatedAt:      time.Now(),
				},
			)
		}

		// 6. Create Inbound Order + Items
		return s.inboundOrder.Create(txCtx, order)
	})
}

func (s *InboundOrderServiceImpl) Update(ctx context.Context, id uuid.UUID, request dto.UpdateInboundOrderRequest) error {
	// order, err := s.inboundOrder.FindByID(ctx, id)
	// if err != nil {
	// 	return err
	// }
	// if !order.Status.IsEditable() {
	// 	return fiberErr("inbound order cannot be edited in its current status")
	// }
	// if request.CustomerID != nil {
	// 	order.CustomerID = *request.CustomerID
	// }
	// if request.ExpectedArrivalAt != nil {
	// 	order.ExpectedArrivalAt = request.ExpectedArrivalAt
	// }
	// if request.Notes != nil {
	// 	order.Notes = request.Notes
	// }
	// for _, item := range request.Items {
	// 	if item.ID != nil {
	// 		existing, findErr := s.items.FindByID(ctx, *item.ID)
	// 		if findErr != nil {
	// 			return findErr
	// 		}
	// 		existing.ProductID = item.ProductID
	// 		existing.ExpectedQty = item.ExpectedQty
	// 		if err = s.items.Update(ctx, existing); err != nil {
	// 			return err
	// 		}
	// 	} else if err = s.items.Create(ctx, entity.InboundOrderItem{ID: uuid.New(), InboundOrderID: id, ProductID: item.ProductID, ExpectedQty: item.ExpectedQty, CreatedAt: time.Now()}); err != nil {
	// 		return err
	// 	}
	// }
	return nil
}

func (s *InboundOrderServiceImpl) UpdateStatus(ctx context.Context, id uuid.UUID, request dto.UpdateInboundOrderStatusRequest) error {
	order, err := s.inboundOrder.FindByID(ctx, id)
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
	return s.inboundOrder.Update(ctx, order)
}

func (s *InboundOrderServiceImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return s.inboundOrder.Delete(ctx, id)
}

// fiberErr is kept local to avoid coupling the application service to HTTP details.
func fiberErr(message string) error { return &serviceError{message: message} }

type serviceError struct{ message string }

func (e *serviceError) Error() string { return e.message }
