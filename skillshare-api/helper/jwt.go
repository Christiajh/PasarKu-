package helper

import (
    "os"
    "skillshare-api/model"
    "time"

   "github.com/golang-jwt/jwt/v4"

)

// JWTSecret returns the JWT secret key from environment variables.
func JWTSecret() string {
    secret := os.Getenv("JWT_SECRET")
    if secret == "" {
        // HARUS sama dengan middleware
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
            ExpiresAt: jwt.NewNumericDate(now.Add(24 * time.Hour)), // berlaku 24 jam
            IssuedAt:  jwt.NewNumericDate(now),
            NotBefore: jwt.NewNumericDate(now.Add(-5 * time.Second)), // sedikit buffer waktu
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    signedToken, err := token.SignedString([]byte(JWTSecret()))
    if err != nil {
        return "", err
    }
    return signedToken, nil
}
