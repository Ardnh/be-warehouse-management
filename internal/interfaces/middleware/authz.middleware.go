package middleware

import (
	"github.com/Ardnh/be-warehouse-management/internal/domain/services"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type AuthorizationMiddleware struct {
	permissionService services.PermissionService
}

func NewAuthorizationMiddleware(permissionService services.PermissionService) *AuthorizationMiddleware {
	return &AuthorizationMiddleware{
		permissionService: permissionService,
	}
}

func (m *AuthorizationMiddleware) Require(resource string, action string) fiber.Handler {
	return func(c fiber.Ctx) error {
		userID, ok := c.Locals("user_id").(uuid.UUID)

		if !ok {
			return fiber.ErrUnauthorized
		}

		allowed, err := m.permissionService.HasPermission(
			c.Context(),
			userID,
			resource,
			action,
		)

		if err != nil {
			return err
		}

		if !allowed {
			return fiber.ErrForbidden
		}

		return c.Next()
	}
}
