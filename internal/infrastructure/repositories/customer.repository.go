package repositories

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"github.com/Ardnh/be-warehouse-management/internal/domain/repositories"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type CustomerRepositoryImpl struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewCustomerRepository(db *gorm.DB, redis *redis.Client) repositories.CustomerRepository {
	return &CustomerRepositoryImpl{
		db:    db,
		redis: redis,
	}
}

func (r *CustomerRepositoryImpl) FindAll(ctx context.Context) ([]entity.Customer, error)
func (r *CustomerRepositoryImpl) FindByID(ctx context.Context, id uuid.UUID) (entity.Customer, error)
func (r *CustomerRepositoryImpl) Create(ctx context.Context, customer entity.Customer) error
func (r *CustomerRepositoryImpl) Update(ctx context.Context, customer entity.Customer) error
func (r *CustomerRepositoryImpl) Delete(ctx context.Context, customer entity.Customer) error
