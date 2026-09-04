package routes

import (
	"github.com/Ardnh/be-warehouse-management/internal/interface/handlers"
	"github.com/Ardnh/be-warehouse-management/internal/interface/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
)

func SetupAPIRoutes(
	app *fiber.App,
	log *logrus.Logger,
	// enforcer *casbin.Enforcer,
	validator *validator.Validate,
	authHandler *handlers.AuthHandler,
) {
	// Middleware
	// casbinMw := middleware.NewCasbinMiddleware(enforcer, log)
	authMiddleware := middleware.NewAuthMiddleware()

	// API v1 group
	api := app.Group("/api/v1")

	// User
	user := api.Group("/user")
	user.Post("/register", authHandler.Register)
	user.Post("/login", authHandler.Login)
}
