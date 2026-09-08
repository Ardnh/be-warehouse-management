package dto

import "github.com/Ardnh/be-warehouse-management/internal/domain/entity"

// type PermissionDTO struct {
// 	ID          string `json:"id"`
// 	Resource    string `json:"resource"`
// 	Action      string `json:"action"`
// 	Description string `json:"description"`
// }

type CreatePermissionRequest struct {
	Resource    string `json:"resource" validate:"required"`
	Action      string `json:"action" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type UpdatePermissionRequest struct {
	Resource    string `json:"resource"`
	Action      string `json:"action"`
	Description string `json:"description"`
}

type PermissionDTO struct {
	ID          string `json:"id"`
	Resource    string `json:"resource"`
	Action      string `json:"action"`
	Description string `json:"description"`
}

type PermissionResponseDTO struct {
	Group  string           `json:"group"`
	Action []*PermissionDTO `json:"action"`
}

func ToPermissionDTO(permission *entity.Permission) *PermissionDTO {
	return &PermissionDTO{
		ID:          permission.ID.String(),
		Resource:    permission.Resource,
		Action:      permission.Action,
		Description: permission.Description,
	}
}

func ToPermissionsDTO(permissions []entity.Permission) []PermissionResponseDTO {
	index := make(map[string]int, len(permissions))
	result := make([]PermissionResponseDTO, 0, len(permissions))

	for i := range permissions {
		p := &permissions[i]

		pos, ok := index[p.Resource]
		if !ok {
			pos = len(result)
			index[p.Resource] = pos
			result = append(result, PermissionResponseDTO{
				Group:  p.Resource,
				Action: make([]*PermissionDTO, 0, 4),
			})
		}

		result[pos].Action = append(result[pos].Action, ToPermissionDTO(p))
	}

	return result
}
