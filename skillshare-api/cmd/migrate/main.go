// File: cmd/migrate/main.go
package main

import (
	"fmt"
	"log"
	"skillshare-api/config"
	"skillshare-api/migration"

	"github.com/joho/godotenv"
	"os"
)

func main() {
	// Load .env
	if os.Getenv("RAILWAY_ENVIRONMENT") == "" {
    if err := godotenv.Load(); err != nil {
        log.Println("⚠️ .env file not found (this is fine in production)")
    }
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
