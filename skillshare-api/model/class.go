package model

import "gorm.io/gorm"

// Class represents a skill class
type Class struct {
	gorm.Model
	Title       string       `json:"title"`
	Description string       `json:"description"`
	UserID      uint         `json:"user_id"` // Foreign key to User
	User        User         `json:"user"`    // Belongs to relationship with User
	CategoryID  uint         `json:"category_id"` // Foreign key to Category
	Category    Category     `json:"category"`    // Belongs to relationship with Category
	Enrollments []Enrollment `json:"enrollments" gorm:"foreignKey:ClassID"`
}