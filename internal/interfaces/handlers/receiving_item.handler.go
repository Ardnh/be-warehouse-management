package handlers

import (
	"github.com/Ardnh/be-warehouse-management/internal/application/dto"
	"github.com/Ardnh/be-warehouse-management/internal/domain/services"
	responses "github.com/Ardnh/be-warehouse-management/internal/interface/response"
	vu "github.com/Ardnh/be-warehouse-management/internal/utils/validator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
)

type ReceivingItemHandler struct {
	service   services.ReceivingItemService
	validator *validator.Validate
	log       *logrus.Logger
}

func NewReceivingItemHandler(s services.ReceivingItemService, v *validator.Validate, l *logrus.Logger) *ReceivingItemHandler {
	return &ReceivingItemHandler{service: s, validator: v, log: l}
}
func (h *ReceivingItemHandler) FindAll(c fiber.Ctx) error {
	id, e := masterID(c)
	if e != nil {
		return responses.HandleError(c, e)
	}
	x, e := h.service.FindAll(c.Context(), id)
	if e != nil {
		return responses.HandleError(c, e)
	}
	return responses.NewSuccessResponse(c, 200, "Receiving items retrieved successfully", x)
}
func (h *ReceivingItemHandler) FindByID(c fiber.Ctx) error {
	id, e := masterParamID(c, "item_id")
	if e != nil {
		return responses.HandleError(c, e)
	}
	x, e := h.service.FindByID(c.Context(), id)
	if e != nil {
		return responses.HandleError(c, e)
	}
	return responses.NewSuccessResponse(c, 200, "Receiving item retrieved successfully", x)
}
func (h *ReceivingItemHandler) Create(c fiber.Ctx) error {
	id, e := masterID(c)
	if e != nil {
		return responses.HandleError(c, e)
	}
	var r dto.CreateReceivingItemRequest
	if e = c.Bind().Body(&r); e != nil {
		return responses.HandleError(c, fiber.ErrBadRequest)
	}
	if e = h.validator.Struct(&r); e != nil {
		return responses.NewErrorResponse(c, 400, fiber.ErrBadRequest.Message, vu.FormatValidationErrors(e))
	}
	if e = h.service.Create(c.Context(), id, r); e != nil {
		return responses.HandleError(c, e)
	}
	return responses.NewSuccessResponse(c, 201, "Receiving item created successfully", nil)
}
func (h *ReceivingItemHandler) Update(c fiber.Ctx) error {
	id, e := masterParamID(c, "item_id")
	if e != nil {
		return responses.HandleError(c, e)
	}
	var r dto.UpdateReceivingItemRequest
	if e = c.Bind().Body(&r); e != nil {
		return responses.HandleError(c, fiber.ErrBadRequest)
	}
	if e = h.validator.Struct(&r); e != nil {
		return responses.NewErrorResponse(c, 400, fiber.ErrBadRequest.Message, vu.FormatValidationErrors(e))
	}
	if e = h.service.Update(c.Context(), id, r); e != nil {
		return responses.HandleError(c, e)
	}
	return responses.NewSuccessResponse(c, 200, "Receiving item updated successfully", nil)
}
func (h *ReceivingItemHandler) Delete(c fiber.Ctx) error {
	id, e := masterParamID(c, "item_id")
	if e != nil {
		return responses.HandleError(c, e)
	}
	if e = h.service.Delete(c.Context(), id); e != nil {
		return responses.HandleError(c, e)
	}
	return responses.NewSuccessResponse(c, 200, "Receiving item deleted successfully", nil)
}
