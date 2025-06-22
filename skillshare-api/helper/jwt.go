package helper

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"skillshare-api/model"
)

// JWTSecret returns the JWT secret from environment or fallback
func JWTSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		fmt.Println("⚠️ WARNING: JWT_SECRET not set. Using fallback. DO NOT use in production.")
		secret = "TEST_SECRET_123"
	}
	return secret
}

// GenerateJWT creates a JWT token for the given user with long expiration
func GenerateJWT(userID uint, email string) (string, error) {
	claims := model.JwtCustomClaims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(10, 0, 0)), // ✅ 10 tahun ke depan
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)

	signedToken, err := token.SignedString([]byte(JWTSecret()))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
