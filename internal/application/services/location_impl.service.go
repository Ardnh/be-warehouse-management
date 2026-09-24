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

type LocationServiceImpl struct {
	repository repositories.LocationRepository
	log        *logrus.Logger
}

func NewLocationService(repository repositories.LocationRepository, log *logrus.Logger) domainservices.LocationService {
	return &LocationServiceImpl{repository: repository, log: log}
}

func (s *LocationServiceImpl) FindAll(ctx context.Context, filter dto.FilterDTO) ([]dto.LocationDTO, int64, error) {
	locations, total, err := s.repository.FindAll(ctx, repositories.Filter{
		Search:   filter.Search,
		Page:     filter.Page,
		PageSize: filter.Size,
		SortBy:   filter.SortBy,
		SortDir:  filter.SortDir,
	})
	if err != nil {
		return nil, 0, err
	}

	result := make([]dto.LocationDTO, 0, len(locations))
	for _, location := range locations {
		result = append(result, dto.ToLocationDTO(location))
	}
	return result, total, nil
}

func (s *LocationServiceImpl) FindByID(ctx context.Context, id uuid.UUID) (*dto.LocationDTO, error) {
	location, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	result := dto.ToLocationDTO(*location)
	return &result, nil
}

func (s *LocationServiceImpl) Create(ctx context.Context, request dto.CreateLocationRequest) error {
	isActive := true
	if request.IsActive != nil {
		isActive = *request.IsActive
	}

	return s.repository.Create(ctx, entity.Location{
		ID:         uuid.New(),
		Code:       request.Code,
		Name:       request.Name,
		Type:       request.Type,
		Address:    request.Address,
		City:       request.City,
		Province:   request.Province,
		PostalCode: request.PostalCode,
		IsActive:   isActive,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	})
}

func (s *LocationServiceImpl) Update(ctx context.Context, id uuid.UUID, request dto.UpdateLocationRequest) error {
	location, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if request.Name != nil {
		location.Name = *request.Name
	}
	if request.Type != nil {
		location.Type = *request.Type
	}
	if request.Address != nil {
		location.Address = request.Address
	}
	if request.City != nil {
		location.City = request.City
	}
	if request.Province != nil {
		location.Province = request.Province
	}
	if request.PostalCode != nil {
		location.PostalCode = request.PostalCode
	}
	if request.IsActive != nil {
		location.IsActive = *request.IsActive
	}
	location.UpdatedAt = time.Now()

	return s.repository.Update(ctx, location)
}

func (s *LocationServiceImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repository.Delete(ctx, id)
}
