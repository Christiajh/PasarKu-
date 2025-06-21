package helper

import (
    "os"
    "skillshare-api/model"
    "time"

    "github.com/golang-jwt/jwt/v5"
)

// JWTSecret returns the JWT secret key from environment variables.
// In a production environment, ensure JWT_SECRET is always set.
func JWTSecret() string {
    secret := os.Getenv("JWT_SECRET")
    if secret == "" {
        // Jika Anda ingin fallback untuk lokal, gunakan nilai yang SAMA PERSIS dengan di middleware
        return "your_super_secure_and_long_consistent_key" // <-- Ganti dengan kunci rahasia asli Anda
    }
    return secret
}

func GenerateJWT(userID uint, email string) (string, error) {
    claims := &model.JwtCustomClaims{
        UserID: userID,
        Email:  email,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
            NotBefore: jwt.NewNumericDate(time.Now()),
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    t, err := token.SignedString([]byte(JWTSecret())) // Menggunakan fungsi JWTSecret() yang sama
    if err != nil {
        return "", err
    }
    return t, nil
}