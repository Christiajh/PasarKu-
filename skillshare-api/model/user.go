package model

import (
	// atau jwt/v5 jika kamu pakai versi v5
	"gorm.io/gorm"
)

// User represents a user in the system
type User struct {
	gorm.Model
	Name        string       `json:"name"`
	Email       string       `json:"email" gorm:"unique"`
	Password    string       `json:"password"`
	Classes     []Class      `json:"classes" gorm:"foreignKey:UserID"`
	Enrollments []Enrollment `json:"enrollments" gorm:"foreignKey:UserID"`
}

