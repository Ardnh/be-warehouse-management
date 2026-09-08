package services

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/application/dto"
	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"github.com/Ardnh/be-warehouse-management/internal/domain/repositories"
	"github.com/Ardnh/be-warehouse-management/internal/domain/services"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type PermissionServiceImpl struct {
	permissionRepository repositories.PermissionRepository
	log                  *logrus.Logger
}

func NewPermissionService(permissionRepository repositories.PermissionRepository, log *logrus.Logger) services.PermissionService {
	return &PermissionServiceImpl{
		permissionRepository: permissionRepository,
		log:                  log,
	}
}

func (s *PermissionServiceImpl) FindAll(ctx context.Context, filter dto.FilterDTO) ([]dto.PermissionResponseDTO, int64, error) {

	filterDomain := repositories.Filter{
		Page:     filter.Page,
		PageSize: filter.Size,
		Search:   filter.Search,
		SortBy:   filter.SortBy,
		SortDir:  filter.SortDir,
	}

	permissions, total, err := s.permissionRepository.FindAll(ctx, filterDomain)
	if err != nil {
		return nil, 0, err
	}

	return dto.ToPermissionsDTO(permissions), total, nil
}

func (s *PermissionServiceImpl) FindByID(ctx context.Context, id uuid.UUID) (*dto.PermissionDTO, error) {
	permission, err := s.permissionRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return dto.ToPermissionDTO(permission), nil
}

func (s *PermissionServiceImpl) Create(ctx context.Context, req dto.CreatePermissionRequest) error {

	permission := entity.Permission{
		Resource:    req.Resource,
		Action:      req.Action,
		Description: req.Description,
	}

	err := s.permissionRepository.Create(ctx, permission)
	if err != nil {
		return err
	}

	return nil
}

func (s *PermissionServiceImpl) Update(ctx context.Context, id uuid.UUID, req dto.UpdatePermissionRequest) error {

	permission, err := s.permissionRepository.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if req.Resource != "" {
		permission.Resource = req.Resource
	}
	if req.Action != "" {
		permission.Action = req.Action
	}
	if req.Description != "" {
		permission.Description = req.Description
	}

	errUpdate := s.permissionRepository.Update(ctx, permission)
	if errUpdate != nil {
		return err
	}

	return nil
}

func (s *PermissionServiceImpl) Delete(ctx context.Context, id uuid.UUID) error {

	err := s.permissionRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
