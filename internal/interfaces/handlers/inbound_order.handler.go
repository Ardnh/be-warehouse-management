package handlers

import (
	"github.com/Ardnh/be-warehouse-management/internal/application/dto"
	"github.com/Ardnh/be-warehouse-management/internal/domain/services"
	responses "github.com/Ardnh/be-warehouse-management/internal/interfaces/response"
	validator_utils "github.com/Ardnh/be-warehouse-management/internal/utils/validator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
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

func (h *InboundOrderHandler) requestLog(c fiber.Ctx, method string) *logrus.Entry {
	return h.log.WithFields(logrus.Fields{
		"component":   "inbound_order_handler",
		"method":      method,
		"http_method": c.Method(),
		"path":        c.Path(),
		"request_id":  c.Get("X-Request-ID"),
	})
}

func (h *InboundOrderHandler) FindAll(c fiber.Ctx) error {
	entry := h.requestLog(c, "FindAll")
	filter, err := masterFilter(c)
	if err != nil {
		entry.WithError(err).Warn("invalid inbound order list filter")
		return responses.HandleError(c, err)
	}
	result, total, err := h.service.FindAll(c.Context(), filter)
	if err != nil {
		entry.WithError(err).Error("failed to retrieve inbound orders")
		return responses.HandleError(c, err)
	}
	entry.WithFields(logrus.Fields{"total": total, "page": filter.Page, "page_size": filter.Size}).Debug("inbound orders retrieved")
	return responses.NewSuccessResponseWithPagination(c, fiber.StatusOK, "Inbound orders retrieved successfully", result, dto.NewPagination(filter, total))
}

func (h *InboundOrderHandler) FindByID(c fiber.Ctx) error {
	entry := h.requestLog(c, "FindByID")
	id, err := masterID(c)
	if err != nil {
		entry.WithError(err).Warn("invalid inbound order ID")
		return responses.HandleError(c, err)
	}
	entry = entry.WithField("inbound_order_id", id.String())
	result, err := h.service.FindByID(c.Context(), id)
	if err != nil {
		entry.WithError(err).Error("failed to retrieve inbound order")
		return responses.HandleError(c, err)
	}
	entry.Debug("inbound order retrieved")
	return responses.NewSuccessResponse(c, fiber.StatusOK, "Inbound order retrieved successfully", result)
}

func (h *InboundOrderHandler) Create(c fiber.Ctx) error {
	entry := h.requestLog(c, "Create")
	var request dto.CreateInboundOrderRequest
	warehouseID, ok := c.Locals("warehouse_id").(string)
	if !ok || warehouseID == "" {
		entry.Warn("warehouse context missing for inbound order creation")
		return responses.HandleError(c, fiber.ErrBadRequest)
	}

	if err := c.Bind().Body(&request); err != nil {
		entry.WithError(err).Warn("failed to parse inbound order request")
		return responses.HandleError(c, fiber.ErrBadRequest)
	}
	if err := h.validator.Struct(&request); err != nil {
		entry.WithError(err).Warn("inbound order request validation failed")
		return responses.NewErrorResponse(c, fiber.StatusBadRequest, fiber.ErrBadRequest.Message, validator_utils.FormatValidationErrors(err))
	}

	warehouseIDUUID, err := uuid.Parse(warehouseID)
	if err != nil {
		entry.WithError(err).Warn("invalid warehouse ID in request context")
		return responses.HandleError(c, err)
	}
	entry = entry.WithFields(logrus.Fields{
		"warehouse_id": warehouseIDUUID.String(),
		"customer_id":  request.CustomerID.String(),
		"item_count":   len(request.Items),
	})
	if err := h.service.Create(c.Context(), warehouseIDUUID, request); err != nil {
		entry.WithError(err).Error("failed to create inbound order")
		return responses.HandleError(c, err)
	}
	entry.Info("inbound order created")
	return responses.NewSuccessResponse(c, fiber.StatusCreated, "Inbound order created successfully", nil)
}

func (h *InboundOrderHandler) Update(c fiber.Ctx) error {
	entry := h.requestLog(c, "Update")
	id, err := masterID(c)
	if err != nil {
		entry.WithError(err).Warn("invalid inbound order ID")
		return responses.HandleError(c, err)
	}
	entry = entry.WithField("inbound_order_id", id.String())
	var request dto.UpdateInboundOrderRequest
	if err = c.Bind().Body(&request); err != nil {
		entry.WithError(err).Warn("failed to parse inbound order update request")
		return responses.HandleError(c, fiber.ErrBadRequest)
	}
	if err = h.validator.Struct(&request); err != nil {
		entry.WithError(err).Warn("inbound order update validation failed")
		return responses.NewErrorResponse(c, fiber.StatusBadRequest, fiber.ErrBadRequest.Message, validator_utils.FormatValidationErrors(err))
	}
	if err = h.service.Update(c.Context(), id, request); err != nil {
		entry.WithError(err).Error("failed to update inbound order")
		return responses.HandleError(c, err)
	}
	entry.Info("inbound order updated")
	return responses.NewSuccessResponse(c, fiber.StatusOK, "Inbound order updated successfully", nil)
}

func (h *InboundOrderHandler) UpdateStatus(c fiber.Ctx) error {
	entry := h.requestLog(c, "UpdateStatus")
	id, err := masterID(c)
	if err != nil {
		entry.WithError(err).Warn("invalid inbound order ID")
		return responses.HandleError(c, err)
	}
	entry = entry.WithField("inbound_order_id", id.String())
	var request dto.UpdateInboundOrderStatusRequest
	if err = c.Bind().Body(&request); err != nil {
		entry.WithError(err).Warn("failed to parse inbound order status request")
		return responses.HandleError(c, fiber.ErrBadRequest)
	}
	if err = h.validator.Struct(&request); err != nil {
		entry.WithError(err).Warn("inbound order status validation failed")
		return responses.NewErrorResponse(c, fiber.StatusBadRequest, fiber.ErrBadRequest.Message, validator_utils.FormatValidationErrors(err))
	}
	if err = h.service.UpdateStatus(c.Context(), id, request); err != nil {
		entry.WithFields(logrus.Fields{"status": request.Status}).WithError(err).Error("failed to update inbound order status")
		return responses.HandleError(c, err)
	}
	entry.WithField("status", request.Status).Info("inbound order status updated")
	return responses.NewSuccessResponse(c, fiber.StatusOK, "Inbound order status updated successfully", nil)
}

func (h *InboundOrderHandler) Delete(c fiber.Ctx) error {
	entry := h.requestLog(c, "Delete")
	id, err := masterID(c)
	if err != nil {
		entry.WithError(err).Warn("invalid inbound order ID")
		return responses.HandleError(c, err)
	}
	entry = entry.WithField("inbound_order_id", id.String())
	if err = h.service.Delete(c.Context(), id); err != nil {
		entry.WithError(err).Error("failed to delete inbound order")
		return responses.HandleError(c, err)
	}
	entry.Info("inbound order deleted")
	return responses.NewSuccessResponse(c, fiber.StatusOK, "Inbound order deleted successfully", nil)
}
