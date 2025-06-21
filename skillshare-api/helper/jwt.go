package helper

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"skillshare-api/model"
	"fmt"
)

// JWTSecret returns the secret key for signing JWT
func JWTSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		fmt.Println("⚠️ WARNING: JWT_SECRET not set. Using fallback. DO NOT use in production.")
		secret = "TEST_SECRET_123"
	}
	return secret
}

// GenerateJWT creates a JWT token for the given user
func GenerateJWT(userID uint, email string) (string, error) {
	claims := model.JwtCustomClaims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)), // Token expires in 1 hour
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()), // Token valid immediately
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)

	signedToken, err := token.SignedString([]byte(JWTSecret()))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
