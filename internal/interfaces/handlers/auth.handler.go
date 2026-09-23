package handlers

import (
	"github.com/Ardnh/be-warehouse-management/internal/application/dto"
	"github.com/Ardnh/be-warehouse-management/internal/domain/services"
	responses "github.com/Ardnh/be-warehouse-management/internal/interfaces/response"
	validator_utils "github.com/Ardnh/be-warehouse-management/internal/utils/validator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
)

type AuthHandler struct {
	authService services.AuthService
	validator   *validator.Validate
	log         *logrus.Logger
}

func NewAuthHandler(authService services.AuthService, validator *validator.Validate, log *logrus.Logger) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		validator:   validator,
		log:         log,
	}
}

// baseFields menyiapkan field yang selalu ikut di setiap log entry.
func (h *AuthHandler) baseFields(c fiber.Ctx, handler string) logrus.Fields {
	return logrus.Fields{
		"handler":    handler,
		"method":     c.Method(),
		"path":       c.Path(),
		"ip":         c.IP(),
		"user_agent": c.Get("User-Agent"),
		"request_id": c.Get("X-Request-ID"),
	}
}

func (h *AuthHandler) Login(c fiber.Ctx) error {
	fields := h.baseFields(c, "Login")

	var req dto.LoginRequestDto
	if err := c.Bind().Body(&req); err != nil {
		h.log.WithFields(fields).WithError(err).Warn("gagal parsing body login")
		return responses.HandleError(c, err)
	}

	// Aman ditambahkan SETELAH bind, karena baru di sini identifier terisi.
	// Jangan pernah masukkan req.Password ke field manapun.
	fields["username"] = req.Username

	if err := h.validator.Struct(&req); err != nil {
		h.log.WithFields(fields).WithError(err).Warn("validasi request login gagal")
		return responses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, validator_utils.FormatValidationErrors(err))
	}

	result, err := h.authService.Login(c.Context(), req)
	if err != nil {
		// Login gagal itu peristiwa keamanan — pakai Warn, bukan Error,
		// karena penyebab tersering adalah password salah, bukan bug.
		h.log.WithFields(fields).WithError(err).Warn("login gagal")
		return responses.HandleError(c, err)
	}

	h.log.WithFields(fields).Info("login berhasil")
	return responses.NewSuccessResponse(c, fiber.StatusOK, "Login successful", result)
}

func (h *AuthHandler) Register(c fiber.Ctx) error {
	fields := h.baseFields(c, "Register")

	var req dto.RegisterRequestDto
	if err := c.Bind().Body(&req); err != nil {
		h.log.WithFields(fields).WithError(err).Warn("gagal parsing body register")
		return responses.HandleError(c, err)
	}

	fields["username"] = req.Username
	fields["email"] = req.Email

	if err := h.validator.Struct(&req); err != nil {
		h.log.WithFields(fields).WithError(err).Warn("validasi request register gagal")
		return responses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, validator_utils.FormatValidationErrors(err))
	}

	if err := h.authService.Register(c.Context(), req); err != nil {
		h.log.WithFields(fields).WithError(err).Error("registrasi gagal")
		return responses.HandleError(c, err)
	}

	h.log.WithFields(fields).Info("registrasi berhasil")
	return responses.NewSuccessResponse(c, fiber.StatusCreated, "Register successful", nil)
}
