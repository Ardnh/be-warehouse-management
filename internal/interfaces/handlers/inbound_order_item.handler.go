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

type InboundOrderItemHandler struct {
	service   services.InboundOrderItemService
	validator *validator.Validate
	log       *logrus.Logger
}

func NewInboundOrderItemHandler(service services.InboundOrderItemService, validator *validator.Validate, log *logrus.Logger) *InboundOrderItemHandler {
	return &InboundOrderItemHandler{service: service, validator: validator, log: log}
}
func (h *InboundOrderItemHandler) FindAll(c fiber.Ctx) error {
	orderID, err := masterID(c)
	if err != nil {
		return responses.HandleError(c, err)
	}
	result, err := h.service.FindAll(c.Context(), orderID)
	if err != nil {
		return responses.HandleError(c, err)
	}
	return responses.NewSuccessResponse(c, fiber.StatusOK, "Inbound order items retrieved successfully", result)
}
func (h *InboundOrderItemHandler) FindByID(c fiber.Ctx) error {
	id, err := masterParamID(c, "item_id")
	if err != nil {
		return responses.HandleError(c, err)
	}
	result, err := h.service.FindByID(c.Context(), id)
	if err != nil {
		return responses.HandleError(c, err)
	}
	return responses.NewSuccessResponse(c, fiber.StatusOK, "Inbound order item retrieved successfully", result)
}
func (h *InboundOrderItemHandler) Create(c fiber.Ctx) error {
	orderID, err := masterID(c)
	if err != nil {
		return responses.HandleError(c, err)
	}
	var request dto.CreateInboundOrderItemRequest
	if err = c.Bind().Body(&request); err != nil {
		return responses.HandleError(c, fiber.ErrBadRequest)
	}
	if err = h.validator.Struct(&request); err != nil {
		return responses.NewErrorResponse(c, fiber.StatusBadRequest, fiber.ErrBadRequest.Message, validator_utils.FormatValidationErrors(err))
	}
	if err = h.service.Create(c.Context(), orderID, request); err != nil {
		return responses.HandleError(c, err)
	}
	return responses.NewSuccessResponse(c, fiber.StatusCreated, "Inbound order item created successfully", nil)
}
func (h *InboundOrderItemHandler) Update(c fiber.Ctx) error {
	id, err := masterParamID(c, "item_id")
	if err != nil {
		return responses.HandleError(c, err)
	}
	var request dto.UpdateInboundOrderItemRequest
	if err = c.Bind().Body(&request); err != nil {
		return responses.HandleError(c, fiber.ErrBadRequest)
	}
	if err = h.validator.Struct(&request); err != nil {
		return responses.NewErrorResponse(c, fiber.StatusBadRequest, fiber.ErrBadRequest.Message, validator_utils.FormatValidationErrors(err))
	}
	if err = h.service.Update(c.Context(), id, request); err != nil {
		return responses.HandleError(c, err)
	}
	return responses.NewSuccessResponse(c, fiber.StatusOK, "Inbound order item updated successfully", nil)
}
func (h *InboundOrderItemHandler) Delete(c fiber.Ctx) error {
	id, err := masterParamID(c, "item_id")
	if err != nil {
		return responses.HandleError(c, err)
	}
	if err = h.service.Delete(c.Context(), id); err != nil {
		return responses.HandleError(c, err)
	}
	return responses.NewSuccessResponse(c, fiber.StatusOK, "Inbound order item deleted successfully", nil)
}
