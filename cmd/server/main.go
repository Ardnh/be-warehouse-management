package main

import (
	"os"
	"path/filepath"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	"github.com/Ardnh/be-warehouse-management/internal/application/services"
	"github.com/Ardnh/be-warehouse-management/internal/config"
	"github.com/Ardnh/be-warehouse-management/internal/infrastructure/database/postgresql"
	"github.com/Ardnh/be-warehouse-management/internal/infrastructure/database/redis"
	"github.com/Ardnh/be-warehouse-management/internal/infrastructure/repositories"
	"github.com/Ardnh/be-warehouse-management/internal/interfaces/handlers"
	"github.com/Ardnh/be-warehouse-management/internal/interfaces/routes"
	"github.com/Ardnh/be-warehouse-management/internal/utils/casbin"
)

func main() {
	godotenv.Load()
	app := fiber.New()
	log := logrus.New()
	cfg := config.LoadConfig()
	validator := validator.New()

	workDir, err := os.Getwd()
	if err != nil {
		log.Fatal("Failed to get working directory:", err)
	}
	modelPath := filepath.Join(workDir, "internal/utils/casbin", "model.conf")

	// Initialize database
	db, err := postgresql.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}
	defer postgresql.CloseDB(db)

	redisDb := redis.NewRedisDB(cfg)
	defer redisDb.Close()

	// Casbin
	enforcer, err := casbin.InitCasbin(modelPath, db)
	if err != nil {
		log.Fatalf("❌ Failed to initialize Casbin: %v", err)
	}
	defer enforcer.SavePolicy()

	// Modules
	// Repository
	userRepository := repositories.NewUserRepository(db, redisDb)
	customerRepository := repositories.NewCustomerRepository(db, redisDb)
	warehouseRepository := repositories.NewWarehouseRepository(db)
	uomRepository := repositories.NewUomRepository(db)
	zoneRepository := repositories.NewZoneRepository(db)
	rackRepository := repositories.NewRackRepository(db)
	storageLocationRepository := repositories.NewStorageLocationRepository(db)
	inboundOrderRepository := repositories.NewInboundOrderRepository(db)
	inboundOrderItemRepository := repositories.NewInboundOrderItemRepository(db)
	handlingUnitRepository := repositories.NewHandlingUnitRepository(db)
	handlingUnitItemRepository := repositories.NewHandlingUnitItemRepository(db)
	productRepository := repositories.NewProductRepository(db)
	roleRepository := repositories.NewRoleRepository(db)
	receivingRepository := repositories.NewReceivingRepository(db)
	receivingItemRepository := repositories.NewReceivingItemRepository(db)
	permissionRepository := repositories.NewPermissionRepository(db)

	// Service
	authService := services.NewAuthService(userRepository, log, cfg)
	customerService := services.NewCustomerService(customerRepository, log)
	warehouseService := services.NewWarehouseService(warehouseRepository, log)
	uomService := services.NewUomService(uomRepository, log)
	zoneService := services.NewZoneService(zoneRepository, log)
	rackService := services.NewRackService(rackRepository, log)
	storageLocationService := services.NewStorageLocationService(storageLocationRepository, log)
	inboundOrderService := services.NewInboundOrderService(inboundOrderRepository, inboundOrderItemRepository, log)
	inboundOrderItemService := services.NewInboundOrderItemService(inboundOrderItemRepository, log)
	handlingUnitService := services.NewHandlingUnitService(handlingUnitRepository, log)
	handlingUnitItemService := services.NewHandlingUnitItemService(handlingUnitItemRepository, log)
	productService := services.NewProductService(productRepository, log)
	roleService := services.NewRoleService(roleRepository, log)
	receivingService := services.NewReceivingService(receivingRepository, receivingItemRepository, log)
	receivingItemService := services.NewReceivingItemService(receivingItemRepository, log)
	permissionService := services.NewPermissionService(permissionRepository, log)

	// Handler
	authHandler := handlers.NewAuthHandler(authService, validator, log)
	customerHandler := handlers.NewCustomerHandler(customerService, validator, log)
	warehouseHandler := handlers.NewWarehouseHandler(warehouseService, validator, log)
	uomHandler := handlers.NewUomHandler(uomService, validator, log)
	zoneHandler := handlers.NewZoneHandler(zoneService, validator, log)
	rackHandler := handlers.NewRackHandler(rackService, validator, log)
	storageLocationHandler := handlers.NewStorageLocationHandler(storageLocationService, validator, log)
	inboundOrderHandler := handlers.NewInboundOrderHandler(inboundOrderService, validator, log)
	inboundOrderItemHandler := handlers.NewInboundOrderItemHandler(inboundOrderItemService, validator, log)
	handlingUnitHandler := handlers.NewHandlingUnitHandler(handlingUnitService, validator, log)
	handlingUnitItemHandler := handlers.NewHandlingUnitItemHandler(handlingUnitItemService, validator, log)
	productHandler := handlers.NewProductHandler(productService, validator, log)
	roleHandler := handlers.NewRoleHandler(roleService, validator, log)
	receivingHandler := handlers.NewReceivingHandler(receivingService, validator, log)
	receivingItemHandler := handlers.NewReceivingItemHandler(receivingItemService, validator, log)
	permissionHandler := handlers.NewPermissionHandler(permissionService, validator, log)

	routes.SetupAPIRoutes(
		app,
		log,
		cfg,
		validator,
		authHandler,
		customerHandler,
		warehouseHandler,
		uomHandler,
		zoneHandler,
		rackHandler,
		storageLocationHandler,
		inboundOrderHandler,
		inboundOrderItemHandler,
		handlingUnitHandler,
		handlingUnitItemHandler,
		productHandler,
		roleHandler,
		receivingHandler,
		receivingItemHandler,
		permissionHandler,
	)

	log.Fatal(app.Listen(":3000"))
}
