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

type LocationHandler struct {
	service   services.LocationService
	validator *validator.Validate
	log       *logrus.Logger
}

func NewLocationHandler(service services.LocationService, validator *validator.Validate, log *logrus.Logger) *LocationHandler {
	return &LocationHandler{service: service, validator: validator, log: log}
}

func (h *LocationHandler) FindAll(c fiber.Ctx) error {
	filter, err := masterFilter(c)
	if err != nil {
		return responses.HandleError(c, err)
	}
	locations, total, err := h.service.FindAll(c.Context(), filter)
	if err != nil {
		return responses.HandleError(c, err)
	}
	return responses.NewSuccessResponseWithPagination(c, fiber.StatusOK, "Locations retrieved successfully", locations, dto.NewPagination(filter, total))
}

func (h *LocationHandler) FindByID(c fiber.Ctx) error {
	id, err := masterID(c)
	if err != nil {
		return responses.HandleError(c, err)
	}
	location, err := h.service.FindByID(c.Context(), id)
	if err != nil {
		return responses.HandleError(c, err)
	}
	return responses.NewSuccessResponse(c, fiber.StatusOK, "Location retrieved successfully", location)
}

func (h *LocationHandler) Create(c fiber.Ctx) error {
	var request dto.CreateLocationRequest
	if err := c.Bind().Body(&request); err != nil {
		return responses.HandleError(c, fiber.ErrBadRequest)
	}
	if err := h.validator.Struct(&request); err != nil {
		return responses.NewErrorResponse(c, fiber.StatusBadRequest, fiber.ErrBadRequest.Message, validator_utils.FormatValidationErrors(err))
	}
	if err := h.service.Create(c.Context(), request); err != nil {
		return responses.HandleError(c, err)
	}
	return responses.NewSuccessResponse(c, fiber.StatusCreated, "Location created successfully", nil)
}

func (h *LocationHandler) Update(c fiber.Ctx) error {
	id, err := masterID(c)
	if err != nil {
		return responses.HandleError(c, err)
	}
	var request dto.UpdateLocationRequest
	if err := c.Bind().Body(&request); err != nil {
		return responses.HandleError(c, fiber.ErrBadRequest)
	}
	if err := h.validator.Struct(&request); err != nil {
		return responses.NewErrorResponse(c, fiber.StatusBadRequest, fiber.ErrBadRequest.Message, validator_utils.FormatValidationErrors(err))
	}
	if err := h.service.Update(c.Context(), id, request); err != nil {
		return responses.HandleError(c, err)
	}
	return responses.NewSuccessResponse(c, fiber.StatusOK, "Location updated successfully", nil)
}

func (h *LocationHandler) Delete(c fiber.Ctx) error {
	id, err := masterID(c)
	if err != nil {
		return responses.HandleError(c, err)
	}
	if err := h.service.Delete(c.Context(), id); err != nil {
		return responses.HandleError(c, err)
	}
	return responses.NewSuccessResponse(c, fiber.StatusOK, "Location deleted successfully", nil)
}
