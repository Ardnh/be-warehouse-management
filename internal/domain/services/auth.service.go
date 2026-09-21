package services

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/application/dto"
	"github.com/google/uuid"
)

type AuthService interface {
	Login(ctx context.Context, req dto.LoginRequestDto) (*dto.LoginTempResponseDto, error)
	Register(ctx context.Context, req dto.RegisterRequestDto) error
	LoginSelection(ctx context.Context, userID uuid.UUID, warehouseID uuid.UUID) (*dto.LoginTempResponseDto, error)
}
