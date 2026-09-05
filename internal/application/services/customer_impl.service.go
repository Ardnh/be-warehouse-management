package services

import (
	"context"
	"time"

	"github.com/Ardnh/be-warehouse-management/internal/application/dto"
	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"github.com/Ardnh/be-warehouse-management/internal/domain/repositories"
	"github.com/Ardnh/be-warehouse-management/internal/domain/services"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type CustomerServiceImpl struct {
	customerRepository repositories.CustomerRepository
	log                *logrus.Logger
}

func NewCustomerService(customerRepository repositories.CustomerRepository, log *logrus.Logger) services.CustomerService {
	return &CustomerServiceImpl{
		customerRepository: customerRepository,
		log:                log,
	}
}

func (s *CustomerServiceImpl) FindAll(ctx context.Context, filter dto.FilterDTO) ([]dto.Customer, int64, error) {

	filterRepo := repositories.Filter{
		Search:   filter.Search,
		Page:     filter.Page,
		PageSize: filter.Size,
		SortBy:   filter.SortBy,
		SortDir:  filter.SortDir,
	}

	customers, total, err := s.customerRepository.FindAll(ctx, filterRepo)
	if err != nil {
		return nil, 0, err
	}

	customerDTOs := dto.ToCustomerDTOs(customers)
	return customerDTOs, total, nil
}

func (s *CustomerServiceImpl) FindByID(ctx context.Context, id uuid.UUID) (*dto.Customer, error) {
	customer, err := s.customerRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	customerDto := dto.ToCustomerDTO(customer)
	return &customerDto, nil
}

func (s *CustomerServiceImpl) Create(ctx context.Context, customer dto.CreateCustomerRequest) error {

	customerEntity := entity.Customer{
		ID:        uuid.New(),
		Code:      customer.Code,
		Name:      customer.Name,
		Email:     customer.Email,
		Address:   customer.Address,
		Phone:     customer.Phone,
		Status:    customer.Status,
		CreatedAt: time.Now(),
	}

	err := s.customerRepository.Create(ctx, customerEntity)
	if err != nil {
		return err
	}

	return nil
}

func (s *CustomerServiceImpl) Update(ctx context.Context, id uuid.UUID, customer dto.UpdateCustomerRequest) error {

	customerById, err := s.customerRepository.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if customer.Name != "" {
		customerById.Name = customer.Name
	}
	if customer.Address != "" {
		customerById.Address = customer.Address
	}
	if customer.Status != "" {
		customerById.Status = customer.Status
	}
	if customer.Email != "" {
		customerById.Email = customer.Email
	}
	if customer.Phone != "" {
		customerById.Phone = customer.Phone
	}

	errUpdate := s.customerRepository.Update(ctx, customerById)
	if errUpdate != nil {
		return err
	}

	return nil
}

func (s *CustomerServiceImpl) Delete(ctx context.Context, id uuid.UUID) error {
	err := s.customerRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
