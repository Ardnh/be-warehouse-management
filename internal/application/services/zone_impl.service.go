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

type ZoneServiceImpl struct {
	repository repositories.ZoneRepository
	log        *logrus.Logger
}

func NewZoneService(repository repositories.ZoneRepository, log *logrus.Logger) domainservices.ZoneService {
	return &ZoneServiceImpl{repository: repository, log: log}
}

func (s *ZoneServiceImpl) FindAll(ctx context.Context, filter dto.FilterDTO) ([]dto.ZoneResponse, int64, error) {
	items, total, err := s.repository.FindAll(ctx, repositories.Filter{Search: filter.Search, Page: filter.Page, PageSize: filter.Size, SortBy: filter.SortBy, SortDir: filter.SortDir})
	if err != nil {
		return nil, 0, err
	}
	result := make([]dto.ZoneResponse, 0, len(items))
	for _, item := range items {
		result = append(result, dto.NewZoneResponse(item))
	}
	return result, total, nil
}

func (s *ZoneServiceImpl) FindByID(ctx context.Context, id uuid.UUID) (*dto.ZoneResponse, error) {
	item, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	result := dto.NewZoneResponse(*item)
	return &result, nil
}

func (s *ZoneServiceImpl) Create(ctx context.Context, request dto.CreateZoneRequest) error {
	return s.repository.Create(ctx, entity.Zone{ID: uuid.New(), WarehouseID: request.WarehouseID, Code: request.Code, Name: request.Name, Type: request.Type, Status: request.Status, CreatedAt: time.Now()})
}

func (s *ZoneServiceImpl) Update(ctx context.Context, id uuid.UUID, request dto.UpdateZoneRequest) error {
	item, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if request.Name != nil {
		item.Name = *request.Name
	}
	if request.Type != nil {
		item.Type = *request.Type
	}
	if request.Status != nil {
		item.Status = *request.Status
	}
	return s.repository.Update(ctx, item)
}

func (s *ZoneServiceImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repository.Delete(ctx, id)
}
