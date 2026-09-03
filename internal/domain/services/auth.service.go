package services

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/application/dto"
)

type AuthService interface {
	Login(ctx context.Context, req dto.LoginRequestDto) (*dto.LoginResponseDto, error)
	Register(ctx context.Context, req dto.RegisterRequestDto) error
}
