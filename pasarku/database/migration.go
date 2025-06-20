package database

import (
	"fmt"
	"log"
	"pasarku/model"
)

// RunMigration melakukan auto-migrate semua tabel
func RunMigration() {
	err := DB.AutoMigrate(
		&model.User{},
		&model.Product{},
		&model.Order{},
		&model.Review{},
		&model.Category{},
		&model.Tag{},
	)
	if err != nil {
		log.Fatalf("❌ Failed to run migration: %v", err)
	}

	fmt.Println("✅ Database migrated successfully!")
}
