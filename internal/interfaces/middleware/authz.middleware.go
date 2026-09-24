package middleware

import (
	"fmt"
	"time"

	"github.com/Ardnh/be-warehouse-management/internal/domain/services"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type AuthorizationMiddleware struct {
	permissionService services.PermissionService
	log               *logrus.Logger
}

func NewAuthorizationMiddleware(permissionService services.PermissionService, log *logrus.Logger) *AuthorizationMiddleware {
	if log == nil {
		log = logrus.StandardLogger()
	}
	return &AuthorizationMiddleware{
		permissionService: permissionService,
		log:               log,
	}
}

func (m *AuthorizationMiddleware) Require(resource string, action string) fiber.Handler {
	return func(c fiber.Ctx) error {
		entry := m.log.WithFields(logrus.Fields{
			"component": "authorization_middleware",
			"method":    c.Method(),
			"path":      c.Path(),
			"ip":        c.IP(),
			"resource":  resource,
			"action":    action,
		})

		if reqID, ok := c.Locals("requestid").(string); ok && reqID != "" {
			entry = entry.WithField("request_id", reqID)
		}

		userID, ok := c.Locals("user_id").(string)
		if !ok {
			entry.WithField("user_id_type", fmt.Sprintf("%T", c.Locals("user_id"))).
				Warn("authorization failed: user_id not found in context")
			return fiber.ErrUnauthorized
		}

		userIdUuid, err := uuid.Parse(userID)
		if err != nil {
			entry.WithError(err).Error("authorization failed: invalid user_id")
			return fiber.ErrUnauthorized
		}
		entry = entry.WithField("user_id", userIdUuid.String())
		start := time.Now()
		allowed, err := m.permissionService.HasPermission(
			c.Context(),
			userIdUuid,
			resource,
			action,
		)
		entry = entry.WithField("duration_ms", time.Since(start).Milliseconds())

		if err != nil {
			entry.WithError(err).Error("authorization failed: error checking permission")
			return err
		}

		if !allowed {
			entry.Warn("authorization denied: insufficient permission")
			return fiber.ErrForbidden
		}

		entry.Debug("authorization granted")

		return c.Next()
	}
}
