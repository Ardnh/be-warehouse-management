package services

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/application/dto"
	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"github.com/Ardnh/be-warehouse-management/internal/domain/repositories"
	"github.com/Ardnh/be-warehouse-management/internal/domain/services"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	UserRepository     repositories.UserRepository
	UserRoleRepository repositories.UserRoleRepository
	tx                 repositories.TxManager
	log                *logrus.Logger
}

func NewUserService(userRepository repositories.UserRepository, userRoleRepository repositories.UserRoleRepository, log *logrus.Logger, tx repositories.TxManager) services.UserService {
	return &UserServiceImpl{
		UserRepository:     userRepository,
		UserRoleRepository: userRoleRepository,
		tx:                 tx,
		log:                log,
	}
}

func (s *UserServiceImpl) FindAll(ctx context.Context, filterDto dto.FilterDTO) ([]dto.User, int64, error) {
	filter := repositories.Filter{
		Search:   filterDto.Search,
		Page:     filterDto.Page,
		PageSize: filterDto.Size,
		SortBy:   filterDto.SortBy,
		SortDir:  filterDto.SortDir,
	}
	users, count, err := s.UserRepository.FindAll(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	userDtos := dto.ToUserDTOs(users)
	return userDtos, count, nil
}

func (s *UserServiceImpl) FindByEmail(ctx context.Context, email string) (*dto.User, error) {
	user, err := s.UserRepository.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	userDto := dto.ToUserDTO(user)
	return &userDto, nil
}

func (s *UserServiceImpl) FindByID(ctx context.Context, userID uuid.UUID) (*dto.User, error) {
	user, err := s.UserRepository.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	userDto := dto.ToUserDTO(user)
	return &userDto, nil
}

func (s *UserServiceImpl) Create(ctx context.Context, user dto.CreateUserRequest) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		s.log.WithField("email", user.Email).Error("failed to hash password")
		return fiber.ErrInternalServerError
	}

	return s.tx.Do(ctx, func(ctx context.Context) error {

		userEntity := entity.User{
			Username:     user.Username,
			Email:        user.Email,
			PasswordHash: string(hashedPassword),
			FullName:     user.FullName,
			Status:       user.Status,
		}
		err := s.UserRepository.Create(ctx, userEntity)
		if err != nil {
			return err
		}

		return s.UserRoleRepository.AssignRoles(ctx, userEntity.ID, user.RoleIDs)
	})
}

func (s *UserServiceImpl) Update(ctx context.Context, id uuid.UUID, user dto.UpdateUserRequest) error {

	userEntity, err := s.UserRepository.FindByID(ctx, id)
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
