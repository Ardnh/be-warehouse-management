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

type WarehouseServiceImpl struct {
	repository repositories.WarehouseRepository
	log        *logrus.Logger
}

func NewWarehouseService(repository repositories.WarehouseRepository, log *logrus.Logger) domainservices.WarehouseService {
	return &WarehouseServiceImpl{repository: repository, log: log}
}

func (s *WarehouseServiceImpl) FindAll(ctx context.Context, filter dto.FilterDTO) ([]dto.Warehouse, int64, error) {
	items, total, err := s.repository.FindAll(ctx, repositories.Filter{Search: filter.Search, Page: filter.Page, PageSize: filter.Size, SortBy: filter.SortBy, SortDir: filter.SortDir})
	if err != nil {
		return nil, 0, err
	}
	result := make([]dto.Warehouse, 0, len(items))
	for _, item := range items {
		result = append(result, dto.NewWarehouseResponse(item))
	}
	return result, total, nil
}

func (s *WarehouseServiceImpl) FindByID(ctx context.Context, id uuid.UUID) (*dto.Warehouse, error) {
	item, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	result := dto.NewWarehouseResponse(*item)
	return &result, nil
}

func (s *WarehouseServiceImpl) Create(ctx context.Context, request dto.CreateWarehouseRequest) error {
	return s.repository.Create(ctx, entity.Warehouse{ID: uuid.New(), Code: request.Code, Name: request.Name, Address: request.Address, Status: request.Status, CreatedAt: time.Now()})
}

func (s *WarehouseServiceImpl) Update(ctx context.Context, id uuid.UUID, request dto.UpdateWarehouseRequest) error {
	item, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if request.Name != nil {
		item.Name = *request.Name
	}
	if request.Address != nil {
		item.Address = *request.Address
	}
	if request.Status != nil {
		item.Status = *request.Status
	}
	return s.repository.Update(ctx, item)
}

func (s *WarehouseServiceImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repository.Delete(ctx, id)
}
