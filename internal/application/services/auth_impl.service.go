package services

import (
	"context"
	"errors"

	"github.com/Ardnh/be-warehouse-management/internal/application/dto"
	"github.com/Ardnh/be-warehouse-management/internal/config"
	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"github.com/Ardnh/be-warehouse-management/internal/domain/repositories"
	"github.com/Ardnh/be-warehouse-management/internal/domain/services"
	utils "github.com/Ardnh/be-warehouse-management/internal/utils/jwt"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	userRepository repositories.UserRepository
	log            *logrus.Logger
	appConfig      *config.Config
}

func NewAuthService(userRepository repositories.UserRepository, log *logrus.Logger, appConfig *config.Config) services.AuthService {
	return &AuthServiceImpl{
		userRepository: userRepository,
		log:            log,
		appConfig:      appConfig,
	}
}

func (s *AuthServiceImpl) Login(ctx context.Context, req dto.LoginRequestDto) (*dto.LoginResponseDto, error) {
	user, err := s.userRepository.FindByUsername(ctx, req.Username)
	if err != nil {
		if errors.Is(err, fiber.ErrNotFound) {
			s.log.WithField("username", req.Username).Warn("login attempt with unregistered username")
			return nil, fiber.ErrNotFound
		}
		return nil, fiber.ErrInternalServerError
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		s.log.WithField("username", req.Username).Warn("invalid credentials provided")
		return nil, fiber.ErrBadRequest
	}

	secretKey := []byte(s.appConfig.App.JWTSecret)
	token, expiredTimeISO, err := utils.GenerateToken(secretKey, user.ID.String())
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"user_id": user.ID.String(),
			"error":   err,
		}).Error("failed to generate token")
		return nil, fiber.ErrInternalServerError
	}

	s.log.WithField("user_id", user.ID).Info("login successful")

	return &dto.LoginResponseDto{
		Token:      *token,
		ExpireDate: *expiredTimeISO,
	}, nil
}

func (s *AuthServiceImpl) Register(ctx context.Context, req dto.RegisterRequestDto) error {
	existingUser, err := s.userRepository.FindByEmail(ctx, req.Email)
	if err != nil && !errors.Is(err, fiber.ErrNotFound) {
		s.log.WithFields(logrus.Fields{
			"email": req.Email,
			"error": err,
		}).Error("failed to check existing user")
		return fiber.ErrNotFound
	}

	if existingUser != nil {
		s.log.WithField("email", req.Email).Warn("registration attempt with existing email")
		return fiber.ErrConflict
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		s.log.WithField("email", req.Email).Error("failed to hash password")
		return fiber.ErrInternalServerError
	}

	user := entity.User{
		ID:           uuid.New(),
		FullName:     req.FullName,
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
	}

	err = s.userRepository.Create(ctx, user)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"email": req.Email,
			"error": err,
		}).Error("failed to create user")
		return fiber.ErrInternalServerError
	}

	// Add role grouping to Casbin
	// _, err = s.casbinEnforcer.AddGroupingPolicy(user.ID.String(), constants.RoleDailyUser)
	// if err != nil {
	// 	s.log.WithFields(logrus.Fields{
	// 		"email":  req.Email,
	// 		"userID": user.ID,
	// 		"error":  err,
	// 	}).Error("failed to add casbin grouping policy")
	// 	return fiber.ErrInternalServerError
	// }

	return nil
}
