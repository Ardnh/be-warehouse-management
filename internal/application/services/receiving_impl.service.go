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

type ReceivingServiceImpl struct {
	repository repositories.ReceivingRepository
	items      repositories.ReceivingItemRepository
	log        *logrus.Logger
}

func NewReceivingService(r repositories.ReceivingRepository, i repositories.ReceivingItemRepository, l *logrus.Logger) domainservices.ReceivingService {
	return &ReceivingServiceImpl{repository: r, items: i, log: l}
}
func (s *ReceivingServiceImpl) FindAll(c context.Context, f dto.FilterDTO) ([]dto.ReceivingResponse, int64, error) {
	x, t, e := s.repository.FindAll(c, repositories.Filter{Search: f.Search, Page: f.Page, PageSize: f.Size, SortBy: f.SortBy, SortDir: f.SortDir})
	if e != nil {
		return nil, 0, e
	}
	return dto.ToReceivingResponses(x), t, nil
}
func (s *ReceivingServiceImpl) FindByID(c context.Context, id uuid.UUID) (*dto.ReceivingResponse, error) {
	x, e := s.repository.FindByID(c, id)
	if e != nil {
		return nil, e
	}
	return dto.ToReceivingResponse(x), nil
}
func (s *ReceivingServiceImpl) Create(c context.Context, r dto.CreateReceivingRequest) error {
	now := time.Now()
	n := "RCV-" + uuid.NewString()[:8]
	x := entity.Receiving{ID: uuid.New(), InboundOrderID: r.InboundOrderID, ReceivingNumber: n, Status: "draft", ReceivedAt: r.ReceivedAt, ReceivedBy: r.ReceivedBy, Notes: r.Notes, CreatedAt: now}
	for _, v := range r.Items {
		x.Items = append(x.Items, entity.ReceivingItem{ID: uuid.New(), ReceivingID: x.ID, InboundOrderItemID: v.InboundOrderItemID, ProductID: v.ProductID, ReceivedQty: v.ReceivedQty, CreatedAt: now})
	}
	return s.repository.Create(c, x)
}
func (s *ReceivingServiceImpl) Update(c context.Context, id uuid.UUID, r dto.UpdateReceivingRequest) error {
	x, e := s.repository.FindByID(c, id)
	if e != nil {
		return e
	}
	if r.Status != nil {
		x.Status = *r.Status
	}
	if r.ReceivedAt != nil {
		x.ReceivedAt = r.ReceivedAt
	}
	if r.ReceivedBy != nil {
		x.ReceivedBy = r.ReceivedBy
	}
	if r.Notes != nil {
		x.Notes = r.Notes
	}
	for _, v := range r.Items {
		if v.ID != nil {
			item, er := s.items.FindByID(c, *v.ID)
			if er != nil {
				return er
			}
			item.InboundOrderItemID = v.InboundOrderItemID
			item.ProductID = v.ProductID
			item.ReceivedQty = v.ReceivedQty
			if er = s.items.Update(c, item); er != nil {
				return er
			}
		} else if er := s.items.Create(c, entity.ReceivingItem{ID: uuid.New(), ReceivingID: id, InboundOrderItemID: v.InboundOrderItemID, ProductID: v.ProductID, ReceivedQty: v.ReceivedQty, CreatedAt: time.Now()}); er != nil {
			return er
		}
	}
	return s.repository.Update(c, x)
}
func (s *ReceivingServiceImpl) Delete(c context.Context, id uuid.UUID) error {
	return s.repository.Delete(c, id)
}
