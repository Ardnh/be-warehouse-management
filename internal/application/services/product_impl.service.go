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

type ProductServiceImpl struct {
	repository repositories.ProductRepository
	log        *logrus.Logger
}

func NewProductService(r repositories.ProductRepository, l *logrus.Logger) domainservices.ProductService {
	return &ProductServiceImpl{repository: r, log: l}
}
func (s *ProductServiceImpl) FindAll(c context.Context, f dto.FilterDTO) ([]dto.Product, int64, error) {
	x, t, e := s.repository.FindAll(c, repositories.Filter{Search: f.Search, Page: f.Page, PageSize: f.Size, SortBy: f.SortBy, SortDir: f.SortDir})
	if e != nil {
		return nil, 0, e
	}
	out := make([]dto.Product, 0, len(x))
	for _, v := range x {
		out = append(out, dto.NewProductResponse(v))
	}
	return out, t, nil
}
func (s *ProductServiceImpl) FindByID(c context.Context, id uuid.UUID) (*dto.Product, error) {
	x, e := s.repository.FindByID(c, id)
	if e != nil {
		return nil, e
	}
	v := dto.NewProductResponse(*x)
	return &v, nil
}
func (s *ProductServiceImpl) Create(c context.Context, r dto.CreateProductRequest) error {
	st := r.Status
	if st == "" {
		st = "active"
	}
	return s.repository.Create(c, entity.Product{ID: uuid.New(), CustomerID: r.CustomerID, SKU: r.SKU, Name: r.Name, UomID: r.UomID, Barcode: r.Barcode, Status: st, CreatedAt: time.Now()})
}
func (s *ProductServiceImpl) Update(c context.Context, id uuid.UUID, r dto.UpdateProductRequest) error {
	x, e := s.repository.FindByID(c, id)
	if e != nil {
		return e
	}
	if r.Name != nil {
		x.Name = *r.Name
	}
	if r.Barcode != nil {
		x.Barcode = r.Barcode
	}
	if r.Status != nil {
		x.Status = *r.Status
	}
	return s.repository.Update(c, x)
}
func (s *ProductServiceImpl) Delete(c context.Context, id uuid.UUID) error {
	return s.repository.Delete(c, id)
}
