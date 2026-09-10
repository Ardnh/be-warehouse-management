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

func (r *UserRepositoryImpl) FindAll(ctx context.Context, filter repositories.Filter) ([]*entity.User, int64, error) {
	var users []*entity.User
	var total int64

	base := r.db.WithContext(ctx).Model(&entity.User{})
	if filter.Search != "" {
		base = base.Where("username ILIKE ? OR email ILIKE ?", "%"+filter.Search+"%", "%"+filter.Search+"%")
	}
	if err := base.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	base = applyMasterListFilter(base, filter)
	if err := base.Find(&users).Error; err != nil {
		return nil, 0, err
	}
	return users, total, nil
}

func (r *UserRepositoryImpl) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) FindByUsername(ctx context.Context, username string) (*entity.User, error) {
	var user entity.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) FindByID(ctx context.Context, userID uuid.UUID) (*entity.User, error) {
	var user entity.User
	if err := r.db.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) Create(ctx context.Context, user entity.User) error {
	if err := r.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepositoryImpl) Update(ctx context.Context, user entity.User) error {
	if err := r.db.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepositoryImpl) Delete(ctx context.Context, userID uuid.UUID) error {
	if err := r.db.Delete(&entity.User{}, userID).Error; err != nil {
		return err
	}
	return nil
}
