package model

import "gorm.io/gorm"

// Class represents a skill class
type Class struct {
	gorm.Model
	Title       string       `json:"title"`
	Description string       `json:"description"`

	UserID  uint   `json:"user_id"`                  // FK to users
	User    User   `json:"-" gorm:"foreignKey:UserID"` // 🧠 disembunyikan agar tidak rekursif di JSON

	CategoryID uint     `json:"category_id"`         // FK to categories
	Category   Category `json:"-" gorm:"foreignKey:CategoryID"` // sembunyikan dari JSON jika tidak diperlukan

	Enrollments []Enrollment `json:"-" gorm:"foreignKey:ClassID"` // optional: sembunyikan untuk efisiensi
}
