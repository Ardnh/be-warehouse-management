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

type RackServiceImpl struct {
	repository repositories.RackRepository
	log        *logrus.Logger
}

func NewRackService(repository repositories.RackRepository, log *logrus.Logger) domainservices.RackService {
	return &RackServiceImpl{repository: repository, log: log}
}

func (s *RackServiceImpl) FindAll(ctx context.Context, filter dto.FilterDTO) ([]dto.RackResponse, int64, error) {
	items, total, err := s.repository.FindAll(ctx, repositories.Filter{Search: filter.Search, Page: filter.Page, PageSize: filter.Size, SortBy: filter.SortBy, SortDir: filter.SortDir})
	if err != nil {
		return nil, 0, err
	}
	result := make([]dto.RackResponse, 0, len(items))
	for _, item := range items {
		result = append(result, dto.NewRackResponse(item))
	}
	return result, total, nil
}

func (s *RackServiceImpl) FindByID(ctx context.Context, id uuid.UUID) (*dto.RackResponse, error) {
	item, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	result := dto.NewRackResponse(*item)
	return &result, nil
}

func (s *RackServiceImpl) Create(ctx context.Context, request dto.CreateRackRequest) error {
	return s.repository.Create(ctx, entity.Rack{ID: uuid.New(), ZoneID: request.ZoneID, Code: request.Code, BayCount: request.BayCount, LevelCount: request.LevelCount, PalletCapacity: request.PalletCapacity, Status: request.Status, CreatedAt: time.Now()})
}

func (s *RackServiceImpl) Update(ctx context.Context, id uuid.UUID, request dto.UpdateRackRequest) error {
	item, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if request.BayCount != nil {
		item.BayCount = *request.BayCount
	}
	if request.LevelCount != nil {
		item.LevelCount = *request.LevelCount
	}
	if request.PalletCapacity != nil {
		item.PalletCapacity = *request.PalletCapacity
	}
	if request.Status != nil {
		item.Status = *request.Status
	}
	return s.repository.Update(ctx, item)
}

func (s *RackServiceImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repository.Delete(ctx, id)
}
