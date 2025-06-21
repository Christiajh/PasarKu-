package migration

import (
	"log"
	"skillshare-api/model"

	"gorm.io/gorm"
)


func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.User{},
		&model.Category{},
		&model.Class{},
		&model.Enrollment{},
	)
	if err != nil {
		log.Fatalf("❌ Failed to migrate database: %v", err)
	}
	log.Println("✅ Database migration completed successfully!")
}
