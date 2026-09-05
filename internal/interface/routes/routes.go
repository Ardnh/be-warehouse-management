package routes

import (
	"github.com/Ardnh/be-warehouse-management/internal/interface/handlers"
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
	customerHandler *handlers.CustomerHandler,
	warehouseHandler *handlers.WarehouseHandler,
	uomHandler *handlers.UomHandler,
	zoneHandler *handlers.ZoneHandler,
	rackHandler *handlers.RackHandler,
	storageLocationHandler *handlers.StorageLocationHandler,
) {
	// Middleware
	// casbinMw := middleware.NewCasbinMiddleware(enforcer, log)
	// authMiddleware := middleware.NewAuthMiddleware()

	// API v1 group
	api := app.Group("/api/v1")

	// User
	user := api.Group("/user")
	user.Post("/register", authHandler.Register)
	user.Post("/login", authHandler.Login)

	// Customer
	customer := api.Group("/customer")
	customer.Get("/", customerHandler.FindAll)
	customer.Get("/:id", customerHandler.FindById)
	customer.Post("/", customerHandler.Create)
	customer.Put("/:id", customerHandler.Update)
	customer.Delete("/:id", customerHandler.Delete)

	warehouse := api.Group("/warehouse")
	warehouse.Get("/", warehouseHandler.FindAll)
	warehouse.Get("/:id", warehouseHandler.FindByID)
	warehouse.Post("/", warehouseHandler.Create)
	warehouse.Put("/:id", warehouseHandler.Update)
	warehouse.Delete("/:id", warehouseHandler.Delete)

	uom := api.Group("/uom")
	uom.Get("/", uomHandler.FindAll)
	uom.Get("/:id", uomHandler.FindByID)
	uom.Post("/", uomHandler.Create)
	uom.Put("/:id", uomHandler.Update)
	uom.Delete("/:id", uomHandler.Delete)

	zone := api.Group("/zone")
	zone.Get("/", zoneHandler.FindAll)
	zone.Get("/:id", zoneHandler.FindByID)
	zone.Post("/", zoneHandler.Create)
	zone.Put("/:id", zoneHandler.Update)
	zone.Delete("/:id", zoneHandler.Delete)

	rack := api.Group("/rack")
	rack.Get("/", rackHandler.FindAll)
	rack.Get("/:id", rackHandler.FindByID)
	rack.Post("/", rackHandler.Create)
	rack.Put("/:id", rackHandler.Update)
	rack.Delete("/:id", rackHandler.Delete)

	storageLocation := api.Group("/storage-location")
	storageLocation.Get("/", storageLocationHandler.FindAll)
	storageLocation.Get("/:id", storageLocationHandler.FindByID)
	storageLocation.Post("/", storageLocationHandler.Create)
	storageLocation.Put("/:id", storageLocationHandler.Update)
	storageLocation.Delete("/:id", storageLocationHandler.Delete)
}
