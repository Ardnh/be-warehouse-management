package services

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/application/dto"
	"github.com/google/uuid"
)

type UserService interface {
	FindAll(ctx context.Context, filter dto.FilterDTO) ([]dto.User, int64, error)
	FindByEmail(ctx context.Context, email string) (*dto.User, error)
	FindByID(ctx context.Context, userID uuid.UUID) (*dto.User, error)
	Create(ctx context.Context, user dto.CreateUserRequest) error
	Update(ctx context.Context, id uuid.UUID, user dto.UpdateUserRequest) error
	Delete(ctx context.Context, userID uuid.UUID) error
}
