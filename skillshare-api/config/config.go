package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDB initializes the database connection
func InitDB() *gorm.DB {
	// Load .env only in local environment
	if os.Getenv("RAILWAY_ENVIRONMENT") == "" {
		if err := godotenv.Load(); err != nil {
			log.Println("⚠️ No .env file found (this is okay for Railway)")
		}
	}

	// If DATABASE_URL is available (e.g., in Railway), use it directly
	if dsn := os.Getenv("DATABASE_URL"); dsn != "" {
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("❌ Failed to connect to database via DATABASE_URL: %v", err)
		}
		log.Println("✅ Connected to database via DATABASE_URL")
		return db
	}

	// Fallback for local development (with manual fields)
	dbHost := GetEnv("DB_HOST", "localhost")
	dbPort := GetEnv("DB_PORT", "5432")
	dbUser := GetEnv("DB_USER", "your_user")
	dbPassword := GetEnv("DB_PASSWORD", "your_password")
	dbName := GetEnv("DB_NAME", "skillshare_db")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		dbHost, dbUser, dbPassword, dbName, dbPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to database (manual config): %v", err)
	}
	log.Println("✅ Connected to database using manual config")
	return db
}

// GetEnv retrieves the value of an environment variable or returns a default value
func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
