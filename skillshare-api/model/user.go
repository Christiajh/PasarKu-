package model

import (
	"github.com/golang-jwt/jwt/v4" // atau jwt/v5 jika kamu pakai versi v5
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

// JwtCustomClaims represents the claims in a JWT token
type JwtCustomClaims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

// ✅ FIX: Valid method with pointer receiver
func (c *JwtCustomClaims) Valid() error {
	return c.RegisteredClaims.Valid()
}
