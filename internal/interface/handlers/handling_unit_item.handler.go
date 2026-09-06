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

type HandlingUnitItemHandler struct {
	service   services.HandlingUnitItemService
	validator *validator.Validate
	log       *logrus.Logger
}

func NewHandlingUnitItemHandler(service services.HandlingUnitItemService, validator *validator.Validate, log *logrus.Logger) *HandlingUnitItemHandler {
	return &HandlingUnitItemHandler{service: service, validator: validator, log: log}
}
func (h *HandlingUnitItemHandler) FindAll(c fiber.Ctx) error {
	unitID, err := masterID(c)
	if err != nil {
		return responses.HandleError(c, err)
	}
	result, err := h.service.FindAll(c.Context(), unitID)
	if err != nil {
		return responses.HandleError(c, err)
	}
	return responses.NewSuccessResponse(c, fiber.StatusOK, "Handling unit items retrieved successfully", result)
}
func (h *HandlingUnitItemHandler) FindByID(c fiber.Ctx) error {
	id, err := masterParamID(c, "item_id")
	if err != nil {
		return responses.HandleError(c, err)
	}
	result, err := h.service.FindByID(c.Context(), id)
	if err != nil {
		return responses.HandleError(c, err)
	}
	return responses.NewSuccessResponse(c, fiber.StatusOK, "Handling unit item retrieved successfully", result)
}
func (h *HandlingUnitItemHandler) Create(c fiber.Ctx) error {
	unitID, err := masterID(c)
	if err != nil {
		return responses.HandleError(c, err)
	}
	var request dto.AddHandlingUnitItemRequest
	if err = c.Bind().Body(&request); err != nil {
		return responses.HandleError(c, fiber.ErrBadRequest)
	}
	if err = h.validator.Struct(&request); err != nil {
		return responses.NewErrorResponse(c, fiber.StatusBadRequest, fiber.ErrBadRequest.Message, validator_utils.FormatValidationErrors(err))
	}
	if err = h.service.Create(c.Context(), unitID, request); err != nil {
		return responses.HandleError(c, err)
	}
	return responses.NewSuccessResponse(c, fiber.StatusCreated, "Handling unit item created successfully", nil)
}
func (h *HandlingUnitItemHandler) Update(c fiber.Ctx) error {
	id, err := masterParamID(c, "item_id")
	if err != nil {
		return responses.HandleError(c, err)
	}
	var request dto.UpdateHandlingUnitItemRequest
	if err = c.Bind().Body(&request); err != nil {
		return responses.HandleError(c, fiber.ErrBadRequest)
	}
	if err = h.validator.Struct(&request); err != nil {
		return responses.NewErrorResponse(c, fiber.StatusBadRequest, fiber.ErrBadRequest.Message, validator_utils.FormatValidationErrors(err))
	}
	if err = h.service.Update(c.Context(), id, request); err != nil {
		return responses.HandleError(c, err)
	}
	return responses.NewSuccessResponse(c, fiber.StatusOK, "Handling unit item updated successfully", nil)
}
func (h *HandlingUnitItemHandler) Delete(c fiber.Ctx) error {
	id, err := masterParamID(c, "item_id")
	if err != nil {
		return responses.HandleError(c, err)
	}
	if err = h.service.Delete(c.Context(), id); err != nil {
		return responses.HandleError(c, err)
	}
	return responses.NewSuccessResponse(c, fiber.StatusOK, "Handling unit item deleted successfully", nil)
}
