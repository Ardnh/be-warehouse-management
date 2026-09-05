package handlers

import (
	"github.com/Ardnh/be-warehouse-management/internal/application/dto"
	"github.com/Ardnh/be-warehouse-management/internal/domain/services"
	responses "github.com/Ardnh/be-warehouse-management/internal/interface/response"
	validator_utils "github.com/Ardnh/be-warehouse-management/internal/utils/validator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
)

type AuthHandler struct {
	authService services.AuthService
	validator   *validator.Validate
	log         *logrus.Logger
}

func NewAuthHandler(authService services.AuthService, validator *validator.Validate, log *logrus.Logger) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		validator:   validator,
		log:         log,
	}
}

func (h *AuthHandler) Login(c fiber.Ctx) error {
	var req dto.LoginRequestDto
	if err := c.Bind().Body(&req); err != nil {
		return responses.HandleError(c, err)
	}

	if err := h.validator.Struct(&req); err != nil {
		return responses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, validator_utils.FormatValidationErrors(err))
	}

	result, err := h.authService.Login(c.Context(), req)
	if err != nil {
		return responses.HandleError(c, err)
	}

	return responses.NewSuccessResponse(c, fiber.StatusOK, "Login successful", result)
}

func (h *AuthHandler) Register(c fiber.Ctx) error {
	var req dto.RegisterRequestDto
	if err := c.Bind().Body(&req); err != nil {
		return responses.HandleError(c, err)
	}

	if err := h.validator.Struct(&req); err != nil {
		return responses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, validator_utils.FormatValidationErrors(err))
	}

	err := h.authService.Register(c.Context(), req)
	if err != nil {
		return responses.HandleError(c, err)
	}

	return responses.NewSuccessResponse(c, fiber.StatusCreated, "Register successful", nil)
}
