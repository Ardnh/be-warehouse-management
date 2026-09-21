package main

import (
	"flag"
	"log"

	"github.com/Ardnh/be-warehouse-management/internal/config"
	"github.com/Ardnh/be-warehouse-management/internal/infrastructure/database/postgresql"
	"github.com/Ardnh/be-warehouse-management/internal/infrastructure/migration"
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

	log.Println("Running migration...")
	if err := migration.Migrate(db); err != nil {
		log.Fatal(err)
	}

	// Seed permission
	log.Println("Running seeder...")
	if err := seeder.SeedPermissions(db); err != nil {
		log.Fatalf("❌ Failed to seed permissions: %v", err)
	}

	log.Println("Done!")
}
