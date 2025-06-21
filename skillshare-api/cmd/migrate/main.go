package main

import (
	"fmt"
	"log"
	"os"
	"skillshare-api/config"
	"skillshare-api/migration"

	"github.com/joho/godotenv"
)

func main() {
	// Hanya load .env di lokal (jangan di Railway)
	if os.Getenv("RAILWAY_ENVIRONMENT") == "" {
		if err := godotenv.Load(); err != nil {
			log.Println("⚠️ .env file not found (this is fine in Railway)")
		}
	}

	// Koneksi ke database via DATABASE_URL
	db := config.ConnectDatabase()

	// Jalankan migrasi
	fmt.Println("🚀 Running database migration...")
	migration.Migrate(db)
	fmt.Println("✅ Migration completed successfully.")
}
