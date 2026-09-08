package handlers

import (
	"github.com/Ardnh/be-warehouse-management/internal/application/dto"
	"github.com/Ardnh/be-warehouse-management/internal/domain/services"
	responses "github.com/Ardnh/be-warehouse-management/internal/interfaces/response"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type PermissionHandler struct {
	service   services.PermissionService
	validator *validator.Validate
	log       *logrus.Logger
}

func NewPermissionHandler(services services.PermissionService, validator *validator.Validate, log *logrus.Logger) *PermissionHandler {
	return &PermissionHandler{
		service:   services,
		validator: validator,
		log:       log,
	}
}

func (h *PermissionHandler) FindAll(c fiber.Ctx) error {
	f, e := masterFilter(c)
	if e != nil {
		return responses.HandleError(c, e)
	}
	x, t, e := h.service.FindAll(c.Context(), f)
	if e != nil {
		return responses.HandleError(c, e)
	}
	return responses.NewSuccessResponseWithPagination(c, 200, "Permissions retrieved successfully", x, dto.NewPagination(f, t))
}

func (h *PermissionHandler) FindByID(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return responses.HandleError(c, fiber.NewError(fiber.StatusBadRequest, "id is required"))
	}

	idUuid, err := uuid.Parse(id)
	if err != nil {
		return responses.HandleError(c, fiber.NewError(fiber.StatusBadRequest, "invalid id"))
	}

	x, e := h.service.FindByID(c.Context(), idUuid)
	if e != nil {
		return responses.HandleError(c, e)
	}
	return responses.NewSuccessResponse(c, 200, "Permission retrieved successfully", x)
}

func (h *PermissionHandler) Create(c fiber.Ctx) error {
	var req dto.CreatePermissionRequest
	if e := c.Bind().Body(&req); e != nil {
		return responses.HandleError(c, e)
	}
	if e := h.validator.Struct(&req); e != nil {
		return responses.HandleError(c, e)
	}
	if e := h.service.Create(c.Context(), req); e != nil {
		return responses.HandleError(c, e)
	}
	return responses.NewSuccessResponse(c, 201, "Permission created successfully", nil)
}

func (h *PermissionHandler) Update(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return responses.HandleError(c, fiber.NewError(fiber.StatusBadRequest, "id is required"))
	}

	idUuid, err := uuid.Parse(id)
	if err != nil {
		return responses.HandleError(c, fiber.NewError(fiber.StatusBadRequest, "invalid id"))
	}
	var req dto.UpdatePermissionRequest
	if e := c.Bind().Body(&req); e != nil {
		return responses.HandleError(c, e)
	}
	if e := h.validator.Struct(&req); e != nil {
		return responses.HandleError(c, e)
	}
	if e := h.service.Update(c.Context(), idUuid, req); e != nil {
		return responses.HandleError(c, e)
	}
	return responses.NewSuccessResponse(c, 200, "Permission updated successfully", nil)
}

func (h *PermissionHandler) Delete(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return responses.HandleError(c, fiber.NewError(fiber.StatusBadRequest, "id is required"))
	}

	idUuid, err := uuid.Parse(id)
	if err != nil {
		return responses.HandleError(c, fiber.NewError(fiber.StatusBadRequest, "invalid id"))
	}
	if e := h.service.Delete(c.Context(), idUuid); e != nil {
		return responses.HandleError(c, e)
	}
	return responses.NewSuccessResponse(c, 200, "Permission deleted successfully", nil)
}
