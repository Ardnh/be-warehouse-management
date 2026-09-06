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

type ReceivingItemServiceImpl struct {
	repository repositories.ReceivingItemRepository
	log        *logrus.Logger
}

func NewReceivingItemService(r repositories.ReceivingItemRepository, l *logrus.Logger) domainservices.ReceivingItemService {
	return &ReceivingItemServiceImpl{repository: r, log: l}
}
func (s *ReceivingItemServiceImpl) FindAll(c context.Context, id uuid.UUID) ([]dto.ReceivingItemResponse, error) {
	x, e := s.repository.FindAllByReceiving(c, id)
	if e != nil {
		return nil, e
	}
	return dto.ToReceivingItemResponses(x), nil
}
func (s *ReceivingItemServiceImpl) FindByID(c context.Context, id uuid.UUID) (*dto.ReceivingItemResponse, error) {
	x, e := s.repository.FindByID(c, id)
	if e != nil {
		return nil, e
	}
	v := dto.ToReceivingItemResponse(x)
	return &v, nil
}
func (s *ReceivingItemServiceImpl) Create(c context.Context, id uuid.UUID, r dto.CreateReceivingItemRequest) error {
	return s.repository.Create(c, entity.ReceivingItem{ID: uuid.New(), ReceivingID: id, InboundOrderItemID: r.InboundOrderItemID, ProductID: r.ProductID, ReceivedQty: r.ReceivedQty, CreatedAt: time.Now()})
}
func (s *ReceivingItemServiceImpl) Update(c context.Context, id uuid.UUID, r dto.UpdateReceivingItemRequest) error {
	x, e := s.repository.FindByID(c, id)
	if e != nil {
		return e
	}
	x.InboundOrderItemID = r.InboundOrderItemID
	x.ProductID = r.ProductID
	x.ReceivedQty = r.ReceivedQty
	return s.repository.Update(c, x)
}
func (s *ReceivingItemServiceImpl) Delete(c context.Context, id uuid.UUID) error {
	return s.repository.Delete(c, id)
}
