package repositories

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"github.com/google/uuid"
)

type UserRepository interface {
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	GetByUsername(ctx context.Context, username string) (*entity.User, error)
	GetByID(ctx context.Context, userID uuid.UUID) (*entity.User, error)
	Create(ctx context.Context, user entity.User) error
	Update(ctx context.Context, user entity.User) error
	Delete(ctx context.Context, userID uuid.UUID) error
}
