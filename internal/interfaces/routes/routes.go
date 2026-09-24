package routes

import (
	"github.com/Ardnh/be-warehouse-management/internal/interfaces/handlers"
	"github.com/Ardnh/be-warehouse-management/internal/interfaces/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func SetupAPIRoutes(
	app *fiber.App,
	authMiddleware *middleware.AuthMiddleware,
	authzMiddleware *middleware.AuthorizationMiddleware,
	validator *validator.Validate,
	authHandler *handlers.AuthHandler,
	customerHandler *handlers.CustomerHandler,
	warehouseHandler *handlers.WarehouseHandler,
	locationHandler *handlers.LocationHandler,
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

	// API v1 group
	api := app.Group("/api/v1")

	// Auth
	api.Post("/register", authHandler.Register)
	api.Post("/login", authHandler.Login)

	// User
	user := api.Group("/user", authMiddleware.Authenticate())
	user.Get("/profile", authzMiddleware.Require("user", "read"), userHandler.GetProfile)
	user.Get("/", authzMiddleware.Require("user", "read"), userHandler.FindAll)
	user.Post("/", authzMiddleware.Require("user", "create"), userHandler.Create)
	user.Put("/:id", authzMiddleware.Require("user", "update"), userHandler.Update)
	user.Delete("/:id", authzMiddleware.Require("user", "delete"), userHandler.Delete)

	// Customer
	customer := api.Group("/customer", authMiddleware.Authenticate())
	customer.Get("/", authzMiddleware.Require("customer", "read"), customerHandler.FindAll)
	customer.Get("/:id", authzMiddleware.Require("customer", "read"), customerHandler.FindById)
	customer.Post("/", authzMiddleware.Require("customer", "create"), customerHandler.Create)
	customer.Put("/:id", authzMiddleware.Require("customer", "update"), customerHandler.Update)
	customer.Delete("/:id", authzMiddleware.Require("customer", "delete"), customerHandler.Delete)

	// Warehouse
	warehouse := api.Group("/warehouse", authMiddleware.Authenticate())
	warehouse.Get("/", authzMiddleware.Require("warehouse", "read"), warehouseHandler.FindAll)
	warehouse.Get("/:id", authzMiddleware.Require("warehouse", "read"), warehouseHandler.FindByID)
	warehouse.Post("/", authzMiddleware.Require("warehouse", "create"), warehouseHandler.Create)
	warehouse.Put("/:id", authzMiddleware.Require("warehouse", "update"), warehouseHandler.Update)
	warehouse.Delete("/:id", authzMiddleware.Require("warehouse", "delete"), warehouseHandler.Delete)

	// Location
	location := api.Group("/location", authMiddleware.Authenticate())
	location.Get("/", authzMiddleware.Require("location", "read"), locationHandler.FindAll)
	location.Get("/:id", authzMiddleware.Require("location", "read"), locationHandler.FindByID)
	location.Post("/", authzMiddleware.Require("location", "create"), locationHandler.Create)
	location.Put("/:id", authzMiddleware.Require("location", "update"), locationHandler.Update)
	location.Delete("/:id", authzMiddleware.Require("location", "delete"), locationHandler.Delete)

	// UOM
	uom := api.Group("/uom", authMiddleware.Authenticate())
	uom.Get("/", authzMiddleware.Require("uoms", "read"), uomHandler.FindAll)
	uom.Get("/:id", authzMiddleware.Require("uoms", "read"), uomHandler.FindByID)
	uom.Post("/", authzMiddleware.Require("uoms", "create"), uomHandler.Create)
	uom.Put("/:id", authzMiddleware.Require("uoms", "update"), uomHandler.Update)
	uom.Delete("/:id", authzMiddleware.Require("uoms", "delete"), uomHandler.Delete)

	// Zone
	zone := api.Group("/zone", authMiddleware.Authenticate())
	zone.Get("/", authzMiddleware.Require("zone", "read"), zoneHandler.FindAll)
	zone.Get("/:id", authzMiddleware.Require("zone", "read"), zoneHandler.FindByID)
	zone.Post("/", authzMiddleware.Require("zone", "create"), zoneHandler.Create)
	zone.Put("/:id", authzMiddleware.Require("zone", "update"), zoneHandler.Update)
	zone.Delete("/:id", authzMiddleware.Require("zone", "delete"), zoneHandler.Delete)

	// Rack
	rack := api.Group("/rack", authMiddleware.Authenticate())
	rack.Get("/", authzMiddleware.Require("rack", "read"), rackHandler.FindAll)
	rack.Get("/:id", authzMiddleware.Require("rack", "read"), rackHandler.FindByID)
	rack.Post("/", authzMiddleware.Require("rack", "create"), rackHandler.Create)
	rack.Put("/:id", authzMiddleware.Require("rack", "update"), rackHandler.Update)
	rack.Delete("/:id", authzMiddleware.Require("rack", "delete"), rackHandler.Delete)

	// Storage Location
	storageLocation := api.Group("/storage-location", authMiddleware.Authenticate())
	storageLocation.Get("/", authzMiddleware.Require("storage-location", "read"), storageLocationHandler.FindAll)
	storageLocation.Get("/:id", authzMiddleware.Require("storage-location", "read"), storageLocationHandler.FindByID)
	storageLocation.Post("/", authzMiddleware.Require("storage-location", "create"), storageLocationHandler.Create)
	storageLocation.Put("/:id", authzMiddleware.Require("storage-location", "update"), storageLocationHandler.Update)
	storageLocation.Delete("/:id", authzMiddleware.Require("storage-location", "delete"), storageLocationHandler.Delete)

	// Inbound Order
	inboundOrder := api.Group("/inbound-order", authMiddleware.Authenticate())
	inboundOrder.Get("/", authzMiddleware.Require("inbound-order", "read"), inboundOrderHandler.FindAll)
	inboundOrder.Get("/:id", authzMiddleware.Require("inbound-order", "read"), inboundOrderHandler.FindByID)
	inboundOrder.Post("/", authzMiddleware.Require("inbound-order", "create"), inboundOrderHandler.Create)
	inboundOrder.Put("/:id", authzMiddleware.Require("inbound-order", "update"), inboundOrderHandler.Update)
	inboundOrder.Patch("/:id/status", authzMiddleware.Require("inbound-order", "update"), inboundOrderHandler.UpdateStatus)
	inboundOrder.Delete("/:id", authzMiddleware.Require("inbound-order", "delete"), inboundOrderHandler.Delete)

	// Inbound Order Item
	inboundOrderItem := inboundOrder.Group("/:id/item")
	inboundOrderItem.Get("/", authzMiddleware.Require("inbound-order-item", "read"), inboundOrderItemHandler.FindAll)
	inboundOrderItem.Get("/:item_id", authzMiddleware.Require("inbound-order-item", "read"), inboundOrderItemHandler.FindByID)
	inboundOrderItem.Post("/", authzMiddleware.Require("inbound-order-item", "create"), inboundOrderItemHandler.Create)
	inboundOrderItem.Put("/:item_id", authzMiddleware.Require("inbound-order-item", "update"), inboundOrderItemHandler.Update)
	inboundOrderItem.Delete("/:item_id", authzMiddleware.Require("inbound-order-item", "delete"), inboundOrderItemHandler.Delete)

	// Handling Unit
	handlingUnit := api.Group("/handling-unit", authMiddleware.Authenticate())
	handlingUnit.Get("/", authzMiddleware.Require("handling-unit", "read"), handlingUnitHandler.FindAll)
	handlingUnit.Get("/:id", authzMiddleware.Require("handling-unit", "read"), handlingUnitHandler.FindByID)
	handlingUnit.Post("/", authzMiddleware.Require("handling-unit", "create"), handlingUnitHandler.Create)
	handlingUnit.Put("/:id", authzMiddleware.Require("handling-unit", "update"), handlingUnitHandler.Update)
	handlingUnit.Delete("/:id", authzMiddleware.Require("handling-unit", "delete"), handlingUnitHandler.Delete)

	// Handling Unit Item
	handlingUnitItem := handlingUnit.Group("/:id/item")
	handlingUnitItem.Get("/", authzMiddleware.Require("handling-unit-item", "read"), handlingUnitItemHandler.FindAll)
	handlingUnitItem.Get("/:item_id", authzMiddleware.Require("handling-unit-item", "read"), handlingUnitItemHandler.FindByID)
	handlingUnitItem.Post("/", authzMiddleware.Require("handling-unit-item", "create"), handlingUnitItemHandler.Create)
	handlingUnitItem.Put("/:item_id", authzMiddleware.Require("handling-unit-item", "update"), handlingUnitItemHandler.Update)
	handlingUnitItem.Delete("/:item_id", authzMiddleware.Require("handling-unit-item", "delete"), handlingUnitItemHandler.Delete)

	// Product
	product := api.Group("/product", authMiddleware.Authenticate())
	product.Get("/", authzMiddleware.Require("product", "read"), productHandler.FindAll)
	product.Get("/:id", authzMiddleware.Require("product", "read"), productHandler.FindByID)
	product.Post("/", authzMiddleware.Require("product", "create"), productHandler.Create)
	product.Put("/:id", authzMiddleware.Require("product", "update"), productHandler.Update)
	product.Delete("/:id", authzMiddleware.Require("product", "delete"), productHandler.Delete)

	// Role
	role := api.Group("/role", authMiddleware.Authenticate())
	role.Get("/", authzMiddleware.Require("role", "read"), roleHandler.FindAll)
	role.Get("/:id", authzMiddleware.Require("role", "read"), roleHandler.FindByID)
	role.Post("/", authzMiddleware.Require("role", "create"), roleHandler.Create)
	role.Put("/:id", authzMiddleware.Require("role", "update"), roleHandler.Update)
	role.Delete("/:id", authzMiddleware.Require("role", "delete"), roleHandler.Delete)

	// Receiving
	receiving := api.Group("/receiving", authMiddleware.Authenticate())
	receiving.Get("/", authzMiddleware.Require("receiving", "read"), receivingHandler.FindAll)
	receiving.Get("/:id", authzMiddleware.Require("receiving", "read"), receivingHandler.FindByID)
	receiving.Post("/", authzMiddleware.Require("receiving", "create"), receivingHandler.Create)
	receiving.Put("/:id", authzMiddleware.Require("receiving", "update"), receivingHandler.Update)
	receiving.Delete("/:id", authzMiddleware.Require("receiving", "delete"), receivingHandler.Delete)

	// Receiving Item
	receivingItem := receiving.Group("/:id/item", authMiddleware.Authenticate())
	receivingItem.Get("/", authzMiddleware.Require("receiving_item", "read"), receivingItemHandler.FindAll)
	receivingItem.Get("/:item_id", authzMiddleware.Require("receiving_item", "read"), receivingItemHandler.FindByID)
	receivingItem.Post("/", authzMiddleware.Require("receiving_item", "create"), receivingItemHandler.Create)
	receivingItem.Put("/:item_id", authzMiddleware.Require("receiving_item", "update"), receivingItemHandler.Update)
	receivingItem.Delete("/:item_id", authzMiddleware.Require("receiving_item", "delete"), receivingItemHandler.Delete)

	// Permission
	permission := api.Group("/permission", authMiddleware.Authenticate())
	permission.Get("/", authzMiddleware.Require("permission", "read"), permissionHandler.FindAll)
	permission.Get("/:id", authzMiddleware.Require("permission", "read"), permissionHandler.FindByID)
	permission.Post("/", authzMiddleware.Require("permission", "create"), permissionHandler.Create)
	permission.Put("/:id", authzMiddleware.Require("permission", "update"), permissionHandler.Update)
	permission.Delete("/:id", authzMiddleware.Require("permission", "delete"), permissionHandler.Delete)
}
