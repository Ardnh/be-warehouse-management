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

type InboundOrderItemHandler struct {
	service   services.InboundOrderItemService
	validator *validator.Validate
	log       *logrus.Logger
}

func NewInboundOrderItemHandler(service services.InboundOrderItemService, validator *validator.Validate, log *logrus.Logger) *InboundOrderItemHandler {
	return &InboundOrderItemHandler{service: service, validator: validator, log: log}
}

func (h *InboundOrderItemHandler) requestLog(c fiber.Ctx, method string) *logrus.Entry {
	return h.log.WithFields(logrus.Fields{
		"component":   "inbound_order_item_handler",
		"method":      method,
		"http_method": c.Method(),
		"path":        c.Path(),
		"request_id":  c.Get("X-Request-ID"),
	})
}

func (h *InboundOrderItemHandler) FindAll(c fiber.Ctx) error {
	entry := h.requestLog(c, "FindAll")
	orderID, err := masterID(c)
	if err != nil {
		entry.WithError(err).Warn("invalid inbound order ID")
		return responses.HandleError(c, err)
	}
	entry = entry.WithField("inbound_order_id", orderID.String())
	result, err := h.service.FindAll(c.Context(), orderID)
	if err != nil {
		entry.WithError(err).Error("failed to retrieve inbound order items")
		return responses.HandleError(c, err)
	}
	entry.WithField("item_count", len(result)).Debug("inbound order items retrieved")
	return responses.NewSuccessResponse(c, fiber.StatusOK, "Inbound order items retrieved successfully", result)
}

func (h *InboundOrderItemHandler) FindByID(c fiber.Ctx) error {
	entry := h.requestLog(c, "FindByID")
	id, err := masterParamID(c, "item_id")
	if err != nil {
		entry.WithError(err).Warn("invalid inbound order item ID")
		return responses.HandleError(c, err)
	}
	entry = entry.WithField("inbound_order_item_id", id.String())
	result, err := h.service.FindByID(c.Context(), id)
	if err != nil {
		entry.WithError(err).Error("failed to retrieve inbound order item")
		return responses.HandleError(c, err)
	}
	entry.Debug("inbound order item retrieved")
	return responses.NewSuccessResponse(c, fiber.StatusOK, "Inbound order item retrieved successfully", result)
}

func (h *InboundOrderItemHandler) Create(c fiber.Ctx) error {
	entry := h.requestLog(c, "Create")
	orderID, err := masterID(c)
	if err != nil {
		entry.WithError(err).Warn("invalid inbound order ID")
		return responses.HandleError(c, err)
	}
	entry = entry.WithField("inbound_order_id", orderID.String())
	var request dto.CreateInboundOrderItemRequest
	if err = c.Bind().Body(&request); err != nil {
		entry.WithError(err).Warn("failed to parse inbound order item request")
		return responses.HandleError(c, fiber.ErrBadRequest)
	}
	if err = h.validator.Struct(&request); err != nil {
		entry.WithError(err).Warn("inbound order item request validation failed")
		return responses.NewErrorResponse(c, fiber.StatusBadRequest, fiber.ErrBadRequest.Message, validator_utils.FormatValidationErrors(err))
	}
	entry = entry.WithField("product_id", request.ProductID.String())
	if err = h.service.Create(c.Context(), orderID, request); err != nil {
		entry.WithError(err).Error("failed to create inbound order item")
		return responses.HandleError(c, err)
	}
	entry.Info("inbound order item created")
	return responses.NewSuccessResponse(c, fiber.StatusCreated, "Inbound order item created successfully", nil)
}

func (h *InboundOrderItemHandler) Update(c fiber.Ctx) error {
	entry := h.requestLog(c, "Update")
	id, err := masterParamID(c, "item_id")
	if err != nil {
		entry.WithError(err).Warn("invalid inbound order item ID")
		return responses.HandleError(c, err)
	}
	entry = entry.WithField("inbound_order_item_id", id.String())
	var request dto.UpdateInboundOrderItemRequest
	if err = c.Bind().Body(&request); err != nil {
		entry.WithError(err).Warn("failed to parse inbound order item update request")
		return responses.HandleError(c, fiber.ErrBadRequest)
	}
	if err = h.validator.Struct(&request); err != nil {
		entry.WithError(err).Warn("inbound order item update validation failed")
		return responses.NewErrorResponse(c, fiber.StatusBadRequest, fiber.ErrBadRequest.Message, validator_utils.FormatValidationErrors(err))
	}
	entry = entry.WithField("product_id", request.ProductID.String())
	if err = h.service.Update(c.Context(), id, request); err != nil {
		entry.WithError(err).Error("failed to update inbound order item")
		return responses.HandleError(c, err)
	}
	entry.Info("inbound order item updated")
	return responses.NewSuccessResponse(c, fiber.StatusOK, "Inbound order item updated successfully", nil)
}

func (h *InboundOrderItemHandler) Delete(c fiber.Ctx) error {
	entry := h.requestLog(c, "Delete")
	id, err := masterParamID(c, "item_id")
	if err != nil {
		entry.WithError(err).Warn("invalid inbound order item ID")
		return responses.HandleError(c, err)
	}
	entry = entry.WithField("inbound_order_item_id", id.String())
	if err = h.service.Delete(c.Context(), id); err != nil {
		entry.WithError(err).Error("failed to delete inbound order item")
		return responses.HandleError(c, err)
	}
	entry.Info("inbound order item deleted")
	return responses.NewSuccessResponse(c, fiber.StatusOK, "Inbound order item deleted successfully", nil)
}
