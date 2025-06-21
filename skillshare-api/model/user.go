package model

import (
	"github.com/golang-jwt/jwt/v5" // Make sure this is the correct import for v5
	"gorm.io/gorm"
)

// User represents a user in the system
type User struct {
	gorm.Model
	Name        string       `json:"name"`
	Email       string       `json:"email" gorm:"unique"`
	Password    string       `json:"-"` // Exclude password from JSON output
	Classes     []Class      `json:"classes" gorm:"foreignKey:UserID"`
	Enrollments []Enrollment `json:"enrollments" gorm:"foreignKey:UserID"`
}

// JwtCustomClaims represents the claims in a JWT token
type JwtCustomClaims struct {
	UserID uint `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims // Corrected: Use RegisteredClaims from jwt/v5
}