package dto

import (
	"time"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"github.com/google/uuid"
)

// dto/rack.go
type CreateRackRequest struct {
	ZoneID         uuid.UUID `json:"zone_id" validate:"required"`
	Code           string    `json:"code" validate:"required,max=50"`
	BayCount       int       `json:"bay_count" validate:"required,min=1"`
	LevelCount     int       `json:"level_count" validate:"required,min=1"`
	PalletCapacity int       `json:"pallet_capacity" validate:"min=0"`
	Status         string    `json:"status" validate:"omitempty,oneof=ACTIVE INACTIVE"`
}

type UpdateRackRequest struct {
	BayCount       *int    `json:"bay_count" validate:"omitempty,min=1"`
	LevelCount     *int    `json:"level_count" validate:"omitempty,min=1"`
	PalletCapacity *int    `json:"pallet_capacity" validate:"omitempty,min=0"`
	Status         *string `json:"status" validate:"omitempty,oneof=ACTIVE INACTIVE"`
}

type RackResponse struct {
	ID             uuid.UUID `json:"id"`
	ZoneID         uuid.UUID `json:"zone_id"`
	ZoneCode       string    `json:"zone_code,omitempty"`
	Code           string    `json:"code"`
	BayCount       int       `json:"bay_count"`
	LevelCount     int       `json:"level_count"`
	PalletCapacity int       `json:"pallet_capacity"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func NewRackResponse(r entity.Rack) RackResponse {
	res := RackResponse{
		ID:             r.ID,
		ZoneID:         r.ZoneID,
		Code:           r.Code,
		BayCount:       r.BayCount,
		LevelCount:     r.LevelCount,
		PalletCapacity: r.PalletCapacity,
		Status:         r.Status,
		CreatedAt:      r.CreatedAt,
		UpdatedAt:      r.UpdatedAt,
	}
	if r.Zone != nil {
		res.ZoneCode = r.Zone.Code
	}
	return res
}
