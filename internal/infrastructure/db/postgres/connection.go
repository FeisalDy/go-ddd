package postgres

import (
	"fmt"
	"log"

	"github.com/FeisalDy/go-ddd/config"
	"github.com/FeisalDy/go-ddd/internal/infrastructure/db/postgres/migrations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init(cfg config.DBConfig) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	log.Println("Database connected")

	// Run migrations
	log.Println("Running database migrations...")
	if err := migrations.RunMigrations(DB); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}
	log.Println("Migrations completed successfully")
}
