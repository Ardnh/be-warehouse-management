package repositories

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"github.com/google/uuid"
)

type UserAssignmentRepository interface {
	AssignRoles(ctx context.Context, userID uuid.UUID, locationID uuid.UUID, roleIDs []uuid.UUID) error
	FindByUserID(ctx context.Context, userID uuid.UUID) ([]entity.UserAssignment, error)
	FindByRoleID(ctx context.Context, roleID uuid.UUID) ([]entity.UserAssignment, error)
	FindByUserAndLocation(ctx context.Context, userID uuid.UUID, locationID uuid.UUID) (*entity.UserAssignment, error)
	FindByID(ctx context.Context, assignmentID uuid.UUID) (*entity.UserAssignment, error)
	Create(ctx context.Context, assignment entity.UserAssignment) error
	Delete(ctx context.Context, assignmentID uuid.UUID) error
	DeleteByUserID(ctx context.Context, userID uuid.UUID) error
}
