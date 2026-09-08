package middleware

import (
	"strings"
	"time"

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
			am.log.WithFields(fields).Warn("akses ditolak: header Authorization kosong")
			return http.NewErrorResponse(c, fiber.StatusUnauthorized, "Unauthorized: No token provided", nil)
		}

		tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer"))
		if tokenString == "" {
			am.log.WithFields(fields).Warn("akses ditolak: format token tidak valid")
			return http.NewErrorResponse(c, fiber.StatusUnauthorized, "Unauthorized: Invalid token format", nil)
		}

		if len(am.secretKey) == 0 {
			// Ini kesalahan konfigurasi server, bukan kesalahan klien.
			am.log.WithFields(fields).Error("JWT secret belum dikonfigurasi")
			return http.NewErrorResponse(c, fiber.StatusInternalServerError, "JWT secret not configured", nil)
		}

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				am.log.WithFields(fields).
					WithField("alg", t.Header["alg"]).
					Warn("signing method tidak sesuai")
				return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid signing method")
			}
			return am.secretKey, nil
		})
		if err != nil {
			am.log.WithFields(fields).WithError(err).Warn("token tidak valid atau kedaluwarsa")
			return http.NewErrorResponse(c, fiber.StatusUnauthorized, "Invalid or expired token", err.Error())
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			am.log.WithFields(fields).Warn("claims token tidak valid")
			return http.NewErrorResponse(c, fiber.StatusUnauthorized, "Invalid token claims", nil)
		}

		if exp, ok := claims["exp"].(float64); ok && time.Now().Unix() > int64(exp) {
			am.log.WithFields(fields).
				WithField("expired_at", time.Unix(int64(exp), 0)).
				Warn("token sudah kedaluwarsa")
			return http.NewErrorResponse(c, fiber.StatusUnauthorized, "Token has expired", nil)
		}

		userID, _ := claims["user_id"].(string)
		if userID == "" {
			am.log.WithFields(fields).Warn("claim user_id kosong")
			return http.NewErrorResponse(c, fiber.StatusUnauthorized, "Invalid token claims", nil)
		}

		fields["user_id"] = userID
		fields["user_type"] = claims["user_type"]

		c.Locals("user_id", claims["user_id"])
		c.Locals("user_type", claims["user_type"])

		am.log.WithFields(fields).Debug("autentikasi berhasil")
		return c.Next()
	}
}
