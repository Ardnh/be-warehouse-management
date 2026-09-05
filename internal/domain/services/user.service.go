package services

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/application/dto"
	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"github.com/google/uuid"
)

type UserService interface {
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	FindByID(ctx context.Context, userID uuid.UUID) (*entity.User, error)
	Create(ctx context.Context, user dto.CreateUserRequest) error
	Update(ctx context.Context, id uuid.UUID, user dto.UpdateUserRequest) error
	Delete(ctx context.Context, userID uuid.UUID) error
}
