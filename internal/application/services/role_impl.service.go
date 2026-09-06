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

type RoleServiceImpl struct {
	repository repositories.RoleRepository
	log        *logrus.Logger
}

func NewRoleService(r repositories.RoleRepository, l *logrus.Logger) domainservices.RoleService {
	return &RoleServiceImpl{repository: r, log: l}
}
func (s *RoleServiceImpl) FindAll(c context.Context, f dto.FilterDTO) ([]dto.RoleResponse, int64, error) {
	x, t, e := s.repository.FindAll(c, repositories.Filter{Search: f.Search, Page: f.Page, PageSize: f.Size, SortBy: f.SortBy, SortDir: f.SortDir})
	if e != nil {
		return nil, 0, e
	}
	o := make([]dto.RoleResponse, 0, len(x))
	for _, v := range x {
		o = append(o, dto.NewRoleResponse(v))
	}
	return o, t, nil
}
func (s *RoleServiceImpl) FindByID(c context.Context, id uuid.UUID) (*dto.RoleResponse, error) {
	x, e := s.repository.FindByID(c, id)
	if e != nil {
		return nil, e
	}
	v := dto.NewRoleResponse(*x)
	return &v, nil
}
func (s *RoleServiceImpl) Create(c context.Context, r dto.CreateRoleRequest) error {
	st := r.Status
	if st == "" {
		st = "ACTIVE"
	}
	return s.repository.Create(c, entity.Role{ID: uuid.New(), Code: r.Code, Name: r.Name, Description: r.Description, Status: st, CreatedAt: time.Now()})
}
func (s *RoleServiceImpl) Update(c context.Context, id uuid.UUID, r dto.UpdateRoleRequest) error {
	x, e := s.repository.FindByID(c, id)
	if e != nil {
		return e
	}
	if r.Name != nil {
		x.Name = *r.Name
	}
	if r.Description != nil {
		x.Description = *r.Description
	}
	if r.Status != nil {
		x.Status = *r.Status
	}
	return s.repository.Update(c, x)
}
func (s *RoleServiceImpl) Delete(c context.Context, id uuid.UUID) error {
	return s.repository.Delete(c, id)
}
