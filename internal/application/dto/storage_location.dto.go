package dto

import (
	"time"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"github.com/google/uuid"
)

type CreateStorageLocationRequest struct {
	RackID   uuid.UUID `json:"rack_id" validate:"required"`
	Code     string    `json:"code" validate:"required,max=50"`
	Bay      int       `json:"bay" validate:"required,min=1"`
	Level    int       `json:"level" validate:"required,min=1"`
	Capacity int       `json:"capacity" validate:"min=0"`
	Status   string    `json:"status" validate:"omitempty,oneof=ACTIVE INACTIVE BLOCKED"`
}

type UpdateStorageLocationRequest struct {
	Capacity *int    `json:"capacity" validate:"omitempty,min=0"`
	Status   *string `json:"status" validate:"omitempty,oneof=ACTIVE INACTIVE BLOCKED"`
}

type StorageLocationResponse struct {
	ID        uuid.UUID `json:"id"`
	RackID    uuid.UUID `json:"rack_id"`
	RackCode  string    `json:"rack_code,omitempty"`
	Code      string    `json:"code"`
	Bay       int       `json:"bay"`
	Level     int       `json:"level"`
	Capacity  int       `json:"capacity"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewStorageLocationResponse(s entity.StorageLocation) StorageLocationResponse {
	res := StorageLocationResponse{
		ID:        s.ID,
		RackID:    s.RackID,
		Code:      s.Code,
		Bay:       s.Bay,
		Level:     s.Level,
		Capacity:  s.Capacity,
		Status:    s.Status,
		CreatedAt: s.CreatedAt, UpdatedAt: s.UpdatedAt,
	}
	if s.Rack != nil {
		res.RackCode = s.Rack.Code
	}
	return res
}
