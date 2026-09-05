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

type StorageLocationServiceImpl struct {
	repository repositories.StorageLocationRepository
	log        *logrus.Logger
}

func NewStorageLocationService(repository repositories.StorageLocationRepository, log *logrus.Logger) domainservices.StorageLocationService {
	return &StorageLocationServiceImpl{repository: repository, log: log}
}

func (s *StorageLocationServiceImpl) FindAll(ctx context.Context, filter dto.FilterDTO) ([]dto.StorageLocationResponse, int64, error) {
	items, total, err := s.repository.FindAll(ctx, repositories.Filter{Search: filter.Search, Page: filter.Page, PageSize: filter.Size, SortBy: filter.SortBy, SortDir: filter.SortDir})
	if err != nil {
		return nil, 0, err
	}
	result := make([]dto.StorageLocationResponse, 0, len(items))
	for _, item := range items {
		result = append(result, dto.NewStorageLocationResponse(item))
	}
	return result, total, nil
}

func (s *StorageLocationServiceImpl) FindByID(ctx context.Context, id uuid.UUID) (*dto.StorageLocationResponse, error) {
	item, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	result := dto.NewStorageLocationResponse(*item)
	return &result, nil
}

func (s *StorageLocationServiceImpl) Create(ctx context.Context, request dto.CreateStorageLocationRequest) error {
	return s.repository.Create(ctx, entity.StorageLocation{ID: uuid.New(), RackID: request.RackID, Code: request.Code, Bay: request.Bay, Level: request.Level, Capacity: request.Capacity, Status: request.Status, CreatedAt: time.Now()})
}

func (s *StorageLocationServiceImpl) Update(ctx context.Context, id uuid.UUID, request dto.UpdateStorageLocationRequest) error {
	item, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if request.Capacity != nil {
		item.Capacity = *request.Capacity
	}
	if request.Status != nil {
		item.Status = *request.Status
	}
	return s.repository.Update(ctx, item)
}

func (s *StorageLocationServiceImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repository.Delete(ctx, id)
}
