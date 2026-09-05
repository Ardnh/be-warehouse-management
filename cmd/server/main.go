package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	"github.com/Ardnh/be-warehouse-management/internal/application/services"
	"github.com/Ardnh/be-warehouse-management/internal/config"
	"github.com/Ardnh/be-warehouse-management/internal/infrastructure/database/postgresql"
	"github.com/Ardnh/be-warehouse-management/internal/infrastructure/database/redis"
	"github.com/Ardnh/be-warehouse-management/internal/infrastructure/repositories"
	"github.com/Ardnh/be-warehouse-management/internal/interface/handlers"
	"github.com/Ardnh/be-warehouse-management/internal/interface/routes"
)

func main() {
	godotenv.Load()
	app := fiber.New()
	log := logrus.New()
	cfg := config.LoadConfig()
	validator := validator.New()

	// Initialize database
	db, err := postgresql.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}
	defer postgresql.CloseDB(db)

	redisDb := redis.NewRedisDB(cfg)
	defer redisDb.Close()

	// Casbin
	// enforcer, err := utils.InitCasbin(modelPath, db)

	// Modules
	// Repository
	userRepository := repositories.NewUserRepository(db, redisDb)
	customerRepository := repositories.NewCustomerRepository(db, redisDb)
	warehouseRepository := repositories.NewWarehouseRepository(db)
	uomRepository := repositories.NewUomRepository(db)
	zoneRepository := repositories.NewZoneRepository(db)
	rackRepository := repositories.NewRackRepository(db)
	storageLocationRepository := repositories.NewStorageLocationRepository(db)

	// Service
	authService := services.NewAuthService(userRepository, log, cfg)
	customerService := services.NewCustomerService(customerRepository, log)
	warehouseService := services.NewWarehouseService(warehouseRepository, log)
	uomService := services.NewUomService(uomRepository, log)
	zoneService := services.NewZoneService(zoneRepository, log)
	rackService := services.NewRackService(rackRepository, log)
	storageLocationService := services.NewStorageLocationService(storageLocationRepository, log)

	// Handler
	authHandler := handlers.NewAuthHandler(authService, validator, log)
	customerHandler := handlers.NewCustomerHandler(customerService, validator, log)
	warehouseHandler := handlers.NewWarehouseHandler(warehouseService, validator, log)
	uomHandler := handlers.NewUomHandler(uomService, validator, log)
	zoneHandler := handlers.NewZoneHandler(zoneService, validator, log)
	rackHandler := handlers.NewRackHandler(rackService, validator, log)
	storageLocationHandler := handlers.NewStorageLocationHandler(storageLocationService, validator, log)

	routes.SetupAPIRoutes(app, log, validator, authHandler, customerHandler, warehouseHandler, uomHandler, zoneHandler, rackHandler, storageLocationHandler)

	log.Fatal(app.Listen(":3000"))
}
