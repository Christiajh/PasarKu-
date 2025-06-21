package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"skillshare-api/model" // Ganti sesuai path sebenarnya

	"github.com/golang-jwt/jwt/v4"

	"github.com/labstack/echo/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
)

// JWTSecret retrieves the JWT secret key from environment variables.
// Fallback digunakan jika variabel tidak diset (untuk debugging lokal).
func JWTSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		fmt.Println("⚠️ WARNING: JWT_SECRET not set. Using fallback. DO NOT use in production.")
		secret = "TEST_SECRET_123"
	}
	return secret
}

// JWTMiddleware mengembalikan middleware Echo untuk memverifikasi JWT token
func JWTMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		// Menggunakan custom claims yang sesuai dengan struktur JWT kamu
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(model.JwtCustomClaims)
		},

		// Kunci dan algoritma signing JWT
		SigningKey:    []byte(JWTSecret()),
		SigningMethod: "HS256",

		// Token akan dicari di header Authorization dan cookie jwt
		TokenLookup: "header:Authorization:Bearer ,cookie:jwt",

		// Key dalam Echo context untuk menyimpan informasi user
		ContextKey: "user",

		// Handler error khusus untuk debug dan keperluan respons custom
		ErrorHandler: func(c echo.Context, err error) error {
			authHeader := c.Request().Header.Get("Authorization")

			// Debug Logging
			fmt.Printf("📥 JWT Auth Header: %q\n", authHeader)
			fmt.Printf("📥 Header Length: %d\n", len(authHeader))
			fmt.Println("🔐 JWT Secret Used:", JWTSecret())
			fmt.Println("🕒 Server Time:", time.Now().Format(time.RFC3339))
			fmt.Printf("❌ JWT Middleware Error: %v\n", err)

			return echo.NewHTTPError(http.StatusUnauthorized, map[string]interface{}{
				"message": "invalid or expired jwt",
				"detail":  err.Error(),
			})
		},
	})
}
