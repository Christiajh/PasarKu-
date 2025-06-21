package helper

import (
	"os"
	"skillshare-api/model"
	"time"

	"github.com/golang-jwt/jwt/v5" // Make sure this is the correct import for v5
)

// JWTSecret returns the JWT secret key from environment variables
func JWTSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "your_jwt_secret_key" // Fallback for local development
	}
	return secret
}

// GenerateJWT generates a new JWT token
func GenerateJWT(userID uint, email string) (string, error) {
	claims := &model.JwtCustomClaims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{ // Corrected: Use RegisteredClaims
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // Corrected for jwt/v5
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	// This syntax `jwt.NewWithClaims(jwt.SigningMethodHS256, claims)` is correct for jwt/v5
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(JWTSecret()))
	if err != nil {
		return "", err
	}
	return t, nil
}