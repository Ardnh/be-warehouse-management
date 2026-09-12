package repositories

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	domainrepositories "github.com/Ardnh/be-warehouse-management/internal/domain/repositories"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewUserRepository(db *gorm.DB, redis *redis.Client) domainrepositories.UserRepository {
	return &UserRepositoryImpl{db: db, redis: redis}
}

func (r *UserRepositoryImpl) FindAll(ctx context.Context, filter domainrepositories.Filter) ([]*entity.User, int64, error) {
	var users []*entity.User
	var total int64

	base := Conn(ctx, r.db).Model(&entity.User{})
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
	if err := Conn(ctx, r.db).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) FindByUsername(ctx context.Context, username string) (*entity.User, error) {
	var user entity.User
	if err := Conn(ctx, r.db).Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) FindByID(ctx context.Context, userID uuid.UUID) (*entity.User, error) {
	var user entity.User
	if err := Conn(ctx, r.db).
		Preload("Roles").
		Preload("UserRoles.Role").
		Preload("UserRoles.Warehouse").
		Preload("UserPermissions.Permission").
		Where("id = ?", userID).
		First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) Create(ctx context.Context, user entity.User) error {
	return Conn(ctx, r.db).Create(&user).Error
}

func (r *UserRepositoryImpl) Update(ctx context.Context, user entity.User) error {
	return Conn(ctx, r.db).Save(&user).Error
}

func (r *UserRepositoryImpl) Delete(ctx context.Context, userID uuid.UUID) error {
	return Conn(ctx, r.db).Delete(&entity.User{}, userID).Error
}
