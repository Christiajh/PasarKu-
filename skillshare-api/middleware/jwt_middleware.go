package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	
	"github.com/labstack/echo/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
)

// JWTSecret retrieves the JWT secret key from environment variables.
func JWTSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		fmt.Println("⚠️ WARNING: JWT_SECRET not set. Using fallback. DO NOT use in production.")
		secret = "your_super_secure_and_long_consistent_key"
	}
	return secret
}

// JWTMiddleware returns a middleware that validates JWT tokens.
func JWTMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:    []byte(JWTSecret()),
		SigningMethod: "HS256",
		ContextKey:    "user",
		TokenLookup:   "header:Authorization:Bearer",
		ErrorHandler: func(c echo.Context, err error) error {
			authHeader := c.Request().Header.Get("Authorization")
			fmt.Println("📥 JWT Middleware menerima Authorization:", authHeader)
			fmt.Println("🔐 JWT_SECRET yang digunakan:", JWTSecret())
			fmt.Println("🕒 Waktu server:", time.Now().Format(time.RFC3339))
			fmt.Printf("❌ JWT Middleware Error: %v\n", err)

			return echo.NewHTTPError(http.StatusUnauthorized, map[string]interface{}{
				"message": "invalid or expired jwt",
				"detail":  err.Error(),
			})
		},
	})
}
