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

type ProductHandler struct {
	service   services.ProductService
	validator *validator.Validate
	log       *logrus.Logger
}

func NewProductHandler(s services.ProductService, v *validator.Validate, l *logrus.Logger) *ProductHandler {
	return &ProductHandler{service: s, validator: v, log: l}
}
func (h *ProductHandler) FindAll(c fiber.Ctx) error {
	f, e := masterFilter(c)
	if e != nil {
		return responses.HandleError(c, e)
	}
	x, t, e := h.service.FindAll(c.Context(), f)
	if e != nil {
		return responses.HandleError(c, e)
	}
	return responses.NewSuccessResponseWithPagination(c, 200, "Products retrieved successfully", x, dto.NewPagination(f, t))
}
func (h *ProductHandler) FindByID(c fiber.Ctx) error {
	id, e := masterID(c)
	if e != nil {
		return responses.HandleError(c, e)
	}
	x, e := h.service.FindByID(c.Context(), id)
	if e != nil {
		return responses.HandleError(c, e)
	}
	return responses.NewSuccessResponse(c, 200, "Product retrieved successfully", x)
}
func (h *ProductHandler) Create(c fiber.Ctx) error {
	var r dto.CreateProductRequest
	if e := c.Bind().Body(&r); e != nil {
		return responses.HandleError(c, fiber.ErrBadRequest)
	}
	if e := h.validator.Struct(&r); e != nil {
		return responses.NewErrorResponse(c, 400, fiber.ErrBadRequest.Message, vu.FormatValidationErrors(e))
	}
	if e := h.service.Create(c.Context(), r); e != nil {
		return responses.HandleError(c, e)
	}
	return responses.NewSuccessResponse(c, 201, "Product created successfully", nil)
}
func (h *ProductHandler) Update(c fiber.Ctx) error {
	id, e := masterID(c)
	if e != nil {
		return responses.HandleError(c, e)
	}
	var r dto.UpdateProductRequest
	if e = c.Bind().Body(&r); e != nil {
		return responses.HandleError(c, fiber.ErrBadRequest)
	}
	if e = h.validator.Struct(&r); e != nil {
		return responses.NewErrorResponse(c, 400, fiber.ErrBadRequest.Message, vu.FormatValidationErrors(e))
	}
	if e = h.service.Update(c.Context(), id, r); e != nil {
		return responses.HandleError(c, e)
	}
	return responses.NewSuccessResponse(c, 200, "Product updated successfully", nil)
}
func (h *ProductHandler) Delete(c fiber.Ctx) error {
	id, e := masterID(c)
	if e != nil {
		return responses.HandleError(c, e)
	}
	if e = h.service.Delete(c.Context(), id); e != nil {
		return responses.HandleError(c, e)
	}
	return responses.NewSuccessResponse(c, 200, "Product deleted successfully", nil)
}
