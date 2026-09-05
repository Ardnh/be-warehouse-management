package services

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/application/dto"
	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"github.com/Ardnh/be-warehouse-management/internal/domain/repositories"
	"github.com/Ardnh/be-warehouse-management/internal/domain/services"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type UserServiceImpl struct {
	UserRepository repositories.UserRepository
	log            *logrus.Logger
}

func NewUserService(userRepository repositories.UserRepository, log *logrus.Logger) services.UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		log:            log,
	}
}

func (s *UserServiceImpl) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	user, err := s.UserRepository.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserServiceImpl) FindByID(ctx context.Context, userID uuid.UUID) (*entity.User, error) {
	user, err := s.UserRepository.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserServiceImpl) Create(ctx context.Context, user dto.CreateUserRequest) error {
	userEntity := entity.User{
		Username:     user.Username,
		Email:        user.Email,
		PasswordHash: user.Password,
		FullName:     user.FullName,
		Status:       user.Status,
	}
	err := s.UserRepository.Create(ctx, userEntity)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserServiceImpl) Update(ctx context.Context, id uuid.UUID, user dto.UpdateUserRequest) error {

	userEntity, err := s.UserRepository.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if user.Email != nil {
		userEntity.Email = *user.Email
	}
	if user.FullName != nil {
		userEntity.FullName = *user.FullName
	}
	if user.Status != nil {
		userEntity.Status = *user.Status
	}

	err = s.UserRepository.Update(ctx, *userEntity)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserServiceImpl) Delete(ctx context.Context, userID uuid.UUID) error {
	err := s.UserRepository.Delete(ctx, userID)
	if err != nil {
		return err
	}
	return nil
}
