package dto

import (
	"time"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"github.com/google/uuid"
)

type LocationResponse struct {
	ID       uuid.UUID           `json:"id"`
	Code     string              `json:"code"`
	Name     string              `json:"name"`
	Type     entity.LocationType `json:"type"`
	IsActive bool                `json:"is_active"`
}

type UserAssignmentDTO struct {
	ID         uuid.UUID         `json:"id"`
	UserID     uuid.UUID         `json:"user_id"`
	LocationID uuid.UUID         `json:"location_id"`
	RoleID     uuid.UUID         `json:"role_id"`
	CreatedAt  time.Time         `json:"created_at"`
	UpdatedAt  time.Time         `json:"updated_at"`
	Location   *LocationResponse `json:"location,omitempty"`
	Role       *RoleResponse     `json:"role,omitempty"`
}

func ToUserAssignmentDTO(assignment *entity.UserAssignment) *UserAssignmentDTO {
	if assignment == nil {
		return nil
	}

	result := &UserAssignmentDTO{
		ID:         assignment.ID,
		UserID:     assignment.UserID,
		LocationID: assignment.LocationID,
		RoleID:     assignment.RoleID,
		CreatedAt:  assignment.CreatedAt,
		UpdatedAt:  assignment.UpdatedAt,
	}
	if assignment.Location != nil {
		result.Location = &LocationResponse{
			ID:       assignment.Location.ID,
			Code:     assignment.Location.Code,
			Name:     assignment.Location.Name,
			Type:     assignment.Location.Type,
			IsActive: assignment.Location.IsActive,
		}
	}
	if assignment.Role != nil {
		role := NewRoleResponse(*assignment.Role)
		result.Role = &role
	}
	return result
}

func ToUserAssignmentDTOs(assignments []entity.UserAssignment) []*UserAssignmentDTO {
	result := make([]*UserAssignmentDTO, 0, len(assignments))
	for i := range assignments {
		result = append(result, ToUserAssignmentDTO(&assignments[i]))
	}
	return result
}
