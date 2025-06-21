package middleware

import (
	"fmt"
	"net/http"
	"os"
	"skillshare-api/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
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
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(model.JwtCustomClaims)
		},
		SigningKey:  []byte(JWTSecret()),
		ContextKey:  "user",
		TokenLookup: "header:Authorization:Bearer",
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
