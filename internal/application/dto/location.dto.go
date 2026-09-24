package dto

import (
	"time"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"github.com/google/uuid"
)

type CreateLocationRequest struct {
	Code       string              `json:"code" validate:"required,max=30"`
	Name       string              `json:"name" validate:"required,max=100"`
	Type       entity.LocationType `json:"type" validate:"required,oneof=HO WAREHOUSE"`
	Address    *string             `json:"address"`
	City       *string             `json:"city" validate:"omitempty,max=100"`
	Province   *string             `json:"province" validate:"omitempty,max=100"`
	PostalCode *string             `json:"postal_code" validate:"omitempty,max=10"`
	IsActive   *bool               `json:"is_active"`
}

type UpdateLocationRequest struct {
	Name       *string              `json:"name" validate:"omitempty,max=100"`
	Type       *entity.LocationType `json:"type" validate:"omitempty,oneof=HO WAREHOUSE"`
	Address    *string              `json:"address"`
	City       *string              `json:"city" validate:"omitempty,max=100"`
	Province   *string              `json:"province" validate:"omitempty,max=100"`
	PostalCode *string              `json:"postal_code" validate:"omitempty,max=10"`
	IsActive   *bool                `json:"is_active"`
}

type LocationResponse struct {
	ID       uuid.UUID           `json:"id"`
	Code     string              `json:"code"`
	Name     string              `json:"name"`
	Type     entity.LocationType `json:"type"`
	Address  *string             `json:"address,omitempty"`
	City     *string             `json:"city,omitempty"`
	Province *string             `json:"province,omitempty"`
	IsActive bool                `json:"is_active"`
}

func ToLocationResponse(location *entity.Location) *LocationResponse {
	if location == nil {
		return nil
	}
	return &LocationResponse{
		ID:       location.ID,
		Code:     location.Code,
		Name:     location.Name,
		Type:     location.Type,
		Address:  location.Address,
		City:     location.City,
		Province: location.Province,
		IsActive: location.IsActive,
	}
}

type LocationDTO struct {
	ID         uuid.UUID           `json:"id"`
	Code       string              `json:"code"`
	Name       string              `json:"name"`
	Type       entity.LocationType `json:"type"`
	Address    *string             `json:"address,omitempty"`
	City       *string             `json:"city,omitempty"`
	Province   *string             `json:"province,omitempty"`
	PostalCode *string             `json:"postal_code,omitempty"`
	IsActive   bool                `json:"is_active"`
	CreatedAt  time.Time           `json:"created_at"`
	UpdatedAt  time.Time           `json:"updated_at"`
}

func ToLocationDTO(location entity.Location) LocationDTO {
	return LocationDTO{
		ID:         location.ID,
		Code:       location.Code,
		Name:       location.Name,
		Type:       location.Type,
		Address:    location.Address,
		City:       location.City,
		Province:   location.Province,
		PostalCode: location.PostalCode,
		IsActive:   location.IsActive,
		CreatedAt:  location.CreatedAt,
		UpdatedAt:  location.UpdatedAt,
	}
}
