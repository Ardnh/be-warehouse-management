package routes

import (
	"github.com/Ardnh/be-warehouse-management/internal/config"
	"github.com/Ardnh/be-warehouse-management/internal/interfaces/handlers"
	"github.com/Ardnh/be-warehouse-management/internal/interfaces/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
)

func SetupAPIRoutes(
	app *fiber.App,
	log *logrus.Logger,
	cfg *config.Config,
	// enforcer *casbin.Enforcer,
	validator *validator.Validate,
	authHandler *handlers.AuthHandler,
	customerHandler *handlers.CustomerHandler,
	warehouseHandler *handlers.WarehouseHandler,
	uomHandler *handlers.UomHandler,
	zoneHandler *handlers.ZoneHandler,
	rackHandler *handlers.RackHandler,
	storageLocationHandler *handlers.StorageLocationHandler,
	inboundOrderHandler *handlers.InboundOrderHandler,
	inboundOrderItemHandler *handlers.InboundOrderItemHandler,
	handlingUnitHandler *handlers.HandlingUnitHandler,
	handlingUnitItemHandler *handlers.HandlingUnitItemHandler,
	productHandler *handlers.ProductHandler,
	roleHandler *handlers.RoleHandler,
	receivingHandler *handlers.ReceivingHandler,
	receivingItemHandler *handlers.ReceivingItemHandler,
	permissionHandler *handlers.PermissionHandler,
	userHandler *handlers.UserHandler,
) {
	// Middleware
	// casbinMw := middleware.NewCasbinMiddleware(enforcer, log)
	authMiddleware := middleware.NewAuthMiddleware(log, cfg)

	// API v1 group
	api := app.Group("/api/v1")

	// Auth
	api.Post("/register", authHandler.Register)
	api.Post("/login", authHandler.Login)

	// User
	user := api.Group("/user", authMiddleware.Authenticate())
	user.Get("/", userHandler.FindAll)
	user.Get("/:id", userHandler.FindByID)
	user.Post("/", userHandler.Create)
	user.Put("/:id", userHandler.Update)
	user.Delete("/:id", userHandler.Delete)

	// Customer
	customer := api.Group("/customer", authMiddleware.Authenticate())
	customer.Get("/", customerHandler.FindAll)
	customer.Get("/:id", customerHandler.FindById)
	customer.Post("/", customerHandler.Create)
	customer.Put("/:id", customerHandler.Update)
	customer.Delete("/:id", customerHandler.Delete)

	// Warehouse
	warehouse := api.Group("/warehouse", authMiddleware.Authenticate())
	warehouse.Get("/", warehouseHandler.FindAll)
	warehouse.Get("/:id", warehouseHandler.FindByID)
	warehouse.Post("/", warehouseHandler.Create)
	warehouse.Put("/:id", warehouseHandler.Update)
	warehouse.Delete("/:id", warehouseHandler.Delete)

	// UOM
	uom := api.Group("/uom", authMiddleware.Authenticate())
	uom.Get("/", uomHandler.FindAll)
	uom.Get("/:id", uomHandler.FindByID)
	uom.Post("/", uomHandler.Create)
	uom.Put("/:id", uomHandler.Update)
	uom.Delete("/:id", uomHandler.Delete)

	// Zone
	zone := api.Group("/zone", authMiddleware.Authenticate())
	zone.Get("/", zoneHandler.FindAll)
	zone.Get("/:id", zoneHandler.FindByID)
	zone.Post("/", zoneHandler.Create)
	zone.Put("/:id", zoneHandler.Update)
	zone.Delete("/:id", zoneHandler.Delete)

	// Rack
	rack := api.Group("/rack", authMiddleware.Authenticate())
	rack.Get("/", rackHandler.FindAll)
	rack.Get("/:id", rackHandler.FindByID)
	rack.Post("/", rackHandler.Create)
	rack.Put("/:id", rackHandler.Update)
	rack.Delete("/:id", rackHandler.Delete)

	// Storage Location
	storageLocation := api.Group("/storage-location", authMiddleware.Authenticate())
	storageLocation.Get("/", storageLocationHandler.FindAll)
	storageLocation.Get("/:id", storageLocationHandler.FindByID)
	storageLocation.Post("/", storageLocationHandler.Create)
	storageLocation.Put("/:id", storageLocationHandler.Update)
	storageLocation.Delete("/:id", storageLocationHandler.Delete)

	// Inbound Order
	inboundOrder := api.Group("/inbound-order", authMiddleware.Authenticate())
	inboundOrder.Get("/", inboundOrderHandler.FindAll)
	inboundOrder.Get("/:id", inboundOrderHandler.FindByID)
	inboundOrder.Post("/", inboundOrderHandler.Create)
	inboundOrder.Put("/:id", inboundOrderHandler.Update)
	inboundOrder.Patch("/:id/status", inboundOrderHandler.UpdateStatus)
	inboundOrder.Delete("/:id", inboundOrderHandler.Delete)

	// Inbound Order Item
	inboundOrderItem := inboundOrder.Group("/:id/item")
	inboundOrderItem.Get("/", inboundOrderItemHandler.FindAll)
	inboundOrderItem.Get("/:item_id", inboundOrderItemHandler.FindByID)
	inboundOrderItem.Post("/", inboundOrderItemHandler.Create)
	inboundOrderItem.Put("/:item_id", inboundOrderItemHandler.Update)
	inboundOrderItem.Delete("/:item_id", inboundOrderItemHandler.Delete)

	// Handling Unit
	handlingUnit := api.Group("/handling-unit", authMiddleware.Authenticate())
	handlingUnit.Get("/", handlingUnitHandler.FindAll)
	handlingUnit.Get("/:id", handlingUnitHandler.FindByID)
	handlingUnit.Post("/", handlingUnitHandler.Create)
	handlingUnit.Put("/:id", handlingUnitHandler.Update)
	handlingUnit.Delete("/:id", handlingUnitHandler.Delete)

	// Handling Unit Item
	handlingUnitItem := handlingUnit.Group("/:id/item")
	handlingUnitItem.Get("/", handlingUnitItemHandler.FindAll)
	handlingUnitItem.Get("/:item_id", handlingUnitItemHandler.FindByID)
	handlingUnitItem.Post("/", handlingUnitItemHandler.Create)
	handlingUnitItem.Put("/:item_id", handlingUnitItemHandler.Update)
	handlingUnitItem.Delete("/:item_id", handlingUnitItemHandler.Delete)

	// Product
	product := api.Group("/product", authMiddleware.Authenticate())
	product.Get("/", productHandler.FindAll)
	product.Get("/:id", productHandler.FindByID)
	product.Post("/", productHandler.Create)
	product.Put("/:id", productHandler.Update)
	product.Delete("/:id", productHandler.Delete)

	// Role
	role := api.Group("/role", authMiddleware.Authenticate())
	role.Get("/", roleHandler.FindAll)
	role.Get("/:id", roleHandler.FindByID)
	role.Post("/", roleHandler.Create)
	role.Put("/:id", roleHandler.Update)
	role.Delete("/:id", roleHandler.Delete)

	// Receiving
	receiving := api.Group("/receiving", authMiddleware.Authenticate())
	receiving.Get("/", receivingHandler.FindAll)
	receiving.Get("/:id", receivingHandler.FindByID)
	receiving.Post("/", receivingHandler.Create)
	receiving.Put("/:id", receivingHandler.Update)
	receiving.Delete("/:id", receivingHandler.Delete)

	// Receiving Item
	receivingItem := receiving.Group("/:id/item", authMiddleware.Authenticate())
	receivingItem.Get("/", receivingItemHandler.FindAll)
	receivingItem.Get("/:item_id", receivingItemHandler.FindByID)
	receivingItem.Post("/", receivingItemHandler.Create)
	receivingItem.Put("/:item_id", receivingItemHandler.Update)
	receivingItem.Delete("/:item_id", receivingItemHandler.Delete)

	// Permission
	permission := api.Group("/permission", authMiddleware.Authenticate())
	permission.Get("/", permissionHandler.FindAll)
	permission.Get("/:id", permissionHandler.FindByID)
	permission.Post("/", permissionHandler.Create)
	permission.Put("/:id", permissionHandler.Update)
	permission.Delete("/:id", permissionHandler.Delete)
}
