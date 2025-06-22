package model

import "github.com/golang-jwt/jwt/v4"

// JwtCustomClaims represents the claims in a JWT token
type JwtCustomClaims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

// Valid validates the token claims (required by jwt library)
func (c *JwtCustomClaims) Valid() error {
	return c.RegisteredClaims.Valid()
}
