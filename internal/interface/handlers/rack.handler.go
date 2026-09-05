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

type RackHandler struct {
	service   services.RackService
	validator *validator.Validate
	log       *logrus.Logger
}

func NewRackHandler(service services.RackService, validator *validator.Validate, log *logrus.Logger) *RackHandler {
	return &RackHandler{service: service, validator: validator, log: log}
}
func (h *RackHandler) FindAll(c fiber.Ctx) error {
	filter, err := masterFilter(c)
	if err != nil {
		return responses.HandleError(c, err)
	}
	result, total, err := h.service.FindAll(c.Context(), filter)
	if err != nil {
		return responses.HandleError(c, err)
	}
	return responses.NewSuccessResponseWithPagination(c, fiber.StatusOK, "Racks retrieved successfully", result, dto.NewPagination(filter, total))
}
func (h *RackHandler) FindByID(c fiber.Ctx) error {
	id, err := masterID(c)
	if err != nil {
		return responses.HandleError(c, err)
	}
	result, err := h.service.FindByID(c.Context(), id)
	if err != nil {
		return responses.HandleError(c, err)
	}
	return responses.NewSuccessResponse(c, fiber.StatusOK, "Rack retrieved successfully", result)
}
func (h *RackHandler) Create(c fiber.Ctx) error {
	var request dto.CreateRackRequest
	if err := c.Bind().Body(&request); err != nil {
		return responses.HandleError(c, fiber.ErrBadRequest)
	}
	if err := h.validator.Struct(&request); err != nil {
		return responses.NewErrorResponse(c, fiber.StatusBadRequest, fiber.ErrBadRequest.Message, validator_utils.FormatValidationErrors(err))
	}
	if err := h.service.Create(c.Context(), request); err != nil {
		return responses.HandleError(c, err)
	}
	return responses.NewSuccessResponse(c, fiber.StatusCreated, "Rack created successfully", nil)
}
func (h *RackHandler) Update(c fiber.Ctx) error {
	id, err := masterID(c)
	if err != nil {
		return responses.HandleError(c, err)
	}
	var request dto.UpdateRackRequest
	if err := c.Bind().Body(&request); err != nil {
		return responses.HandleError(c, fiber.ErrBadRequest)
	}
	if err := h.validator.Struct(&request); err != nil {
		return responses.NewErrorResponse(c, fiber.StatusBadRequest, fiber.ErrBadRequest.Message, validator_utils.FormatValidationErrors(err))
	}
	if err := h.service.Update(c.Context(), id, request); err != nil {
		return responses.HandleError(c, err)
	}
	return responses.NewSuccessResponse(c, fiber.StatusOK, "Rack updated successfully", nil)
}
func (h *RackHandler) Delete(c fiber.Ctx) error {
	id, err := masterID(c)
	if err != nil {
		return responses.HandleError(c, err)
	}
	if err := h.service.Delete(c.Context(), id); err != nil {
		return responses.HandleError(c, err)
	}
	return responses.NewSuccessResponse(c, fiber.StatusOK, "Rack deleted successfully", nil)
}
