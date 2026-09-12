package repositories

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"github.com/google/uuid"
)

type UserRoleRepository interface {
	AssignRoles(ctx context.Context, userID uuid.UUID, roleIDs []uuid.UUID) error
	FindByUserID(ctx context.Context, userID uuid.UUID) ([]entity.UserRole, error)
	FindByRoleID(ctx context.Context, roleID uuid.UUID) ([]entity.UserRole, error)
	FindByUserAndRole(ctx context.Context, userID uuid.UUID, roleID uuid.UUID) (*entity.UserRole, error)
	Create(ctx context.Context, userRole entity.UserRole) error
	Delete(ctx context.Context, userID uuid.UUID, roleID uuid.UUID) error
	DeleteByUserID(ctx context.Context, userID uuid.UUID) error
}
