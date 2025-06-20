package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"pasarku/model"
)

var DB *gorm.DB

func Init() {
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	DB = db
	db.AutoMigrate(&model.User{}, &model.Product{}, &model.Order{}, &model.Review{})
}
