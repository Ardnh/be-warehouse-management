package repositories

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"github.com/Ardnh/be-warehouse-management/internal/domain/repositories"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewUserRepository(db *gorm.DB, redis *redis.Client) repositories.UserRepository {
	return &UserRepositoryImpl{
		db:    db,
		redis: redis,
	}
}

func (r *UserRepositoryImpl) GetByEmail(ctx context.Context, email string) (*entity.User, error)
func (r *UserRepositoryImpl) GetByID(ctx context.Context, userID uuid.UUID) (*entity.User, error)
func (r *UserRepositoryImpl) Create(ctx context.Context, user entity.User) error
func (r *UserRepositoryImpl) Update(ctx context.Context, user entity.User) error
func (r *UserRepositoryImpl) Delete(ctx context.Context, userID uuid.UUID) error
