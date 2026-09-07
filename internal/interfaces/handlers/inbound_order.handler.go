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

type InboundOrderHandler struct {
	service   services.InboundOrderService
	validator *validator.Validate
	log       *logrus.Logger
}

func NewInboundOrderHandler(service services.InboundOrderService, validator *validator.Validate, log *logrus.Logger) *InboundOrderHandler {
	return &InboundOrderHandler{service: service, validator: validator, log: log}
}
func (h *InboundOrderHandler) FindAll(c fiber.Ctx) error {
	filter, err := masterFilter(c)
	if err != nil {
		return responses.HandleError(c, err)
	}
	result, total, err := h.service.FindAll(c.Context(), filter)
	if err != nil {
		return responses.HandleError(c, err)
	}
	return responses.NewSuccessResponseWithPagination(c, fiber.StatusOK, "Inbound orders retrieved successfully", result, dto.NewPagination(filter, total))
}
func (h *InboundOrderHandler) FindByID(c fiber.Ctx) error {
	id, err := masterID(c)
	if err != nil {
		return responses.HandleError(c, err)
	}
	result, err := h.service.FindByID(c.Context(), id)
	if err != nil {
		return responses.HandleError(c, err)
	}
	return responses.NewSuccessResponse(c, fiber.StatusOK, "Inbound order retrieved successfully", result)
}
func (h *InboundOrderHandler) Create(c fiber.Ctx) error {
	var request dto.CreateInboundOrderRequest
	if err := c.Bind().Body(&request); err != nil {
		return responses.HandleError(c, fiber.ErrBadRequest)
	}
	if err := h.validator.Struct(&request); err != nil {
		return responses.NewErrorResponse(c, fiber.StatusBadRequest, fiber.ErrBadRequest.Message, validator_utils.FormatValidationErrors(err))
	}
	if err := h.service.Create(c.Context(), request); err != nil {
		return responses.HandleError(c, err)
	}
	return responses.NewSuccessResponse(c, fiber.StatusCreated, "Inbound order created successfully", nil)
}
func (h *InboundOrderHandler) Update(c fiber.Ctx) error {
	id, err := masterID(c)
	if err != nil {
		return responses.HandleError(c, err)
	}
	var request dto.UpdateInboundOrderRequest
	if err = c.Bind().Body(&request); err != nil {
		return responses.HandleError(c, fiber.ErrBadRequest)
	}
	if err = h.validator.Struct(&request); err != nil {
		return responses.NewErrorResponse(c, fiber.StatusBadRequest, fiber.ErrBadRequest.Message, validator_utils.FormatValidationErrors(err))
	}
	if err = h.service.Update(c.Context(), id, request); err != nil {
		return responses.HandleError(c, err)
	}
	return responses.NewSuccessResponse(c, fiber.StatusOK, "Inbound order updated successfully", nil)
}
func (h *InboundOrderHandler) UpdateStatus(c fiber.Ctx) error {
	id, err := masterID(c)
	if err != nil {
		return responses.HandleError(c, err)
	}
	var request dto.UpdateInboundOrderStatusRequest
	if err = c.Bind().Body(&request); err != nil {
		return responses.HandleError(c, fiber.ErrBadRequest)
	}
	if err = h.validator.Struct(&request); err != nil {
		return responses.NewErrorResponse(c, fiber.StatusBadRequest, fiber.ErrBadRequest.Message, validator_utils.FormatValidationErrors(err))
	}
	if err = h.service.UpdateStatus(c.Context(), id, request); err != nil {
		return responses.HandleError(c, err)
	}
	return responses.NewSuccessResponse(c, fiber.StatusOK, "Inbound order status updated successfully", nil)
}
func (h *InboundOrderHandler) Delete(c fiber.Ctx) error {
	id, err := masterID(c)
	if err != nil {
		return responses.HandleError(c, err)
	}
	if err = h.service.Delete(c.Context(), id); err != nil {
		return responses.HandleError(c, err)
	}
	return responses.NewSuccessResponse(c, fiber.StatusOK, "Inbound order deleted successfully", nil)
}
