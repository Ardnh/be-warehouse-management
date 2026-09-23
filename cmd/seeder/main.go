package main

import (
	"flag"
	"log"

	"github.com/Ardnh/be-warehouse-management/internal/config"
	"github.com/Ardnh/be-warehouse-management/internal/infrastructure/database/postgresql"
	"github.com/Ardnh/be-warehouse-management/internal/infrastructure/seeder"
	"github.com/joho/godotenv"
)

func main() {
	// seed := flag.Bool("seed", false, "run seeder after migration")
	flag.Parse()
	godotenv.Load()

	cfg := config.LoadConfig()

	db, err := postgresql.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}
	defer postgresql.CloseDB(db)

	// Seed permission
	log.Println("Running seeder...")
	if err := seeder.SeedPermissions(db); err != nil {
		log.Fatalf("❌ Failed to seed permissions: %v", err)
	}

	// Seed locations
	if err := seeder.SeedLocations(db); err != nil {
		log.Fatalf("❌ Failed to seed locations: %v", err)
	}

	// Seed user
	if err := seeder.SeedUser(db); err != nil {
		log.Fatalf("❌ Failed to seed user: %v", err)
	}

	// Seed customers
	if err := seeder.SeedCustomers(db); err != nil {
		log.Fatalf("❌ Failed to seed customers: %v", err)
	}

	// Seed UOMs
	if err := seeder.SeedUoms(db); err != nil {
		log.Fatalf("❌ Failed to seed UOMs: %v", err)
	}

	// Seed products using the existing customers and UOM records
	if err := seeder.SeedProducts(db); err != nil {
		log.Fatalf("❌ Failed to seed products: %v", err)
	}

	// Seed roles
	systemAdminRole, err := seeder.SeedSystemAdminRole(db)
	if err != nil {
		log.Fatalf("❌ Failed to seed system admin role: %v", err)
	}

	// Seed roles permissions
	if err := seeder.SeedSystemAdminPermissions(db, systemAdminRole); err != nil {
		log.Fatalf("❌ Failed to seed system admin permissions: %v", err)
	}

	log.Println("Done!")
}
