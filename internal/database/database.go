package database

import (
	"fmt"
	"log"
	"os"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"github.com/joho/godotenv"
	"github.com/bpietrzakk/swift_codes/internal/models"
)

var DB *gorm.DB

func Connect() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Cannot load .env: %v", err)
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, name,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("cannot connect with database: %v", err)
	}

	if err := db.AutoMigrate(&models.SwiftCode{}); err != nil {
		log.Fatalf("cannot migrate database: %v", err)
	}

	DB = db
	log.Println(("Polaczono z PostgreSQL i wykonano migracje"))
}