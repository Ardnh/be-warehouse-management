package middleware

import (
	"strings"

	"github.com/Ardnh/be-warehouse-management/internal/config"
	http "github.com/Ardnh/be-warehouse-management/internal/interfaces/response"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

type AuthMiddleware struct {
	log       *logrus.Logger
	secretKey []byte // lihat catatan di bawah
}

func NewAuthMiddleware(log *logrus.Logger, cfg *config.Config) *AuthMiddleware {
	return &AuthMiddleware{
		log:       log,
		secretKey: []byte(cfg.App.JWTSecret),
	}
}

func (am *AuthMiddleware) Authenticate() fiber.Handler {
	return func(c fiber.Ctx) error {
		fields := logrus.Fields{
			"middleware": "Authenticate",
			"method":     c.Method(),
			"path":       c.Path(),
			"ip":         c.IP(),
			"user_agent": c.Get("User-Agent"),
		}

		authHeader := c.Get("Authorization")
		if authHeader == "" {
			am.log.WithFields(fields).
				Warn("akses ditolak: header Authorization kosong")

			return http.NewErrorResponse(
				c,
				fiber.StatusUnauthorized,
				"Unauthorized: No token provided",
				nil,
			)
		}

		parts := strings.Fields(authHeader)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			am.log.WithFields(fields).
				Warn("akses ditolak: format Authorization tidak valid")

			return http.NewErrorResponse(
				c,
				fiber.StatusUnauthorized,
				"Unauthorized: Invalid token format",
				nil,
			)
		}

		tokenString := parts[1]

		if len(am.secretKey) == 0 {
			am.log.WithFields(fields).
				Error("JWT secret belum dikonfigurasi")

			return http.NewErrorResponse(
				c,
				fiber.StatusInternalServerError,
				"JWT secret not configured",
				nil,
			)
		}

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				am.log.WithFields(fields).
					WithField("alg", t.Header["alg"]).
					Warn("signing method tidak sesuai")

				return nil, fiber.NewError(
					fiber.StatusUnauthorized,
					"Invalid signing method",
				)
			}

			return am.secretKey, nil
		})

		if err != nil {
			am.log.WithFields(fields).
				WithError(err).
				Warn("token tidak valid atau kedaluwarsa")

			return http.NewErrorResponse(
				c,
				fiber.StatusUnauthorized,
				"Invalid or expired token",
				err.Error(),
			)
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			am.log.WithFields(fields).
				Warn("claims token tidak valid")

			return http.NewErrorResponse(
				c,
				fiber.StatusUnauthorized,
				"Invalid token claims",
				nil,
			)
		}

		userID, ok := claims["user_id"].(string)
		if !ok || userID == "" {
			am.log.WithFields(fields).
				Warn("claim user_id kosong")

			return http.NewErrorResponse(
				c,
				fiber.StatusUnauthorized,
				"Invalid token claims",
				nil,
			)
		}

		fields["user_id"] = userID
		c.Locals("user_id", userID)
		am.log.WithFields(fields).Debug("autentikasi berhasil")

		return c.Next()
	}
}
