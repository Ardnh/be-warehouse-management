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

type HandlingUnitServiceImpl struct {
	repository repositories.HandlingUnitRepository
	log        *logrus.Logger
}

func NewHandlingUnitService(repository repositories.HandlingUnitRepository, log *logrus.Logger) domainservices.HandlingUnitService {
	return &HandlingUnitServiceImpl{repository: repository, log: log}
}
func (s *HandlingUnitServiceImpl) FindAll(ctx context.Context, filter dto.FilterDTO) ([]dto.HandlingUnitResponse, int64, error) {
	list, total, err := s.repository.FindAll(ctx, repositories.Filter{Search: filter.Search, Page: filter.Page, PageSize: filter.Size, SortBy: filter.SortBy, SortDir: filter.SortDir})
	if err != nil {
		return nil, 0, err
	}
	return dto.NewHandlingUnitResponses(list), total, nil
}
func (s *HandlingUnitServiceImpl) FindByID(ctx context.Context, id uuid.UUID) (*dto.HandlingUnitResponse, error) {
	item, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	result := dto.NewHandlingUnitResponse(*item)
	return &result, nil
}
func (s *HandlingUnitServiceImpl) Create(ctx context.Context, request dto.CreateHandlingUnitRequest) error {
	status := request.Status
	if status == "" {
		status = "EMPTY"
	}
	return s.repository.Create(ctx, entity.HandlingUnit{ID: uuid.New(), Code: request.Code, WarehouseID: request.WarehouseID, Status: status, CreatedAt: time.Now()})
}
func (s *HandlingUnitServiceImpl) Update(ctx context.Context, id uuid.UUID, request dto.UpdateHandlingUnitRequest) error {
	item, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if request.Status != nil {
		item.Status = *request.Status
	}
	return s.repository.Update(ctx, item)
}
func (s *HandlingUnitServiceImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repository.Delete(ctx, id)
}
