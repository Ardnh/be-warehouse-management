package repositories

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	domainrepositories "github.com/Ardnh/be-warehouse-management/internal/domain/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRoleRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRoleRepository(db *gorm.DB) domainrepositories.UserRoleRepository {
	return &UserRoleRepositoryImpl{db: db}
}

func (r *UserRoleRepositoryImpl) AssignRoles(ctx context.Context, userID uuid.UUID, roleIDs []uuid.UUID) error {
	rows := make([]entity.UserRole, 0, len(roleIDs))
	for _, id := range roleIDs {
		rows = append(rows, entity.UserRole{UserID: userID, RoleID: id})
	}
	return Conn(ctx, r.db).Create(&rows).Error
}

func (r *UserRoleRepositoryImpl) FindByUserID(ctx context.Context, userID uuid.UUID) ([]entity.UserRole, error) {
	var userRoles []entity.UserRole
	if err := Conn(ctx, r.db).
		Preload("Role").
		Preload("Warehouse").
		Where("user_id = ?", userID).
		Find(&userRoles).Error; err != nil {
		return nil, err
	}
	return userRoles, nil
}

func (r *UserRoleRepositoryImpl) FindByRoleID(ctx context.Context, roleID uuid.UUID) ([]entity.UserRole, error) {
	var userRoles []entity.UserRole
	if err := Conn(ctx, r.db).
		Preload("User").
		Preload("Warehouse").
		Where("role_id = ?", roleID).
		Find(&userRoles).Error; err != nil {
		return nil, err
	}
	return userRoles, nil
}

func (r *UserRoleRepositoryImpl) FindByUserAndRole(ctx context.Context, userID uuid.UUID, roleID uuid.UUID) (*entity.UserRole, error) {
	var userRole entity.UserRole
	if err := Conn(ctx, r.db).
		Preload("Role").
		Preload("Warehouse").
		Where("user_id = ? AND role_id = ?", userID, roleID).
		First(&userRole).Error; err != nil {
		return nil, err
	}
	return &userRole, nil
}

func (r *UserRoleRepositoryImpl) Create(ctx context.Context, userRole entity.UserRole) error {
	return Conn(ctx, r.db).Create(&userRole).Error
}

func (r *UserRoleRepositoryImpl) Delete(ctx context.Context, userID uuid.UUID, roleID uuid.UUID) error {
	return Conn(ctx, r.db).
		Where("user_id = ? AND role_id = ?", userID, roleID).
		Delete(&entity.UserRole{}).Error
}

func (r *UserRoleRepositoryImpl) DeleteByUserID(ctx context.Context, userID uuid.UUID) error {
	return Conn(ctx, r.db).
		Where("user_id = ?", userID).
		Delete(&entity.UserRole{}).Error
}
