package repositories

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	domainrepositories "github.com/Ardnh/be-warehouse-management/internal/domain/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserAssignmentRepositoryImpl struct {
	db *gorm.DB
}

func NewUserAssignmentRepository(db *gorm.DB) domainrepositories.UserAssignmentRepository {
	return &UserAssignmentRepositoryImpl{db: db}
}

func (r *UserAssignmentRepositoryImpl) AssignRoles(ctx context.Context, userID uuid.UUID, locationID uuid.UUID, roleIDs []uuid.UUID) error {
	assignments := make([]entity.UserAssignment, 0, len(roleIDs))
	for _, roleID := range roleIDs {
		assignments = append(assignments, entity.UserAssignment{
			UserID:     userID,
			LocationID: locationID,
			RoleID:     roleID,
		})
	}
	if len(assignments) == 0 {
		return nil
	}
	return Conn(ctx, r.db).Create(&assignments).Error
}

func (r *UserAssignmentRepositoryImpl) FindByUserID(ctx context.Context, userID uuid.UUID) ([]entity.UserAssignment, error) {
	var assignments []entity.UserAssignment
	if err := Conn(ctx, r.db).
		Preload("Role").
		Preload("Location").
		Where("user_id = ?", userID).
		Find(&assignments).Error; err != nil {
		return nil, err
	}
	return assignments, nil
}

func (r *UserAssignmentRepositoryImpl) FindByRoleID(ctx context.Context, roleID uuid.UUID) ([]entity.UserAssignment, error) {
	var assignments []entity.UserAssignment
	if err := Conn(ctx, r.db).
		Preload("User").
		Preload("Location").
		Where("role_id = ?", roleID).
		Find(&assignments).Error; err != nil {
		return nil, err
	}
	return assignments, nil
}

func (r *UserAssignmentRepositoryImpl) FindByUserAndLocation(ctx context.Context, userID uuid.UUID, locationID uuid.UUID) (*entity.UserAssignment, error) {
	var assignment entity.UserAssignment
	if err := Conn(ctx, r.db).
		Preload("Role").
		Preload("Location").
		Where("user_id = ? AND location_id = ?", userID, locationID).
		First(&assignment).Error; err != nil {
		return nil, err
	}
	return &assignment, nil
}

func (r *UserAssignmentRepositoryImpl) FindByID(ctx context.Context, assignmentID uuid.UUID) (*entity.UserAssignment, error) {
	var assignment entity.UserAssignment
	if err := Conn(ctx, r.db).
		Preload("User").
		Preload("Role").
		Preload("Location").
		First(&assignment, assignmentID).Error; err != nil {
		return nil, err
	}
	return &assignment, nil
}

func (r *UserAssignmentRepositoryImpl) Create(ctx context.Context, assignment entity.UserAssignment) error {
	return Conn(ctx, r.db).Create(&assignment).Error
}

func (r *UserAssignmentRepositoryImpl) Delete(ctx context.Context, assignmentID uuid.UUID) error {
	return Conn(ctx, r.db).Delete(&entity.UserAssignment{}, assignmentID).Error
}

func (r *UserAssignmentRepositoryImpl) DeleteByUserID(ctx context.Context, userID uuid.UUID) error {
	return Conn(ctx, r.db).Where("user_id = ?", userID).Delete(&entity.UserAssignment{}).Error
}
