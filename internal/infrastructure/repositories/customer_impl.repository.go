package repositories

import (
	"context"
	"fmt"

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

func (r *CustomerRepositoryImpl) FindAll(ctx context.Context, filter repositories.Filter) ([]entity.Customer, int64, error) {

	var (
		customers []entity.Customer
		total     int64
	)

	// --- Default pagination guard
	if filter.Page <= 0 {
		filter.Page = 1
	}
	if filter.PageSize <= 0 || filter.PageSize >= 1000 {
		filter.PageSize = 30
	}

	offset := (filter.Page - 1) * filter.PageSize

	// --- Base query TANPA preload (untuk count & filter)
	baseQuery := r.db.WithContext(ctx).
		Model(&entity.Customer{})

	// --- Search (by name)
	if filter.Search != "" {
		baseQuery = baseQuery.Where("name ILIKE ?", "%"+filter.Search+"%")
	}

	// --- Hitung total data dulu (WAJIB sebelum limit)
	if err := baseQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// --- Limit & offset
	baseQuery = baseQuery.Offset(offset).Limit(filter.PageSize)

	// --- Fetch data
	if err := baseQuery.Find(&customers).Error; err != nil {
		return nil, 0, err
	}

	return customers, total, nil
}

func (r *CustomerRepositoryImpl) FindByID(ctx context.Context, id uuid.UUID) (*entity.Customer, error) {
	var customer entity.Customer

	if err := r.db.WithContext(ctx).First(&customer, id).Error; err != nil {
		return nil, err
	}

	return &customer, nil
}

func (r *CustomerRepositoryImpl) Create(ctx context.Context, customer entity.Customer) error {
	if err := r.db.WithContext(ctx).Create(&customer).Error; err != nil {
		return err
	}
	return nil
}

func (r *CustomerRepositoryImpl) Update(ctx context.Context, customer *entity.Customer) error {
	if err := r.db.WithContext(ctx).Save(&customer).Error; err != nil {
		return err
	}
	return nil
}

func (r *CustomerRepositoryImpl) Delete(ctx context.Context, customer uuid.UUID) error {
	if err := r.db.WithContext(ctx).Delete(&customer).Error; err != nil {
		return err
	}
	return nil
}

func GenerateCustomerCode(tx *gorm.DB) (string, error) {
	var sequence int64

	err := tx.Raw(
		"SELECT nextval('customer_code_seq')",
	).Scan(&sequence).Error

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("CST%04d", sequence), nil
}
