package seeder

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

	// Seed user
	if err := seeder.SeedUser(db); err != nil {
		log.Fatalf("❌ Failed to seed user: %v", err)
	}

	log.Println("Done!")
}
