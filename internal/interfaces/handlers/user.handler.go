package handlers

import (
	"github.com/Ardnh/be-warehouse-management/internal/application/dto"
	"github.com/Ardnh/be-warehouse-management/internal/domain/services"
	responses "github.com/Ardnh/be-warehouse-management/internal/interface/response"
	validator_utils "github.com/Ardnh/be-warehouse-management/internal/utils/validator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	UserService services.UserService
	validator   *validator.Validate
	log         *logrus.Logger
}

func NewUserHandler(userService services.UserService, validator *validator.Validate, log *logrus.Logger) *UserHandler {
	return &UserHandler{
		UserService: userService,
		validator:   validator,
		log:         log,
	}
}

func (h *UserHandler) findByID(c fiber.Ctx) error {
	userID := c.Params("id", "")

	if userID == "" {
		return responses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, "Id is required")
	}

	userIDUuid, err := uuid.Parse(userID)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid id")
	}

	user, err := h.UserService.FindByID(c.Context(), userIDUuid)
	if err != nil {
		return err
	}

	if user == nil {
		return responses.NewErrorResponse(c, fiber.ErrNotFound.Code, fiber.ErrNotFound.Message, "User not found")
	}

	return responses.NewSuccessResponse(c, fiber.StatusOK, "User found", user)
}

func (h *UserHandler) create(c fiber.Ctx) error {

	var request dto.CreateUserRequest
	if err := c.Bind().Body(&request); err != nil {
		return responses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, err)
	}

	if err := h.validator.Struct(&request); err != nil {
		return responses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, validator_utils.FormatValidationErrors(err))
	}

	err := h.UserService.Create(c.Context(), request)
	if err != nil {
		return responses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}

	return responses.NewSuccessResponse(c, fiber.StatusOK, "User created", nil)
}

func (h *UserHandler) update(c fiber.Ctx) error {
	userID := c.Params("id", "")

	if userID == "" {
		return responses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, "Id is required")
	}

	userIDUuid, err := uuid.Parse(userID)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid id")
	}

	var request dto.UpdateUserRequest
	if err := c.Bind().Body(&request); err != nil {
		return responses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, err)
	}

	if err := h.validator.Struct(&request); err != nil {
		return responses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, validator_utils.FormatValidationErrors(err))
	}

	err = h.UserService.Update(c.Context(), userIDUuid, request)
	if err != nil {
		return responses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}

	return responses.NewSuccessResponse(c, fiber.StatusOK, "User updated", nil)
}

func (h *UserHandler) delete(c fiber.Ctx) error {
	userID := c.Params("id", "")

	if userID == "" {
		return responses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, "Id is required")
	}

	userIDUuid, err := uuid.Parse(userID)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid id")
	}

	err = h.UserService.Delete(c.Context(), userIDUuid)
	if err != nil {
		return responses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}

	return responses.NewSuccessResponse(c, fiber.StatusOK, "User deleted", nil)
}
