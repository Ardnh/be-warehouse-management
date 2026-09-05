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

	// 2. Initialize database
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

	// Service
	authService := services.NewAuthService(userRepository, log, cfg)

	// Handler
	authHandler := handlers.NewAuthHandler(authService, validator, log)

	routes.SetupAPIRoutes(app, log, validator, authHandler)

	log.Fatal(app.Listen(":3000"))
}
