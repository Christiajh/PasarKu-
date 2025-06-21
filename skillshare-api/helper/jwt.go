package helper

import (
	"os"
	"skillshare-api/model"
	"time"

	"github.com/golang-jwt/jwt/v4" // Pastikan versi v4
)

// JWTSecret returns the JWT secret key from environment variables.
// Ini adalah sumber kebenaran untuk secret key saat token DIBUAT.
func JWTSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		// HARUS sama dengan middleware jika JWT_SECRET tidak diset di environment.
		// Pesan warning hanya di middleware agar tidak terlalu banyak log.
		return "your_super_secure_and_long_consistent_key"
	}
	return secret
}

// GenerateJWT generates a JWT token for a user.
func GenerateJWT(userID uint, email string) (string, error) {
	now := time.Now()

	claims := &model.JwtCustomClaims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(24 * time.Hour)),    // Berlaku 24 jam
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now.Add(-5 * time.Second)), // Sedikit buffer waktu
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(JWTSecret()))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}