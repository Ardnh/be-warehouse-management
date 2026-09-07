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

type StorageLocationHandler struct {
	service   services.StorageLocationService
	validator *validator.Validate
	log       *logrus.Logger
}

func NewStorageLocationHandler(service services.StorageLocationService, validator *validator.Validate, log *logrus.Logger) *StorageLocationHandler {
	return &StorageLocationHandler{service: service, validator: validator, log: log}
}
func (h *StorageLocationHandler) FindAll(c fiber.Ctx) error {
	filter, err := masterFilter(c)
	if err != nil {
		return responses.HandleError(c, err)
	}
	result, total, err := h.service.FindAll(c.Context(), filter)
	if err != nil {
		return responses.HandleError(c, err)
	}
	return responses.NewSuccessResponseWithPagination(c, fiber.StatusOK, "Storage locations retrieved successfully", result, dto.NewPagination(filter, total))
}
func (h *StorageLocationHandler) FindByID(c fiber.Ctx) error {
	id, err := masterID(c)
	if err != nil {
		return responses.HandleError(c, err)
	}
	result, err := h.service.FindByID(c.Context(), id)
	if err != nil {
		return responses.HandleError(c, err)
	}
	return responses.NewSuccessResponse(c, fiber.StatusOK, "Storage location retrieved successfully", result)
}
func (h *StorageLocationHandler) Create(c fiber.Ctx) error {
	var request dto.CreateStorageLocationRequest
	if err := c.Bind().Body(&request); err != nil {
		return responses.HandleError(c, fiber.ErrBadRequest)
	}
	if err := h.validator.Struct(&request); err != nil {
		return responses.NewErrorResponse(c, fiber.StatusBadRequest, fiber.ErrBadRequest.Message, validator_utils.FormatValidationErrors(err))
	}
	if err := h.service.Create(c.Context(), request); err != nil {
		return responses.HandleError(c, err)
	}
	return responses.NewSuccessResponse(c, fiber.StatusCreated, "Storage location created successfully", nil)
}
func (h *StorageLocationHandler) Update(c fiber.Ctx) error {
	id, err := masterID(c)
	if err != nil {
		return responses.HandleError(c, err)
	}
	var request dto.UpdateStorageLocationRequest
	if err := c.Bind().Body(&request); err != nil {
		return responses.HandleError(c, fiber.ErrBadRequest)
	}
	if err := h.validator.Struct(&request); err != nil {
		return responses.NewErrorResponse(c, fiber.StatusBadRequest, fiber.ErrBadRequest.Message, validator_utils.FormatValidationErrors(err))
	}
	if err := h.service.Update(c.Context(), id, request); err != nil {
		return responses.HandleError(c, err)
	}
	return responses.NewSuccessResponse(c, fiber.StatusOK, "Storage location updated successfully", nil)
}
func (h *StorageLocationHandler) Delete(c fiber.Ctx) error {
	id, err := masterID(c)
	if err != nil {
		return responses.HandleError(c, err)
	}
	if err := h.service.Delete(c.Context(), id); err != nil {
		return responses.HandleError(c, err)
	}
	return responses.NewSuccessResponse(c, fiber.StatusOK, "Storage location deleted successfully", nil)
}
