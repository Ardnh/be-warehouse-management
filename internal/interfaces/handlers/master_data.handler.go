package handlers

import (
	"strconv"

	"github.com/Ardnh/be-warehouse-management/internal/application/dto"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func masterFilter(c fiber.Ctx) (dto.FilterDTO, error) {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return dto.FilterDTO{}, fiber.ErrBadRequest
	}
	size, err := strconv.Atoi(c.Query("page_size", "30"))
	if err != nil {
		return dto.FilterDTO{}, fiber.ErrBadRequest
	}
	if page <= 0 {
		page = 1
	}
	if size <= 0 || size >= 1000 {
		size = 30
	}
	return dto.FilterDTO{
		Page:    page,
		Size:    size,
		Search:  c.Query("search"),
		SortBy:  c.Query("sort_by", "resource"),
		SortDir: c.Query("sort_dir", "asc"),
	}, nil
}

func masterID(c fiber.Ctx) (uuid.UUID, error) { return masterParamID(c, "id") }

func masterParamID(c fiber.Ctx, name string) (uuid.UUID, error) {
	id := c.Params(name)
	if id == "" {
		return uuid.Nil, fiber.ErrBadRequest
	}
	parsed, err := uuid.Parse(id)
	if err != nil {
		return uuid.Nil, fiber.ErrBadRequest
	}
	return parsed, nil
}
