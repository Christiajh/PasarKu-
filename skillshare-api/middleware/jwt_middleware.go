package middleware

import (
    "fmt"
    "net/http"
    "os"
    "skillshare-api/model"

    "github.com/golang-jwt/jwt/v5"
    "github.com/labstack/echo/v4"
    echojwt "github.com/labstack/echo-jwt/v4"
)

// JWTSecret retrieves the JWT secret key from environment variables.
func JWTSecret() string {
    secret := os.Getenv("JWT_SECRET")
    if secret == "" {
        fmt.Println("WARNING: JWT_SECRET environment variable is not set. Using a default fallback secret. ENSURE THIS IS INTENTIONAL FOR DEVELOPMENT. FOR PRODUCTION, ALWAYS SET VIA ENV VAR.")
        // Ini HARUS SAMA PERSIS dengan fallback di helper/jwt.go
        secret = "your_super_secure_and_long_consistent_key" // <-- Ganti dengan kunci rahasia asli Anda
    }
    return secret
}

func JWTMiddleware() echo.MiddlewareFunc {
    return echojwt.WithConfig(echojwt.Config{
        NewClaimsFunc: func(c echo.Context) jwt.Claims {
            return new(model.JwtCustomClaims)
        },
        SigningKey: []byte(JWTSecret()), // Menggunakan fungsi JWTSecret() yang sama
        ContextKey: "user",
        TokenLookup: "header:Authorization:Bearer",
        ErrorHandler: func(c echo.Context, err error) error {
            c.Logger().Errorf("JWT Middleware Error: %v", err)
            return echo.NewHTTPError(http.StatusUnauthorized, "invalid or expired jwt")
        },
    })
}