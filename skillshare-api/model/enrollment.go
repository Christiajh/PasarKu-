package model

import "gorm.io/gorm"

// Enrollment represents a user's enrollment in a class
type Enrollment struct {
	gorm.Model
	UserID  uint `json:"user_id"`
	User    User `json:"user"` // Belongs to relationship with User
	ClassID uint `json:"class_id"`
	Class   Class `json:"class"` // Belongs to relationship with Class
}