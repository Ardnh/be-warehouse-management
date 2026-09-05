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

type UomServiceImpl struct {
	repository repositories.UomRepository
	log        *logrus.Logger
}

func NewUomService(repository repositories.UomRepository, log *logrus.Logger) domainservices.UomService {
	return &UomServiceImpl{repository: repository, log: log}
}

func (s *UomServiceImpl) FindAll(ctx context.Context, filter dto.FilterDTO) ([]dto.Uom, int64, error) {
	items, total, err := s.repository.FindAll(ctx, repositories.Filter{Search: filter.Search, Page: filter.Page, PageSize: filter.Size, SortBy: filter.SortBy, SortDir: filter.SortDir})
	if err != nil {
		return nil, 0, err
	}
	result := make([]dto.Uom, 0, len(items))
	for _, item := range items {
		result = append(result, dto.NewUomResponse(item))
	}
	return result, total, nil
}

func (s *UomServiceImpl) FindByID(ctx context.Context, id uuid.UUID) (*dto.Uom, error) {
	item, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	result := dto.NewUomResponse(*item)
	return &result, nil
}

func (s *UomServiceImpl) Create(ctx context.Context, request dto.CreateUomRequest) error {
	return s.repository.Create(ctx, entity.Uom{ID: uuid.New(), Code: request.Code, Name: request.Name, Type: request.Type, CreatedAt: time.Now()})
}

func (s *UomServiceImpl) Update(ctx context.Context, id uuid.UUID, request dto.UpdateUomRequest) error {
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
	return s.repository.Update(ctx, item)
}

func (s *UomServiceImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repository.Delete(ctx, id)
}
