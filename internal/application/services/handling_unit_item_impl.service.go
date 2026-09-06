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

type HandlingUnitItemServiceImpl struct {
	repository repositories.HandlingUnitItemRepository
	log        *logrus.Logger
}

func NewHandlingUnitItemService(repository repositories.HandlingUnitItemRepository, log *logrus.Logger) domainservices.HandlingUnitItemService {
	return &HandlingUnitItemServiceImpl{repository: repository, log: log}
}
func (s *HandlingUnitItemServiceImpl) FindAll(ctx context.Context, unitID uuid.UUID) ([]dto.HandlingUnitItemResponse, error) {
	list, err := s.repository.FindAllByHandlingUnit(ctx, unitID)
	if err != nil {
		return nil, err
	}
	result := make([]dto.HandlingUnitItemResponse, 0, len(list))
	for _, item := range list {
		result = append(result, dto.NewHandlingUnitItemResponse(item))
	}
	return result, nil
}
func (s *HandlingUnitItemServiceImpl) FindByID(ctx context.Context, id uuid.UUID) (*dto.HandlingUnitItemResponse, error) {
	item, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	result := dto.NewHandlingUnitItemResponse(*item)
	return &result, nil
}
func (s *HandlingUnitItemServiceImpl) Create(ctx context.Context, unitID uuid.UUID, request dto.AddHandlingUnitItemRequest) error {
	return s.repository.Create(ctx, entity.HandlingUnitItem{ID: uuid.New(), HandlingUnitID: unitID, ProductID: request.ProductID, Quantity: request.Quantity, CreatedAt: time.Now()})
}
func (s *HandlingUnitItemServiceImpl) Update(ctx context.Context, id uuid.UUID, request dto.UpdateHandlingUnitItemRequest) error {
	item, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if request.Quantity != nil {
		item.Quantity = *request.Quantity
	}
	return s.repository.Update(ctx, item)
}
func (s *HandlingUnitItemServiceImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repository.Delete(ctx, id)
}
