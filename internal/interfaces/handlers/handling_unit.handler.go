package handlers

import (
	"github.com/Ardnh/be-warehouse-management/internal/application/dto"
	"github.com/Ardnh/be-warehouse-management/internal/domain/services"
	responses "github.com/Ardnh/be-warehouse-management/internal/interfaces/response"
	validator_utils "github.com/Ardnh/be-warehouse-management/internal/utils/validator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
)

type HandlingUnitHandler struct {
	service   services.HandlingUnitService
	validator *validator.Validate
	log       *logrus.Logger
}

func NewHandlingUnitHandler(service services.HandlingUnitService, validator *validator.Validate, log *logrus.Logger) *HandlingUnitHandler {
	return &HandlingUnitHandler{service: service, validator: validator, log: log}
}
func (h *HandlingUnitHandler) FindAll(c fiber.Ctx) error {
	filter, err := masterFilter(c)
	if err != nil {
		return responses.HandleError(c, err)
	}
	result, total, err := h.service.FindAll(c.Context(), filter)
	if err != nil {
		return responses.HandleError(c, err)
	}
	return responses.NewSuccessResponseWithPagination(c, fiber.StatusOK, "Handling units retrieved successfully", result, dto.NewPagination(filter, total))
}
func (h *HandlingUnitHandler) FindByID(c fiber.Ctx) error {
	id, err := masterID(c)
	if err != nil {
		return responses.HandleError(c, err)
	}
	result, err := h.service.FindByID(c.Context(), id)
	if err != nil {
		return responses.HandleError(c, err)
	}
	return responses.NewSuccessResponse(c, fiber.StatusOK, "Handling unit retrieved successfully", result)
}
func (h *HandlingUnitHandler) Create(c fiber.Ctx) error {
	var request dto.CreateHandlingUnitRequest
	if err := c.Bind().Body(&request); err != nil {
		return responses.HandleError(c, fiber.ErrBadRequest)
	}
	if err := h.validator.Struct(&request); err != nil {
		return responses.NewErrorResponse(c, fiber.StatusBadRequest, fiber.ErrBadRequest.Message, validator_utils.FormatValidationErrors(err))
	}
	if err := h.service.Create(c.Context(), request); err != nil {
		return responses.HandleError(c, err)
	}
	return responses.NewSuccessResponse(c, fiber.StatusCreated, "Handling unit created successfully", nil)
}
func (h *HandlingUnitHandler) Update(c fiber.Ctx) error {
	id, err := masterID(c)
	if err != nil {
		return responses.HandleError(c, err)
	}
	var request dto.UpdateHandlingUnitRequest
	if err = c.Bind().Body(&request); err != nil {
		return responses.HandleError(c, fiber.ErrBadRequest)
	}
	if err = h.validator.Struct(&request); err != nil {
		return responses.NewErrorResponse(c, fiber.StatusBadRequest, fiber.ErrBadRequest.Message, validator_utils.FormatValidationErrors(err))
	}
	if err = h.service.Update(c.Context(), id, request); err != nil {
		return responses.HandleError(c, err)
	}
	return responses.NewSuccessResponse(c, fiber.StatusOK, "Handling unit updated successfully", nil)
}
func (h *HandlingUnitHandler) Delete(c fiber.Ctx) error {
	id, err := masterID(c)
	if err != nil {
		return responses.HandleError(c, err)
	}
	if err = h.service.Delete(c.Context(), id); err != nil {
		return responses.HandleError(c, err)
	}
	return responses.NewSuccessResponse(c, fiber.StatusOK, "Handling unit deleted successfully", nil)
}
