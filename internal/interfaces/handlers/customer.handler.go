package handlers

import (
	"strconv"

	"github.com/Ardnh/be-warehouse-management/internal/application/dto"
	"github.com/Ardnh/be-warehouse-management/internal/domain/services"
	responses "github.com/Ardnh/be-warehouse-management/internal/interfaces/response"
	validator_utils "github.com/Ardnh/be-warehouse-management/internal/utils/validator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type CustomerHandler struct {
	customerService services.CustomerService
	validator       *validator.Validate
	log             *logrus.Logger
}

func NewCustomerHandler(customerService services.CustomerService, validator *validator.Validate, log *logrus.Logger) *CustomerHandler {
	return &CustomerHandler{
		customerService: customerService,
		validator:       validator,
		log:             log,
	}
}

func (h *CustomerHandler) FindAll(c fiber.Ctx) error {

	page := c.Query("page", "1")
	pageSize := c.Query("page_size", "30")
	search := c.Query("search", "")
	sortBy := c.Query("sort_by", "created_at")
	sortDir := c.Query("sort_dir", "asc")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return responses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, err)
	}

	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil {
		return responses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, err)
	}

	if pageInt <= 0 {
		pageInt = 1
	}
	if pageSizeInt <= 0 || pageSizeInt >= 1000 {
		pageSizeInt = 30
	}

	filterDto := dto.FilterDTO{
		Page:    pageInt,
		Size:    pageSizeInt,
		Search:  search,
		SortBy:  sortBy,
		SortDir: sortDir,
	}

	result, total, err := h.customerService.FindAll(c.Context(), filterDto)
	if err != nil {
		return responses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}

	pagination := dto.Pagination{
		CurrentPage: pageInt,
		PageSize:    pageSizeInt,
		TotalItems:  int(total),
		TotalPages:  (int(total) + pageSizeInt - 1) / pageSizeInt,
		HasNext:     pageInt*pageSizeInt < int(total),
		HasPrevious: pageInt > 1,
	}

	return responses.NewSuccessResponseWithPagination(c, fiber.StatusOK, "Schedules retrieved successfully", result, pagination)

}

func (h *CustomerHandler) FindById(c fiber.Ctx) error {

	id := c.Params("id", "")
	if id == "" {
		return responses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, nil)
	}

	idUUID, err := uuid.Parse(id)
	if err != nil {
		return responses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, err)
	}

	result, err := h.customerService.FindByID(c.Context(), idUUID)
	if err != nil {
		return responses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}

	return responses.NewSuccessResponse(c, fiber.StatusOK, "Customer retrieved successfully", result)
}

func (h *CustomerHandler) Create(c fiber.Ctx) error {
	var request dto.CreateCustomerRequest
	if err := c.Bind().Body(&request); err != nil {
		return responses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, err)
	}

	if err := h.validator.Struct(&request); err != nil {
		return responses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, validator_utils.FormatValidationErrors(err))
	}

	err := h.customerService.Create(c.Context(), request)
	if err != nil {
		return responses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}

	return responses.NewSuccessResponse(c, fiber.StatusCreated, "Customer created successfully", nil)
}

func (h *CustomerHandler) Update(c fiber.Ctx) error {
	id := c.Params("id", "")
	if id == "" {
		return responses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, nil)
	}

	idUUID, err := uuid.Parse(id)
	if err != nil {
		return responses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, err)
	}

	var request dto.UpdateCustomerRequest
	if err := c.Bind().Body(&request); err != nil {
		return responses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, err)
	}

	if err := h.validator.Struct(&request); err != nil {
		return responses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, validator_utils.FormatValidationErrors(err))
	}

	err = h.customerService.Update(c.Context(), idUUID, request)
	if err != nil {
		return responses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}

	return responses.NewSuccessResponse(c, fiber.StatusOK, "Customer updated successfully", nil)
}

func (h *CustomerHandler) Delete(c fiber.Ctx) error {
	id := c.Params("id", "")
	if id == "" {
		return responses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, nil)
	}

	idUUID, err := uuid.Parse(id)
	if err != nil {
		return responses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, err)
	}

	err = h.customerService.Delete(c.Context(), idUUID)
	if err != nil {
		return responses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}

	return responses.NewSuccessResponse(c, fiber.StatusOK, "Customer deleted successfully", nil)
}
