// File: cmd/migrate/main.go
package main

import (
	"fmt"
	"log"
	"skillshare-api/config"
	"skillshare-api/migration"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	// Connect ke database
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}

	// Jalankan migrasi
	fmt.Println("🚀 Running database migration...")
	migration.Migrate(db)
}
